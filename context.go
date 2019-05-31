package ssh

import (
	"context"
	"encoding/hex"
	"net"
	"sync"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type contextKey struct {
	name string
}

var (
	ContextKeySessionID     = &contextKey{"session-id"}
	ContextKeyClientVersion = &contextKey{"client-version"}
	ContextKeyServerVersion = &contextKey{"server-version"}
	ContextKeyUser          = &contextKey{"user"}
	ContextKeyLocalAddr     = &contextKey{"local-addr"}
	ContextKeyRemoteAddr    = &contextKey{"remote-addr"}
	ContextKeyPermissions   = &contextKey{"permissions"}
	ContextKeyLogger        = &contextKey{"logger"}

	ContextKeyServer    = &contextKey{"ssh-server"}
	ContextKeyConn      = &contextKey{"ssh-conn"}
	ContextKeyPublicKey = &contextKey{"public-key"}
)

type Context interface {
	context.Context
	sync.Locker

	SetValue(key, value interface{})
	SessionID() string
	ClientVersion() string
	ServerVersion() string
	User() string
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Permissions() *Permissions
	Logger() Logger
}

type sshContext struct {
	context.Context
	sync.Mutex
}

func newContext(srv *Server) (*sshContext, context.CancelFunc) {
	innerCtx, cancel := context.WithCancel(context.Background())

	ctx := &sshContext{Context: innerCtx}
	ctx.SetValue(ContextKeyServer, srv)

	perms := &Permissions{Permissions: new(ssh.Permissions)}
	ctx.SetValue(ContextKeyPermissions, perms)

	return ctx, cancel
}

func applyConnMetadata(ctx Context, conn ssh.ConnMetadata) {
	if ctx.Value(ContextKeySessionID) != nil {
		return
	}
	ctx.SetValue(ContextKeySessionID, hex.EncodeToString(conn.SessionID()))
	ctx.SetValue(ContextKeyClientVersion, string(conn.ClientVersion()))
	ctx.SetValue(ContextKeyServerVersion, string(conn.ServerVersion()))
	ctx.SetValue(ContextKeyUser, conn.User())
	ctx.SetValue(ContextKeyLocalAddr, conn.LocalAddr())
	ctx.SetValue(ContextKeyRemoteAddr, conn.RemoteAddr())
}

func (ctx *sshContext) SetValue(key, value interface{}) {
	ctx.Lock()
	ctx.Context = context.WithValue(ctx.Context, key, value)
	ctx.Unlock()
}

func (ctx *sshContext) SessionID() string {
	return ctx.Value(ContextKeySessionID).(string)
}

func (ctx *sshContext) ClientVersion() string {
	return ctx.Value(ContextKeyClientVersion).(string)
}

func (ctx *sshContext) ServerVersion() string {
	return ctx.Value(ContextKeyServerVersion).(string)
}

func (ctx *sshContext) User() string {
	return ctx.Value(ContextKeyUser).(string)
}

func (ctx *sshContext) LocalAddr() net.Addr {
	return ctx.Value(ContextKeyLocalAddr).(net.Addr)
}

func (ctx *sshContext) RemoteAddr() net.Addr {
	return ctx.Value(ContextKeyRemoteAddr).(net.Addr)
}

func (ctx *sshContext) Permissions() *Permissions {
	return ctx.Value(ContextKeyPermissions).(*Permissions)
}

func (ctx *sshContext) Logger() Logger {
	log := ctx.logger()
	if v := ctx.Value(ContextKeySessionID); v == nil {
		return log
	}

	return log.WithFields(logrus.Fields{
		"user":           ctx.User(),
		"session_id":     ctx.SessionID(),
		"client_version": ctx.ClientVersion(),
		"server_version": ctx.ServerVersion(),
		"remote_addr":    ctx.RemoteAddr().String(),
		"local_addr":     ctx.LocalAddr().String(),
	})
}

func (ctx *sshContext) logger() Logger {
	v, _ := ctx.Value(ContextKeyLogger).(Logger)
	if v != nil {
		return v
	}

	srv, _ := ctx.Value(ContextKeyServer).(*Server)
	if srv != nil {
		return srv
	}

	return logrus.NewEntry(logrus.StandardLogger())
}
