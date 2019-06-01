// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net"
	"sync"
	"time"

	ssh "github.com/demosdemon/go-ssh"
)

type FakeContext struct {
	ClientVersionStub        func() string
	clientVersionMutex       sync.RWMutex
	clientVersionArgsForCall []struct {
	}
	clientVersionReturns struct {
		result1 string
	}
	clientVersionReturnsOnCall map[int]struct {
		result1 string
	}
	DeadlineStub        func() (time.Time, bool)
	deadlineMutex       sync.RWMutex
	deadlineArgsForCall []struct {
	}
	deadlineReturns struct {
		result1 time.Time
		result2 bool
	}
	deadlineReturnsOnCall map[int]struct {
		result1 time.Time
		result2 bool
	}
	DoneStub        func() <-chan struct{}
	doneMutex       sync.RWMutex
	doneArgsForCall []struct {
	}
	doneReturns struct {
		result1 <-chan struct{}
	}
	doneReturnsOnCall map[int]struct {
		result1 <-chan struct{}
	}
	ErrStub        func() error
	errMutex       sync.RWMutex
	errArgsForCall []struct {
	}
	errReturns struct {
		result1 error
	}
	errReturnsOnCall map[int]struct {
		result1 error
	}
	LocalAddrStub        func() net.Addr
	localAddrMutex       sync.RWMutex
	localAddrArgsForCall []struct {
	}
	localAddrReturns struct {
		result1 net.Addr
	}
	localAddrReturnsOnCall map[int]struct {
		result1 net.Addr
	}
	LockStub        func()
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
	}
	LoggerStub        func() ssh.Logger
	loggerMutex       sync.RWMutex
	loggerArgsForCall []struct {
	}
	loggerReturns struct {
		result1 ssh.Logger
	}
	loggerReturnsOnCall map[int]struct {
		result1 ssh.Logger
	}
	PermissionsStub        func() *ssh.Permissions
	permissionsMutex       sync.RWMutex
	permissionsArgsForCall []struct {
	}
	permissionsReturns struct {
		result1 *ssh.Permissions
	}
	permissionsReturnsOnCall map[int]struct {
		result1 *ssh.Permissions
	}
	RemoteAddrStub        func() net.Addr
	remoteAddrMutex       sync.RWMutex
	remoteAddrArgsForCall []struct {
	}
	remoteAddrReturns struct {
		result1 net.Addr
	}
	remoteAddrReturnsOnCall map[int]struct {
		result1 net.Addr
	}
	ServerVersionStub        func() string
	serverVersionMutex       sync.RWMutex
	serverVersionArgsForCall []struct {
	}
	serverVersionReturns struct {
		result1 string
	}
	serverVersionReturnsOnCall map[int]struct {
		result1 string
	}
	SessionIDStub        func() string
	sessionIDMutex       sync.RWMutex
	sessionIDArgsForCall []struct {
	}
	sessionIDReturns struct {
		result1 string
	}
	sessionIDReturnsOnCall map[int]struct {
		result1 string
	}
	SetValueStub        func(interface{}, interface{})
	setValueMutex       sync.RWMutex
	setValueArgsForCall []struct {
		arg1 interface{}
		arg2 interface{}
	}
	UnlockStub        func()
	unlockMutex       sync.RWMutex
	unlockArgsForCall []struct {
	}
	UserStub        func() string
	userMutex       sync.RWMutex
	userArgsForCall []struct {
	}
	userReturns struct {
		result1 string
	}
	userReturnsOnCall map[int]struct {
		result1 string
	}
	ValueStub        func(interface{}) interface{}
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
		arg1 interface{}
	}
	valueReturns struct {
		result1 interface{}
	}
	valueReturnsOnCall map[int]struct {
		result1 interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContext) ClientVersion() string {
	fake.clientVersionMutex.Lock()
	ret, specificReturn := fake.clientVersionReturnsOnCall[len(fake.clientVersionArgsForCall)]
	fake.clientVersionArgsForCall = append(fake.clientVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("ClientVersion", []interface{}{})
	fake.clientVersionMutex.Unlock()
	if fake.ClientVersionStub != nil {
		return fake.ClientVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.clientVersionReturns
	return fakeReturns.result1
}

func (fake *FakeContext) ClientVersionCallCount() int {
	fake.clientVersionMutex.RLock()
	defer fake.clientVersionMutex.RUnlock()
	return len(fake.clientVersionArgsForCall)
}

func (fake *FakeContext) ClientVersionCalls(stub func() string) {
	fake.clientVersionMutex.Lock()
	defer fake.clientVersionMutex.Unlock()
	fake.ClientVersionStub = stub
}

func (fake *FakeContext) ClientVersionReturns(result1 string) {
	fake.clientVersionMutex.Lock()
	defer fake.clientVersionMutex.Unlock()
	fake.ClientVersionStub = nil
	fake.clientVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) ClientVersionReturnsOnCall(i int, result1 string) {
	fake.clientVersionMutex.Lock()
	defer fake.clientVersionMutex.Unlock()
	fake.ClientVersionStub = nil
	if fake.clientVersionReturnsOnCall == nil {
		fake.clientVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.clientVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) Deadline() (time.Time, bool) {
	fake.deadlineMutex.Lock()
	ret, specificReturn := fake.deadlineReturnsOnCall[len(fake.deadlineArgsForCall)]
	fake.deadlineArgsForCall = append(fake.deadlineArgsForCall, struct {
	}{})
	fake.recordInvocation("Deadline", []interface{}{})
	fake.deadlineMutex.Unlock()
	if fake.DeadlineStub != nil {
		return fake.DeadlineStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deadlineReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContext) DeadlineCallCount() int {
	fake.deadlineMutex.RLock()
	defer fake.deadlineMutex.RUnlock()
	return len(fake.deadlineArgsForCall)
}

func (fake *FakeContext) DeadlineCalls(stub func() (time.Time, bool)) {
	fake.deadlineMutex.Lock()
	defer fake.deadlineMutex.Unlock()
	fake.DeadlineStub = stub
}

func (fake *FakeContext) DeadlineReturns(result1 time.Time, result2 bool) {
	fake.deadlineMutex.Lock()
	defer fake.deadlineMutex.Unlock()
	fake.DeadlineStub = nil
	fake.deadlineReturns = struct {
		result1 time.Time
		result2 bool
	}{result1, result2}
}

func (fake *FakeContext) DeadlineReturnsOnCall(i int, result1 time.Time, result2 bool) {
	fake.deadlineMutex.Lock()
	defer fake.deadlineMutex.Unlock()
	fake.DeadlineStub = nil
	if fake.deadlineReturnsOnCall == nil {
		fake.deadlineReturnsOnCall = make(map[int]struct {
			result1 time.Time
			result2 bool
		})
	}
	fake.deadlineReturnsOnCall[i] = struct {
		result1 time.Time
		result2 bool
	}{result1, result2}
}

func (fake *FakeContext) Done() <-chan struct{} {
	fake.doneMutex.Lock()
	ret, specificReturn := fake.doneReturnsOnCall[len(fake.doneArgsForCall)]
	fake.doneArgsForCall = append(fake.doneArgsForCall, struct {
	}{})
	fake.recordInvocation("Done", []interface{}{})
	fake.doneMutex.Unlock()
	if fake.DoneStub != nil {
		return fake.DoneStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.doneReturns
	return fakeReturns.result1
}

func (fake *FakeContext) DoneCallCount() int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return len(fake.doneArgsForCall)
}

func (fake *FakeContext) DoneCalls(stub func() <-chan struct{}) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = stub
}

func (fake *FakeContext) DoneReturns(result1 <-chan struct{}) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	fake.doneReturns = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *FakeContext) DoneReturnsOnCall(i int, result1 <-chan struct{}) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	if fake.doneReturnsOnCall == nil {
		fake.doneReturnsOnCall = make(map[int]struct {
			result1 <-chan struct{}
		})
	}
	fake.doneReturnsOnCall[i] = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *FakeContext) Err() error {
	fake.errMutex.Lock()
	ret, specificReturn := fake.errReturnsOnCall[len(fake.errArgsForCall)]
	fake.errArgsForCall = append(fake.errArgsForCall, struct {
	}{})
	fake.recordInvocation("Err", []interface{}{})
	fake.errMutex.Unlock()
	if fake.ErrStub != nil {
		return fake.ErrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.errReturns
	return fakeReturns.result1
}

func (fake *FakeContext) ErrCallCount() int {
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	return len(fake.errArgsForCall)
}

func (fake *FakeContext) ErrCalls(stub func() error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = stub
}

func (fake *FakeContext) ErrReturns(result1 error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = nil
	fake.errReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContext) ErrReturnsOnCall(i int, result1 error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = nil
	if fake.errReturnsOnCall == nil {
		fake.errReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.errReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContext) LocalAddr() net.Addr {
	fake.localAddrMutex.Lock()
	ret, specificReturn := fake.localAddrReturnsOnCall[len(fake.localAddrArgsForCall)]
	fake.localAddrArgsForCall = append(fake.localAddrArgsForCall, struct {
	}{})
	fake.recordInvocation("LocalAddr", []interface{}{})
	fake.localAddrMutex.Unlock()
	if fake.LocalAddrStub != nil {
		return fake.LocalAddrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.localAddrReturns
	return fakeReturns.result1
}

func (fake *FakeContext) LocalAddrCallCount() int {
	fake.localAddrMutex.RLock()
	defer fake.localAddrMutex.RUnlock()
	return len(fake.localAddrArgsForCall)
}

func (fake *FakeContext) LocalAddrCalls(stub func() net.Addr) {
	fake.localAddrMutex.Lock()
	defer fake.localAddrMutex.Unlock()
	fake.LocalAddrStub = stub
}

func (fake *FakeContext) LocalAddrReturns(result1 net.Addr) {
	fake.localAddrMutex.Lock()
	defer fake.localAddrMutex.Unlock()
	fake.LocalAddrStub = nil
	fake.localAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeContext) LocalAddrReturnsOnCall(i int, result1 net.Addr) {
	fake.localAddrMutex.Lock()
	defer fake.localAddrMutex.Unlock()
	fake.LocalAddrStub = nil
	if fake.localAddrReturnsOnCall == nil {
		fake.localAddrReturnsOnCall = make(map[int]struct {
			result1 net.Addr
		})
	}
	fake.localAddrReturnsOnCall[i] = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeContext) Lock() {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
	}{})
	fake.recordInvocation("Lock", []interface{}{})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		fake.LockStub()
	}
}

func (fake *FakeContext) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeContext) LockCalls(stub func()) {
	fake.lockMutex.Lock()
	defer fake.lockMutex.Unlock()
	fake.LockStub = stub
}

