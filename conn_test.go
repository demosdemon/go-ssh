package ssh

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/demosdemon/go-ssh/fakes/net_fakes"
)

func Test_serverConn_Write(t *testing.T) {
	type fields struct {
		Conn          net.Conn
		idleTimeout   time.Duration
		maxDeadline   time.Time
		closeCanceler context.CancelFunc
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "ErrorUpdatingDeadline",
			fields: fields{
				Conn: &net_fakes.FakeConn{
					SetDeadlineStub: func(i time.Time) error {
						return assert.AnError
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ErrorWritingConnNetErrorWithCanceller",
			fields: fields{
				Conn: &net_fakes.FakeConn{
					WriteStub: func(bytes []byte) (i int, e error) {
						return 0, &net.OpError{Op: "Write"}
					},
				},
				closeCanceler: func() {},
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				Conn: &net_fakes.FakeConn{
					WriteStub: func(bytes []byte) (i int, e error) {
						return 42, nil
					},
				},
			},
			wantN: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &serverConn{
				Conn:          tt.fields.Conn,
				idleTimeout:   tt.fields.idleTimeout,
				maxDeadline:   tt.fields.maxDeadline,
				closeCanceler: tt.fields.closeCanceler,
			}
			gotN, err := c.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("serverConn.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("serverConn.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_serverConn_Read(t *testing.T) {
	type fields struct {
		Conn          net.Conn
		idleTimeout   time.Duration
		maxDeadline   time.Time
		closeCanceler context.CancelFunc
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "ErrorUpdatingDeadline",
			fields: fields{
				Conn: &net_fakes.FakeConn{
					SetDeadlineStub: func(i time.Time) error {
						return assert.AnError
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ErrorReadingConnNetErrorWithCanceller",
			fields: fields{
				Conn: &net_fakes.FakeConn{
					ReadStub: func(bytes []byte) (i int, e error) {
						return 0, &net.OpError{Op: "Read"}
					},
				},
				closeCanceler: func() {},
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				Conn: &net_fakes.FakeConn{
					ReadStub: func(bytes []byte) (i int, e error) {
						return 42, nil
					},
				},
			},
			wantN: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &serverConn{
				Conn:          tt.fields.Conn,
				idleTimeout:   tt.fields.idleTimeout,
				maxDeadline:   tt.fields.maxDeadline,
				closeCanceler: tt.fields.closeCanceler,
			}
			gotN, err := c.Read(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("serverConn.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("serverConn.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_serverConn_Close(t *testing.T) {
	type fields struct {
		Conn          net.Conn
		idleTimeout   time.Duration
		maxDeadline   time.Time
		closeCanceler context.CancelFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:   "NoCanceler",
			fields: fields{Conn: &net_fakes.FakeConn{}},
		},
		{
			name: "Canceler",
			fields: fields{
				Conn: &net_fakes.FakeConn{},
				closeCanceler: func() {

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &serverConn{
				Conn:          tt.fields.Conn,
				idleTimeout:   tt.fields.idleTimeout,
				maxDeadline:   tt.fields.maxDeadline,
				closeCanceler: tt.fields.closeCanceler,
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("serverConn.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serverConn_updateDeadline(t *testing.T) {
	type fields struct {
		Conn          net.Conn
		idleTimeout   time.Duration
		maxDeadline   time.Time
		closeCanceler context.CancelFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "IdleTimeout",
			fields: fields{
				Conn:        &net_fakes.FakeConn{},
				idleTimeout: time.Millisecond,
			},
		},
		{
			name: "IdleTimeoutMaxDeadline",
			fields: fields{
				Conn:        &net_fakes.FakeConn{},
				idleTimeout: time.Millisecond,
				maxDeadline: time.Now(),
			},
		},
		{
			name: "NoIdleTimeout",
			fields: fields{
				Conn: &net_fakes.FakeConn{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &serverConn{
				Conn:          tt.fields.Conn,
				idleTimeout:   tt.fields.idleTimeout,
				maxDeadline:   tt.fields.maxDeadline,
				closeCanceler: tt.fields.closeCanceler,
			}
			if err := c.updateDeadline(); (err != nil) != tt.wantErr {
				t.Errorf("serverConn.updateDeadline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
