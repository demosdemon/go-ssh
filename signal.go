package ssh

import (
	"os"
	"syscall"

	"golang.org/x/crypto/ssh"
)

type Signal ssh.Signal

const (
	SIGABRT      = Signal(ssh.SIGABRT)
	SIGALRM      = Signal(ssh.SIGALRM)
	SIGFPE       = Signal(ssh.SIGFPE)
	SIGHUP       = Signal(ssh.SIGHUP)
	SIGILL       = Signal(ssh.SIGILL)
	SIGKILL      = Signal(ssh.SIGKILL)
	SIGPIPE      = Signal(ssh.SIGPIPE)
	SIGQUIT      = Signal(ssh.SIGQUIT)
	SIGSEGV      = Signal(ssh.SIGSEGV)
	SIGTERM      = Signal(ssh.SIGTERM)
	SIGUSR1      = Signal(ssh.SIGUSR1)
	SIGUSR2      = Signal(ssh.SIGUSR2)
	totalSignals = iota
)

var signals = [totalSignals]Signal{
	SIGABRT,
	SIGALRM,
	SIGFPE,
	SIGHUP,
	SIGILL,
	SIGKILL,
	SIGPIPE,
	SIGQUIT,
	SIGSEGV,
	SIGTERM,
	SIGUSR1,
	SIGUSR2,
}

var signalMap = map[os.Signal]Signal{
	syscall.SIGABRT: SIGABRT,
	syscall.SIGALRM: SIGALRM,
	syscall.SIGFPE:  SIGFPE,
	syscall.SIGHUP:  SIGHUP,
	syscall.SIGILL:  SIGILL,
	syscall.SIGKILL: SIGKILL,
	syscall.SIGPIPE: SIGPIPE,
	syscall.SIGQUIT: SIGQUIT,
	syscall.SIGSEGV: SIGSEGV,
	syscall.SIGTERM: SIGTERM,
	syscall.SIGUSR1: SIGUSR1,
	syscall.SIGUSR2: SIGUSR2,
}

var osSignalMap = map[Signal]os.Signal{
	SIGABRT: syscall.SIGABRT,
	SIGALRM: syscall.SIGALRM,
	SIGFPE:  syscall.SIGFPE,
	SIGHUP:  syscall.SIGHUP,
	SIGILL:  syscall.SIGILL,
	SIGKILL: syscall.SIGKILL,
	SIGPIPE: syscall.SIGPIPE,
	SIGQUIT: syscall.SIGQUIT,
	SIGSEGV: syscall.SIGSEGV,
	SIGTERM: syscall.SIGTERM,
	SIGUSR1: syscall.SIGUSR1,
	SIGUSR2: syscall.SIGUSR2,
}

func NewSignal(sig os.Signal) Signal {
	if v, ok := signalMap[sig]; ok {
		return v
	}

	return "unknown"
}

func (s Signal) OSSignal() os.Signal {
	if v, ok := osSignalMap[s]; ok {
		return v
	}

	return syscall.Signal(-1)
}