func (fake *FakeContext) Logger() ssh.Logger {
	fake.loggerMutex.Lock()
	ret, specificReturn := fake.loggerReturnsOnCall[len(fake.loggerArgsForCall)]
	fake.loggerArgsForCall = append(fake.loggerArgsForCall, struct {
	}{})
	fake.recordInvocation("Logger", []interface{}{})
	fake.loggerMutex.Unlock()
	if fake.LoggerStub != nil {
		return fake.LoggerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.loggerReturns
	return fakeReturns.result1
}

func (fake *FakeContext) LoggerCallCount() int {
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	return len(fake.loggerArgsForCall)
}

func (fake *FakeContext) LoggerCalls(stub func() ssh.Logger) {
	fake.loggerMutex.Lock()
	defer fake.loggerMutex.Unlock()
	fake.LoggerStub = stub
}

func (fake *FakeContext) LoggerReturns(result1 ssh.Logger) {
	fake.loggerMutex.Lock()
	defer fake.loggerMutex.Unlock()
	fake.LoggerStub = nil
	fake.loggerReturns = struct {
		result1 ssh.Logger
	}{result1}
}

func (fake *FakeContext) LoggerReturnsOnCall(i int, result1 ssh.Logger) {
	fake.loggerMutex.Lock()
	defer fake.loggerMutex.Unlock()
	fake.LoggerStub = nil
	if fake.loggerReturnsOnCall == nil {
		fake.loggerReturnsOnCall = make(map[int]struct {
			result1 ssh.Logger
		})
	}
	fake.loggerReturnsOnCall[i] = struct {
		result1 ssh.Logger
	}{result1}
}

func (fake *FakeContext) Permissions() *ssh.Permissions {
	fake.permissionsMutex.Lock()
	ret, specificReturn := fake.permissionsReturnsOnCall[len(fake.permissionsArgsForCall)]
	fake.permissionsArgsForCall = append(fake.permissionsArgsForCall, struct {
	}{})
	fake.recordInvocation("Permissions", []interface{}{})
	fake.permissionsMutex.Unlock()
	if fake.PermissionsStub != nil {
		return fake.PermissionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.permissionsReturns
	return fakeReturns.result1
}

func (fake *FakeContext) PermissionsCallCount() int {
	fake.permissionsMutex.RLock()
	defer fake.permissionsMutex.RUnlock()
	return len(fake.permissionsArgsForCall)
}

func (fake *FakeContext) PermissionsCalls(stub func() *ssh.Permissions) {
	fake.permissionsMutex.Lock()
	defer fake.permissionsMutex.Unlock()
	fake.PermissionsStub = stub
}

func (fake *FakeContext) PermissionsReturns(result1 *ssh.Permissions) {
	fake.permissionsMutex.Lock()
	defer fake.permissionsMutex.Unlock()
	fake.PermissionsStub = nil
	fake.permissionsReturns = struct {
		result1 *ssh.Permissions
	}{result1}
}

func (fake *FakeContext) PermissionsReturnsOnCall(i int, result1 *ssh.Permissions) {
	fake.permissionsMutex.Lock()
	defer fake.permissionsMutex.Unlock()
	fake.PermissionsStub = nil
	if fake.permissionsReturnsOnCall == nil {
		fake.permissionsReturnsOnCall = make(map[int]struct {
			result1 *ssh.Permissions
		})
	}
	fake.permissionsReturnsOnCall[i] = struct {
		result1 *ssh.Permissions
	}{result1}
}

func (fake *FakeContext) RemoteAddr() net.Addr {
	fake.remoteAddrMutex.Lock()
	ret, specificReturn := fake.remoteAddrReturnsOnCall[len(fake.remoteAddrArgsForCall)]
	fake.remoteAddrArgsForCall = append(fake.remoteAddrArgsForCall, struct {
	}{})
	fake.recordInvocation("RemoteAddr", []interface{}{})
	fake.remoteAddrMutex.Unlock()
	if fake.RemoteAddrStub != nil {
		return fake.RemoteAddrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.remoteAddrReturns
	return fakeReturns.result1
}

func (fake *FakeContext) RemoteAddrCallCount() int {
	fake.remoteAddrMutex.RLock()
	defer fake.remoteAddrMutex.RUnlock()
	return len(fake.remoteAddrArgsForCall)
}

func (fake *FakeContext) RemoteAddrCalls(stub func() net.Addr) {
	fake.remoteAddrMutex.Lock()
	defer fake.remoteAddrMutex.Unlock()
	fake.RemoteAddrStub = stub
}

func (fake *FakeContext) RemoteAddrReturns(result1 net.Addr) {
	fake.remoteAddrMutex.Lock()
	defer fake.remoteAddrMutex.Unlock()
	fake.RemoteAddrStub = nil
	fake.remoteAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeContext) RemoteAddrReturnsOnCall(i int, result1 net.Addr) {
	fake.remoteAddrMutex.Lock()
	defer fake.remoteAddrMutex.Unlock()
	fake.RemoteAddrStub = nil
	if fake.remoteAddrReturnsOnCall == nil {
		fake.remoteAddrReturnsOnCall = make(map[int]struct {
			result1 net.Addr
		})
	}
	fake.remoteAddrReturnsOnCall[i] = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeContext) ServerVersion() string {
	fake.serverVersionMutex.Lock()
	ret, specificReturn := fake.serverVersionReturnsOnCall[len(fake.serverVersionArgsForCall)]
	fake.serverVersionArgsForCall = append(fake.serverVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("ServerVersion", []interface{}{})
	fake.serverVersionMutex.Unlock()
	if fake.ServerVersionStub != nil {
		return fake.ServerVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.serverVersionReturns
	return fakeReturns.result1
}

func (fake *FakeContext) ServerVersionCallCount() int {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	return len(fake.serverVersionArgsForCall)
}

func (fake *FakeContext) ServerVersionCalls(stub func() string) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = stub
}

func (fake *FakeContext) ServerVersionReturns(result1 string) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	fake.serverVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) ServerVersionReturnsOnCall(i int, result1 string) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	if fake.serverVersionReturnsOnCall == nil {
		fake.serverVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.serverVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) SessionID() string {
	fake.sessionIDMutex.Lock()
	ret, specificReturn := fake.sessionIDReturnsOnCall[len(fake.sessionIDArgsForCall)]
	fake.sessionIDArgsForCall = append(fake.sessionIDArgsForCall, struct {
	}{})
	fake.recordInvocation("SessionID", []interface{}{})
	fake.sessionIDMutex.Unlock()
	if fake.SessionIDStub != nil {
		return fake.SessionIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sessionIDReturns
	return fakeReturns.result1
}

func (fake *FakeContext) SessionIDCallCount() int {
	fake.sessionIDMutex.RLock()
	defer fake.sessionIDMutex.RUnlock()
	return len(fake.sessionIDArgsForCall)
}

func (fake *FakeContext) SessionIDCalls(stub func() string) {
	fake.sessionIDMutex.Lock()
	defer fake.sessionIDMutex.Unlock()
	fake.SessionIDStub = stub
}

func (fake *FakeContext) SessionIDReturns(result1 string) {
	fake.sessionIDMutex.Lock()
	defer fake.sessionIDMutex.Unlock()
	fake.SessionIDStub = nil
	fake.sessionIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) SessionIDReturnsOnCall(i int, result1 string) {
	fake.sessionIDMutex.Lock()
	defer fake.sessionIDMutex.Unlock()
	fake.SessionIDStub = nil
	if fake.sessionIDReturnsOnCall == nil {
		fake.sessionIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sessionIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) SetValue(arg1 interface{}, arg2 interface{}) {
	fake.setValueMutex.Lock()
	fake.setValueArgsForCall = append(fake.setValueArgsForCall, struct {
		arg1 interface{}
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("SetValue", []interface{}{arg1, arg2})
	fake.setValueMutex.Unlock()
	if fake.SetValueStub != nil {
		fake.SetValueStub(arg1, arg2)
	}
}

func (fake *FakeContext) SetValueCallCount() int {
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	return len(fake.setValueArgsForCall)
}

func (fake *FakeContext) SetValueCalls(stub func(interface{}, interface{})) {
	fake.setValueMutex.Lock()
	defer fake.setValueMutex.Unlock()
	fake.SetValueStub = stub
}

func (fake *FakeContext) SetValueArgsForCall(i int) (interface{}, interface{}) {
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	argsForCall := fake.setValueArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContext) Unlock() {
	fake.unlockMutex.Lock()
	fake.unlockArgsForCall = append(fake.unlockArgsForCall, struct {
	}{})
	fake.recordInvocation("Unlock", []interface{}{})
	fake.unlockMutex.Unlock()
	if fake.UnlockStub != nil {
		fake.UnlockStub()
	}
}

func (fake *FakeContext) UnlockCallCount() int {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return len(fake.unlockArgsForCall)
}

func (fake *FakeContext) UnlockCalls(stub func()) {
	fake.unlockMutex.Lock()
	defer fake.unlockMutex.Unlock()
	fake.UnlockStub = stub
}

func (fake *FakeContext) User() string {
	fake.userMutex.Lock()
	ret, specificReturn := fake.userReturnsOnCall[len(fake.userArgsForCall)]
	fake.userArgsForCall = append(fake.userArgsForCall, struct {
	}{})
	fake.recordInvocation("User", []interface{}{})
	fake.userMutex.Unlock()
	if fake.UserStub != nil {
		return fake.UserStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.userReturns
	return fakeReturns.result1
}

func (fake *FakeContext) UserCallCount() int {
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	return len(fake.userArgsForCall)
}

func (fake *FakeContext) UserCalls(stub func() string) {
	fake.userMutex.Lock()
	defer fake.userMutex.Unlock()
	fake.UserStub = stub
}

func (fake *FakeContext) UserReturns(result1 string) {
	fake.userMutex.Lock()
	defer fake.userMutex.Unlock()
	fake.UserStub = nil
	fake.userReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) UserReturnsOnCall(i int, result1 string) {
	fake.userMutex.Lock()
	defer fake.userMutex.Unlock()
	fake.UserStub = nil
	if fake.userReturnsOnCall == nil {
		fake.userReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.userReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeContext) Value(arg1 interface{}) interface{} {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Value", []interface{}{arg1})
	fake.valueMutex.Unlock()
	if fake.ValueStub != nil {
		return fake.ValueStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.valueReturns
	return fakeReturns.result1
}

func (fake *FakeContext) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeContext) ValueCalls(stub func(interface{}) interface{}) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeContext) ValueArgsForCall(i int) interface{} {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	argsForCall := fake.valueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContext) ValueReturns(result1 interface{}) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeContext) ValueReturnsOnCall(i int, result1 interface{}) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clientVersionMutex.RLock()
	defer fake.clientVersionMutex.RUnlock()
	fake.deadlineMutex.RLock()
	defer fake.deadlineMutex.RUnlock()
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	fake.localAddrMutex.RLock()
	defer fake.localAddrMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	fake.permissionsMutex.RLock()
	defer fake.permissionsMutex.RUnlock()
	fake.remoteAddrMutex.RLock()
	defer fake.remoteAddrMutex.RUnlock()
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	fake.sessionIDMutex.RLock()
	defer fake.sessionIDMutex.RUnlock()
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContext) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ ssh.Context = new(FakeContext)
