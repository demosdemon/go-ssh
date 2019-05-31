package ssh

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

var (
	ErrServerClosed     = errors.New("server closed")
	ErrPermissionDenied = errors.New("permission denied")
)

type Server struct {
	Addr        string
	Handler     Handler
	HostSigners []Signer
	Version     string
	Logger      Logger

	KeyboardInteractiveHandler    KeyboardInteractiveHandler
	PasswordHandler               PasswordHandler
	PublicKeyHandler              PublicKeyHandler
	PtyCallback                   PtyCallback
	ConnCallback                  ConnCallback
	LocalPortForwardingCallback   LocalPortForwardingCallback
	ReversePortForwardingCallback ReversePortForwardingCallback
	DefaultServerConfigCallback   DefaultServerConfigCallback
	SessionRequestCallback        SessionRequestCallback
	IdleTimeout                   time.Duration
	MaxTimeout                    time.Duration
	ChannelHandlers               map[string]ChannelHandler
	RequestHandlers               map[string]RequestHandler

	listenerWg sync.WaitGroup
	mu         sync.Mutex
	listeners  map[net.Listener]struct{}
	conns      map[*ssh.ServerConn]struct{}
	connWg     sync.WaitGroup
	doneChan   chan struct{}
}

func (srv *Server) Handle(fn Handler) {
	srv.Handler = fn
}

func (srv *Server) AddHostKey(key Signer) {
	srv.HostSigners = append(srv.HostSigners, key)
}

func (srv *Server) SetOption(option Option) error {
	return option(srv)
}

func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":22"
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(l)
}

func (srv *Server) Serve(l net.Listener) error {
	srv.ensureHandlers()

	if err := srv.ensureHostSigner(); err != nil {
		return errors.Wrap(err, "error generating default host key")
	}

	var tempDelay time.Duration

	srv.trackListener(l, true)
	defer srv.trackListener(l, false)

	for {
		conn, err := l.Accept()
		if err != nil {
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}

			if err, ok := err.(net.Error); ok && err.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}

				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}

				time.Sleep(tempDelay)
				continue
			}

			return err
		}
		go srv.handleConn(conn)
	}
}

func (srv *Server) Close() error {
	srv.mu.Lock()
	srv.closeDoneChanLocked()
	err := multierror.Append(nil, srv.closeListenersLocked())
	for c := range srv.conns {
		err = multierror.Append(err, c.Close())
		delete(srv.conns, c)
	}
	srv.mu.Unlock()
	return err.ErrorOrNil()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	srv.mu.Lock()
	srv.closeDoneChanLocked()
	err := multierror.Append(nil, srv.closeListenersLocked())
	srv.mu.Unlock()

	done := make(chan struct{})
	go func() {
		srv.listenerWg.Wait()
		srv.connWg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		err = multierror.Append(err, ctx.Err())
	case <-done:
	}

	return err.ErrorOrNil()
}

func (srv *Server) handleConn(newConn net.Conn) {
	if srv.ConnCallback != nil {
		cbConn := srv.ConnCallback(newConn)
		if cbConn == nil {
			newConn.Close()
			return
		}
		newConn = cbConn
	}

	ctx, cancel := newContext(srv)
	conn := serverConn{
		Conn:          newConn,
		idleTimeout:   srv.IdleTimeout,
		closeCanceler: cancel,
	}
	if srv.MaxTimeout > 0 {
		conn.maxDeadline = time.Now().Add(srv.MaxTimeout)
	}
	defer conn.Close()

	sshConn, chans, reqs, err := ssh.NewServerConn(&conn, srv.config(ctx))
	if err != nil {
		return
	}

	srv.trackConn(sshConn, true)
	defer srv.trackConn(sshConn, false)

	ctx.SetValue(ContextKeyConn, sshConn)
	applyConnMetadata(ctx, sshConn)

	go srv.handleRequests(ctx, reqs)

	for ch := range chans {
		handler := srv.ChannelHandlers[ch.ChannelType()]
		if handler == nil {
			_ = ch.Reject(ssh.UnknownChannelType, "unsupported channel type")
			continue
		}

		go handler.HandleChannel(ctx, srv, sshConn, ch)
	}
}

func (srv *Server) handleRequests(ctx Context, in <-chan *ssh.Request) {
	for req := range in {
		var ok bool
		var payload []byte
		if handler := srv.RequestHandlers[req.Type]; handler != nil {
			ok, payload = handler.HandleRequest(ctx, srv, req)
		}
		if req.WantReply {
			_ = req.Reply(ok, payload)
		}
	}
}

