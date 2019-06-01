package ssh_test

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"

	. "github.com/demosdemon/go-ssh"
	. "github.com/demosdemon/go-ssh/fakes"
	"github.com/demosdemon/go-ssh/fakes/net_fakes"
	"github.com/demosdemon/go-ssh/fakes/ssh_fakes"
)

var testFormatter = logrus.TextFormatter{
	DisableColors:    true,
	DisableTimestamp: true,
	QuoteEmptyFields: true,
}

func TestSetAgentRequested(t *testing.T) {
	var fakeContext FakeContext

	SetAgentRequested(&fakeContext)

	assert.Equal(t, 1, fakeContext.SetValueCallCount())
}

func TestAgentRequested(t *testing.T) {
	var fakeSession FakeSession
	var fakeContext FakeContext

	fakeSession.ContextReturns(&fakeContext)

	fakeContext.ValueReturns(true)
	assert.True(t, AgentRequested(&fakeSession))

	fakeContext.ValueReturns(false)
	assert.False(t, AgentRequested(&fakeSession))

	assert.Equal(t, 2, fakeSession.ContextCallCount())
	assert.Equal(t, 2, fakeContext.ValueCallCount())
}

func TestNewAgentListener(t *testing.T) {
	tmpdirEnv := os.Getenv("TMPDIR")
	tmpdir, err := ioutil.TempDir("/tmp", "listener")
	require.NoError(t, err)

	require.NoError(t, os.Setenv("TMPDIR", tmpdir))
	defer func() {
		require.NoError(t, os.Setenv("TMPDIR", tmpdirEnv))
	}()

	t.Run("success", func(t *testing.T) {
		l, err := NewAgentListener()
		require.NoError(t, err)
		require.NotNil(t, l)
		require.NoError(t, l.Close())
	})

	t.Run("error", func(t *testing.T) {
		require.NoError(t, os.Chmod(tmpdir, 0x200))
		l, err := NewAgentListener()
		require.Error(t, err)
		require.Nil(t, l)
	})
}

type ForwardAgentConnectionsCase struct {
	Name               string
	ConnPayload        []byte
	ChannelPayload     []byte
	OpenChannelError   error
	LoggerOut          string
	ConnCloseCallCount int
	ConnReadCallCount  int
	ConnWriteCallCount int
	ChanCloseCallCount int
	ChanReadCallCount  int
	ChanWriteCallCount int

	initOnce   sync.Once
	logger     logrus.Logger
	loggerOut  strings.Builder
	conn       net_fakes.FakeConn
	connBuf    bytes.Buffer
	channel    ssh_fakes.FakeChannel
	channelBuf bytes.Buffer
}

func (c *ForwardAgentConnectionsCase) Validate(t *testing.T) {
	assert.Equal(t, c.LoggerOut, c.loggerOut.String())
	assert.Equal(t, c.ConnCloseCallCount, c.conn.CloseCallCount())
	assert.Equal(t, c.ConnReadCallCount, c.conn.ReadCallCount())
	assert.Equal(t, c.ConnWriteCallCount, c.conn.WriteCallCount())
	assert.Equal(t, c.ChanCloseCallCount, c.channel.CloseCallCount())
	assert.Equal(t, c.ChanReadCallCount, c.channel.ReadCallCount())
	assert.Equal(t, c.ChanWriteCallCount, c.channel.WriteCallCount())
}

func (c *ForwardAgentConnectionsCase) Init() {
	c.initOnce.Do(c.init)
}

func (c *ForwardAgentConnectionsCase) init() {
	c.logger = logrus.Logger{
		Out:          &c.loggerOut,
		Formatter:    &testFormatter,
		ReportCaller: false,
		Level:        logrus.TraceLevel,
		ExitFunc:     func(code int) { panic(errors.Errorf("exit %d", code)) },
	}

	connReader := bytes.NewReader(c.ConnPayload)
	c.conn.ReadStub = connReader.Read
	c.conn.WriteStub = c.connBuf.Write

	chanReader := bytes.NewReader(c.ChannelPayload)
	c.channel.ReadStub = chanReader.Read
	c.channel.WriteStub = c.channelBuf.Write
}

func (c *ForwardAgentConnectionsCase) Logger() Logger {
	c.Init()
	return &c.logger
}

func (c *ForwardAgentConnectionsCase) Conn() net.Conn {
	c.Init()
	return &c.conn
}

func (c *ForwardAgentConnectionsCase) Channel() ssh.Channel {
	c.Init()
	return &c.channel
}

func TestForwardAgentConnections(t *testing.T) {
	var (
		fakeSession  FakeSession
		fakeContext  FakeContext
		fakeSSHConn  ssh_fakes.FakeConn
		fakeListener net_fakes.FakeListener
	)

	fakeSession.ContextReturns(&fakeContext)
	fakeContext.ValueReturns(&fakeSSHConn)

	closedChan := make(chan *ssh.Request)
	close(closedChan)

	cases := []ForwardAgentConnectionsCase{
		{
			Name:               "OpenChannelError",
			OpenChannelError:   assert.AnError,
			LoggerOut:          "level=trace msg=\"new ssh agent connection\"\nlevel=trace msg=\"error opening agent channel\" error=\"assert.AnError general error for testing\"\nlevel=trace msg=\"ssh agent connection ended\"\n",
			ConnCloseCallCount: 1,
		},
		{
			Name:               "NoChannelError",
			ConnPayload:        []byte("the connection is sending this to the channel"),
			ChannelPayload:     []byte("the channel is sending this to the connection"),
			LoggerOut:          "level=trace msg=\"new ssh agent connection\"\nlevel=trace msg=\"ssh agent connection ended\"\n",
			ConnCloseCallCount: 1,
			ConnReadCallCount:  2,
			ConnWriteCallCount: 1,
			ChanReadCallCount:  2,
			ChanWriteCallCount: 1,
		},
	}

	idx := 0
	loggerDone := make(chan struct{})
	close(loggerDone)
	openChannelDone := make(chan struct{})
	close(openChannelDone)

	fakeListener.AcceptStub = func() (conn net.Conn, e error) {
		<-loggerDone
		loggerDone = make(chan struct{})
		<-openChannelDone
		openChannelDone = make(chan struct{})

		if idx < len(cases) {
			return cases[idx].Conn(), nil
		}

		return nil, assert.AnError
	}

	fakeSession.LoggerStub = func() Logger {
		l := cases[idx].Logger()
		close(loggerDone)
		return l
	}

	fakeSSHConn.OpenChannelStub = func(string, []byte) (channel ssh.Channel, requests <-chan *ssh.Request, e error) {
		channel = cases[idx].Channel()
		requests = closedChan
		e = cases[idx].OpenChannelError
		idx++
		close(openChannelDone)
		return channel, requests, e
	}

	ForwardAgentConnections(&fakeListener, &fakeSession)
	time.Sleep(time.Millisecond)
	require.Equal(t, len(cases)+1, fakeListener.AcceptCallCount())

	for _, c := range cases {
		t.Run(c.Name, c.Validate)
	}
}
