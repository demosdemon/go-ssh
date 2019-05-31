package ssh

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func PasswordAuth(fn PasswordHandler) Option {
	return func(srv *Server) error {
		srv.PasswordHandler = fn
		return nil
	}
}

func PublicKeyAuth(fn PublicKeyHandler) Option {
	return func(srv *Server) error {
		srv.PublicKeyHandler = fn
		return nil
	}
}

func HostKeyFile(filepath string) Option {
	return func(srv *Server) error {
		pemBytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		signer, err := ssh.ParsePrivateKey(pemBytes)
		if err != nil {
			return err
		}

		srv.AddHostKey(signer)
		return nil
	}
}

func HostKeyPEM(bytes []byte) Option {
	return func(srv *Server) error {
		signer, err := ssh.ParsePrivateKey(bytes)
		if err != nil {
			return err
		}

		srv.AddHostKey(signer)
		return nil
	}
}

func NoPty() Option {
	return func(srv *Server) error {
		srv.PtyCallback = func(ctx Context, pty Pty) bool { return false }
		return nil
	}
}

func WrapConn(fn ConnCallback) Option {
	return func(srv *Server) error {
		srv.ConnCallback = fn
		return nil
	}
}
