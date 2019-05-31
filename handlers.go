package ssh

import "golang.org/x/crypto/ssh"

type RequestHandler interface {
	HandleRequest(ctx Context, srv *Server, req *ssh.Request) (ok bool, payload []byte)
}

type ChannelHandler interface {
	HandleChannel(ctx Context, srv *Server, conn *ssh.ServerConn, newChan ssh.NewChannel)
}
