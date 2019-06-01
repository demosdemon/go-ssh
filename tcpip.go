package ssh

import (
	"net"
	"strconv"
	"sync"

	"golang.org/x/crypto/ssh"

	"github.com/demosdemon/go-copier"
)

const (
	forwardedTCPChannelType = "forwarded-tcpip"
	tcpipForwardType        = "tcpip-forward"
	cancelTcpIpForwardType  = "cancel-tcpip-forward"
)

type DirectTcpIpHandler struct {
	net.Dialer
}

func (h *DirectTcpIpHandler) HandleChannel(ctx Context, srv *Server, conn *ssh.ServerConn, newChan ssh.NewChannel) {
	var msg msgForwarding

	if err := ssh.Unmarshal(newChan.ExtraData(), &msg); err != nil {
		_ = newChan.Reject(ssh.ConnectionFailed, "error parsing forwarded data: "+err.Error())
		return
	}

	if srv.LocalPortForwardingCallback == nil || !srv.LocalPortForwardingCallback(ctx, msg.DestAddr, msg.DestPort) {
		_ = newChan.Reject(ssh.Prohibited, "port forwarding is disabled")
		return
	}

	dest := net.JoinHostPort(msg.DestAddr, strconv.Itoa(int(msg.DestPort)))
	dconn, err := h.DialContext(ctx, "tcp", dest)
	if err != nil {
		_ = newChan.Reject(ssh.ConnectionFailed, err.Error())
		return
	}

	ch, reqs, err := newChan.Accept()
	if err != nil {
		_ = dconn.Close()
		return
	}

	go ssh.DiscardRequests(reqs)

	var cg copier.Group
	cg.Add(ch, dconn)
	cg.Add(dconn, ch)
	_ = cg.Wait(ctx)
}

type forwardedTCPHandler struct {
	sync.Mutex
	forwards map[string]net.Listener
}

func (h *forwardedTCPHandler) forward(ctx Context, srv *Server, req *ssh.Request) (bool, []byte) {
	conn := ctx.Value(ContextKeyConn).(*ssh.ServerConn)
	var msg msgRemoteForward
	if err := ssh.Unmarshal(req.Payload, &msg); err != nil {
		return false, nil
	}

	if srv.ReversePortForwardingCallback == nil || !srv.ReversePortForwardingCallback(ctx, msg.BindAddr, msg.BindPort) {
		return false, []byte("port forwarding is disabled")
	}

	addr := net.JoinHostPort(msg.BindAddr, strconv.Itoa(int(msg.BindPort)))
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return false, nil
	}

	h.Lock()
	h.forwards[addr] = l
	h.Unlock()

	destAddr := l.Addr().(*net.TCPAddr)

	go func() {
		<-ctx.Done()

		h.Lock()
		l := h.forwards[addr]
		h.Unlock()

		if l != nil {
			_ = l.Close()
		}
	}()

	go func() {
		for {
			subConn, err := l.Accept()

			if err != nil {
				break
			}

			originAddr := subConn.RemoteAddr().(*net.TCPAddr)

			go func() {
				ch, reqs, err := conn.OpenChannel(forwardedTCPChannelType, ssh.Marshal(msgForwarding{
					DestAddr:   destAddr.IP.String(),
					DestPort:   uint32(destAddr.Port),
					OriginAddr: originAddr.IP.String(),
					OriginPort: uint32(originAddr.Port),
				}))

				if err != nil {
					_ = subConn.Close()
					return
				}

				go ssh.DiscardRequests(reqs)

				var cg copier.Group
				cg.Add(ch, subConn)
				cg.Add(subConn, ch)
				_ = cg.Wait(ctx)
			}()
		}

		h.Lock()
		delete(h.forwards, addr)
		h.Unlock()
	}()

	return true, ssh.Marshal(msgRemoteForwardSuccess{uint32(destAddr.Port)})
}

func (h *forwardedTCPHandler) cancel(ctx Context, srv *Server, req *ssh.Request) (bool, []byte) {
	var msg msgRemoteForward
	if err := ssh.Unmarshal(req.Payload, &msg); err != nil {
		return false, nil
	}

	addr := net.JoinHostPort(msg.BindAddr, strconv.Itoa(int(msg.BindPort)))

	h.Lock()
	l := h.forwards[addr]
	h.Unlock()

	if l != nil {
		_ = l.Close()
	}

	return true, nil
}

func (h *forwardedTCPHandler) HandleRequest(ctx Context, srv *Server, req *ssh.Request) (bool, []byte) {
	h.Lock()
	if h.forwards == nil {
		h.forwards = make(map[string]net.Listener)
	}
	h.Unlock()

	switch req.Type {
	case tcpipForwardType:
		return h.forward(ctx, srv, req)
	case cancelTcpIpForwardType:
		return h.cancel(ctx, srv, req)
	default:
		return false, nil
	}
}
