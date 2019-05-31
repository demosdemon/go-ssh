package ssh

import (
	"crypto/subtle"
	"net"

	"golang.org/x/crypto/ssh"
)

var DefaultHandler Handler

type Option func(*Server) error

type Handler func(Session)

type PublicKeyHandler func(ctx Context, key PublicKey) bool

type PasswordHandler func(ctx Context, password string) bool

type KeyboardInteractiveHandler func(ctx Context, challenger ssh.KeyboardInteractiveChallenge) bool

type PtyCallback func(ctx Context, pty Pty) bool

type SessionRequestCallback func(sess Session, requestType string) bool

type ConnCallback func(conn net.Conn) net.Conn

type LocalPortForwardingCallback func(ctx Context, destinationHost string, destinationPort uint32) bool

type ReversePortForwardingCallback func(ctx Context, bindHost string, bindPort uint32) bool

type DefaultServerConfigCallback func(ctx Context) *ssh.ServerConfig

type Window struct {
	Cols   uint32
	Rows   uint32
	Width  uint32
	Height uint32
}

type Pty struct {
	Term   string
	Window Window
	Modes  []byte
}

func Serve(l net.Listener, handler Handler, options ...Option) error {
	srv := &Server{Handler: handler}
	for _, option := range options {
		if err := srv.SetOption(option); err != nil {
			return err
		}
	}
	return srv.Serve(l)
}

func ListenAndServe(addr string, handler Handler, options ...Option) error {
	srv := &Server{Addr: addr, Handler: handler}
	for _, option := range options {
		if err := srv.SetOption(option); err != nil {
			return err
		}
	}
	return srv.ListenAndServe()
}

func Handle(handler Handler) {
	DefaultHandler = handler
}

func KeysEqual(ak, bk PublicKey) bool {
	if ak == nil || bk == nil {
		return false
	}

	a := ak.Marshal()
	b := bk.Marshal()
	return (len(a) == len(b)) && subtle.ConstantTimeCompare(a, b) == 1
}