func (srv *Server) config(ctx Context) *ssh.ServerConfig {
	var config *ssh.ServerConfig

	if srv.DefaultServerConfigCallback == nil {
		config = new(ssh.ServerConfig)
	} else {
		config = srv.DefaultServerConfigCallback(ctx)
	}

	for _, signer := range srv.HostSigners {
		config.AddHostKey(signer)
	}

	if srv.PasswordHandler == nil && srv.PublicKeyHandler == nil && srv.KeyboardInteractiveHandler == nil {
		config.NoClientAuth = true
	}

	if srv.Version != "" {
		config.ServerVersion = "SSH-2.0-" + srv.Version
	}

	if srv.PasswordHandler != nil {
		config.PasswordCallback = func(conn ssh.ConnMetadata, password []byte) (permissions *ssh.Permissions, err error) {
			applyConnMetadata(ctx, conn)
			if ok := srv.PasswordHandler(ctx, string(password)); !ok {
				err = ErrPermissionDenied
			}
			permissions = ctx.Permissions().Permissions
			return permissions, err
		}
	}

	if srv.PublicKeyHandler != nil {
		config.PublicKeyCallback = func(conn ssh.ConnMetadata, key ssh.PublicKey) (permissions *ssh.Permissions, err error) {
			applyConnMetadata(ctx, conn)
			if ok := srv.PublicKeyHandler(ctx, key); !ok {
				err = ErrPermissionDenied
			}
			permissions = ctx.Permissions().Permissions
			return permissions, err
		}
	}

	if srv.KeyboardInteractiveHandler != nil {
		config.KeyboardInteractiveCallback = func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (permissions *ssh.Permissions, err error) {
			applyConnMetadata(ctx, conn)
			if ok := srv.KeyboardInteractiveHandler(ctx, client); !ok {
				err = ErrPermissionDenied
			}
			permissions = ctx.Permissions().Permissions
			return permissions, err
		}
	}

	return config
}

func (srv *Server) ensureHostSigner() error {
	if len(srv.HostSigners) == 0 {
		signer, err := generateSigner()
		if err != nil {
			return err
		}
		srv.HostSigners = append(srv.HostSigners, signer)
	}
	return nil
}

func (srv *Server) ensureHandlers() {
	if srv.Handler == nil {
		srv.Handler = DefaultHandler
	}

	if srv.RequestHandlers == nil {
		srv.RequestHandlers = make(map[string]RequestHandler)
	}
	if _, ok := srv.RequestHandlers[tcpipForwardType]; !ok {
		hdlr := &forwardedTCPHandler{}
		srv.RequestHandlers[tcpipForwardType] = hdlr
		srv.RequestHandlers[cancelTcpIpForwardType] = hdlr
	}

	if srv.ChannelHandlers == nil {
		srv.ChannelHandlers = make(map[string]ChannelHandler)
	}
	if _, ok := srv.ChannelHandlers["session"]; !ok {
		srv.ChannelHandlers["session"] = SessionHandler{}
	}
	if _, ok := srv.ChannelHandlers["direct-tcpip"]; !ok {
		srv.ChannelHandlers["direct-tcpip"] = &DirectTcpIpHandler{}
	}
}

func (srv *Server) getDoneChan() <-chan struct{} {
	srv.mu.Lock()
	ch := srv.getDoneChanLocked()
	srv.mu.Unlock()
	return ch
}

func (srv *Server) getDoneChanLocked() chan struct{} {
	if srv.doneChan == nil {
		srv.doneChan = make(chan struct{})
	}
	return srv.doneChan
}

func (srv *Server) closeDoneChanLocked() {
	ch := srv.getDoneChanLocked()
	select {
	case <-ch:
	default:
		close(ch)
	}
}

func (srv *Server) closeListenersLocked() error {
	var err *multierror.Error
	for l := range srv.listeners {
		err = multierror.Append(err, l.Close())
		delete(srv.listeners, l)
	}
	return err.ErrorOrNil()
}

func (srv *Server) trackListener(l net.Listener, add bool) {
	srv.mu.Lock()

	if srv.listeners == nil {
		srv.listeners = make(map[net.Listener]struct{})
	}

	if add {
		if len(srv.listeners) == 0 && len(srv.conns) == 0 {
			srv.doneChan = nil
		}
		srv.listeners[l] = struct{}{}
		srv.listenerWg.Add(1)
	} else {
		delete(srv.listeners, l)
		srv.listenerWg.Done()
	}

	srv.mu.Unlock()
}

func (srv *Server) trackConn(c *ssh.ServerConn, add bool) {
	srv.mu.Lock()

	if srv.conns == nil {
		srv.conns = make(map[*ssh.ServerConn]struct{})
	}

	if add {
		srv.conns[c] = struct{}{}
		srv.connWg.Add(1)
	} else {
		delete(srv.conns, c)
		srv.connWg.Done()
	}

	srv.mu.Unlock()
}

func (srv *Server) logger() Logger {
	e := srv.Logger
	if e == nil {
		e = logrus.NewEntry(logrus.StandardLogger())
	}
	return e
}
