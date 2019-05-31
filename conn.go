package ssh

import (
	"context"
	"net"
	"time"
)

type serverConn struct {
	net.Conn

	idleTimeout   time.Duration
	maxDeadline   time.Time
	closeCanceler context.CancelFunc
}

func (c *serverConn) Write(b []byte) (n int, err error) {
	err = c.updateDeadline()
	if err != nil {
		return n, err
	}
	n, err = c.Conn.Write(b)
	if _, ok := err.(net.Error); ok && c.closeCanceler != nil {
		c.closeCanceler()
	}
	return n, err
}

func (c *serverConn) Read(b []byte) (n int, err error) {
	err = c.updateDeadline()
	if err != nil {
		return n, err
	}
	n, err = c.Conn.Read(b)
	if _, ok := err.(net.Error); ok && c.closeCanceler != nil {
		c.closeCanceler()
	}
	return n, err
}

func (c *serverConn) Close() (err error) {
	err = c.Conn.Close()
	if c.closeCanceler != nil {
		c.closeCanceler()
	}
	return err
}

func (c *serverConn) updateDeadline() error {
	switch {
	case c.idleTimeout > 0:
		idleDeadline := time.Now().Add(c.idleTimeout)
		if idleDeadline.Unix() < c.maxDeadline.Unix() || c.maxDeadline.IsZero() {
			return c.Conn.SetDeadline(idleDeadline)
		}
		fallthrough
	default:
		return c.Conn.SetDeadline(c.maxDeadline)
	}
}
