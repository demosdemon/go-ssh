package ssh

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/anmitsu/go-shlex"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type Session interface {
	ssh.Channel

	Logger() Logger
	User() string
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
	Environ() []string
	Exit(code int) error
	ExitSignal(signal os.Signal, coreDumped bool, message, languageTag string) error
	Command() []string
	PublicKey() PublicKey
	Context() context.Context
	Permissions() Permissions
	Pty() (Pty, <-chan Window, bool)

	Notify(c chan<- Signal, sig ...Signal)
	Reset(sig ...Signal)
	Stop(c chan<- Signal)
}

type SessionHandler struct{}

func (SessionHandler) HandleChannel(ctx Context, srv *Server, conn *ssh.ServerConn, newChan ssh.NewChannel) {
	ch, reqs, err := newChan.Accept()
	if err != nil {
		return
	}

	sess := session{
		Channel:   ch,
		conn:      conn,
		handler:   srv.Handler,
		ptyCb:     srv.PtyCallback,
		sessReqCb: srv.SessionRequestCallback,
		ctx:       ctx,
	}

	sess.handleRequests(reqs)
}

type session struct {
	sync.Mutex
	ssh.Channel
	conn      *ssh.ServerConn
	handler   Handler
	handled   bool
	exited    bool
	pty       *Pty
	winch     chan Window
	env       []string
	ptyCb     PtyCallback
	sessReqCb SessionRequestCallback
	cmd       []string
	ctx       Context
	signals   map[Signal][]chan<- Signal
}

func (s *session) Logger() Logger           { return s.ctx.Logger() }
func (s *session) User() string             { return s.conn.User() }
func (s *session) RemoteAddr() net.Addr     { return s.conn.RemoteAddr() }
func (s *session) LocalAddr() net.Addr      { return s.conn.LocalAddr() }
func (s *session) Environ() []string        { return append([]string(nil), s.env...) }
func (s *session) Command() []string        { return append([]string(nil), s.cmd...) }
func (s *session) Context() context.Context { return s.ctx }
func (s *session) Permissions() Permissions { return Permissions{Permissions: s.conn.Permissions} }

func (s *session) exit() error {
	s.Lock()
	defer s.Unlock()

	if s.exited {
		return errors.New("exit called multiple times")
	}

	s.exited = true
	return nil
}

func (s *session) Exit(code int) error {
	err := s.exit()
	if err != nil {
		return err
	}

	msg := msgExitStatus{uint32(code)}
	_, err = s.SendRequest("exit-status", false, ssh.Marshal(msg))
	return err
}

func (s *session) ExitSignal(signal os.Signal, coreDumped bool, message, languageTag string) error {
	err := s.exit()
	if err != nil {
		return err
	}

	sig := NewSignal(signal)
	if sig == "unknown" {
		return errors.Errorf("unknown signal %v", signal)
	}

	msg := msgExitSignal{
		Signal:       sig,
		CoreDumped:   coreDumped,
		ErrorMessage: message,
		LanguageTag:  languageTag,
	}
	_, err = s.SendRequest("exit-signal", false, ssh.Marshal(msg))
	return err
}

func (s *session) PublicKey() PublicKey {
	v, _ := s.ctx.Value(ContextKeyPublicKey).(PublicKey)
	return v
}

func (s *session) Pty() (Pty, <-chan Window, bool) {
	s.Lock()
	defer s.Unlock()
	if s.pty == nil {
		return Pty{}, nil, false
	}
	return *s.pty, s.winch, true
}

func (s *session) Notify(c chan<- Signal, sig ...Signal) {
	if len(sig) == 0 {
		sig = signals[:]
	}

	s.Lock()

	if s.signals == nil {
		s.signals = make(map[Signal][]chan<- Signal, totalSignals)
	}

	for _, sig := range sig {
		s.signals[sig] = append(s.signals[sig], c)
	}

	s.Unlock()
}

func (s *session) Reset(sig ...Signal) {
	s.Lock()

	for _, sig := range sig {
		delete(s.signals, sig)
	}

	s.Unlock()
}

func (s *session) Stop(c chan<- Signal) {
	s.Lock()
	defer s.Unlock()

	if c == nil {
		s.signals = nil
		return
	}

	for k, v := range s.signals {
		for idx, ch := range v {
			if c == ch {
				copy(v[idx:], v[idx+1:])
				v[len(v)-1] = nil
				v = v[:len(v)-1]
				break
			}
		}
		s.signals[k] = v
	}
}

func (s *session) signal(sig Signal) {
	v := s.signals[sig]
	for _, c := range v {
		go func(c chan<- Signal) {
			c <- sig
		}(c)
	}
}

func (s *session) handleRequests(reqs <-chan *ssh.Request) {
	defer func() {
		if s.winch != nil {
			close(s.winch)
		}
	}()

	for req := range reqs {
		ok := false

		switch req.Type {
		case "shell", "exec":
			if s.handled {
				break
			}

			var msg msgExec
			_ = ssh.Unmarshal(req.Payload, &msg)
			s.cmd, _ = shlex.Split(msg.Command, true)

			if s.sessReqCb != nil && !s.sessReqCb(s, req.Type) {
				s.cmd = nil
				break
			}

			s.handled = true
			go func() {
				s.handler(s)
				_ = s.Exit(0)
				_ = s.Close()
			}()

			ok = true
		case "env":
			if s.handled {
				break
			}

			var msg msgEnv
			_ = ssh.Unmarshal(req.Payload, &msg)
			s.env = append(s.env, fmt.Sprintf("%s=%s", msg.Key, msg.Value))

			ok = true
		case "signal":
			var msg msgSignal
			_ = ssh.Unmarshal(req.Payload, &msg)
			s.signal(msg.Signal)
			ok = true
		case "pty-req":
			if s.handled || s.pty != nil {
				break
			}

			var msg msgPtyReq
			_ = ssh.Unmarshal(req.Payload, &msg)

			ptyReq := msg.PTY()

			ok = true
			if s.ptyCb != nil {
				ok = s.ptyCb(s.ctx, ptyReq)
			}
			if !ok {
				break
			}

			s.pty = &ptyReq
			s.winch = make(chan Window, 1)
			s.winch <- ptyReq.Window
		case "window-change":
			if s.pty == nil {
				break
			}

			var msg Window
			_ = ssh.Unmarshal(req.Payload, &msg)
			s.pty.Window = msg
			s.winch <- msg
			ok = true
		case agentRequestType:
			SetAgentRequested(s.ctx)
			ok = true
		}

		if req.WantReply {
			_ = req.Reply(ok, nil)
		}
	}
}
