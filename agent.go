package ssh

import (
	"io/ioutil"
	"net"
	"path"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"

	"github.com/demosdemon/go-copier"
)

const (
	agentRequestType = "auth-agent-req@openssh.com"
	agentChannelType = "auth-agent@openssh.com"
	agentTempDir     = "auth-agent"
	agentListenFile  = "listener.sock"
)

var contextKeyAgentRequest = &contextKey{"auth-agent-req"}

func SetAgentRequested(ctx Context) {
	ctx.SetValue(contextKeyAgentRequest, true)
}

func AgentRequested(sess Session) bool {
	return sess.Context().Value(contextKeyAgentRequest) == true
}

func NewAgentListener() (net.Listener, error) {
	dir, err := ioutil.TempDir("", agentTempDir)
	if err != nil {
		return nil, err
	}
	return net.Listen("unix", path.Join(dir, agentListenFile))
}

func ForwardAgentConnections(l net.Listener, s Session) {
	sshConn := s.Context().Value(ContextKeyConn).(ssh.Conn)
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}

		go func(conn net.Conn) {
			defer conn.Close()

			log := s.Logger().WithFields(logrus.Fields{
				"local_addr":  conn.LocalAddr(),
				"remote_addr": conn.RemoteAddr(),
			})

			defer log.Trace("ssh agent connection ended")

			log.Trace("new ssh agent connection")

			channel, reqs, err := sshConn.OpenChannel(agentChannelType, nil)
			if err != nil {
				log.WithError(err).Trace("error opening agent channel")
				return
			}

			go ssh.DiscardRequests(reqs)

			var cg copier.Group
			defer cg.Shutdown(s.Context())

			cg.Add(conn, channel)
			cg.Add(channel, conn)

			_ = cg.Wait(s.Context())
		}(conn)
	}
}
