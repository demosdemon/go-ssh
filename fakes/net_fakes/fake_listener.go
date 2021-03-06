// Code generated by counterfeiter. DO NOT EDIT.
package net_fakes

import (
	"net"
	"sync"
)

type FakeListener struct {
	AcceptStub        func() (net.Conn, error)
	acceptMutex       sync.RWMutex
	acceptArgsForCall []struct {
	}
	acceptReturns struct {
		result1 net.Conn
		result2 error
	}
	acceptReturnsOnCall map[int]struct {
		result1 net.Conn
		result2 error
	}
	AddrStub        func() net.Addr
	addrMutex       sync.RWMutex
	addrArgsForCall []struct {
	}
	addrReturns struct {
		result1 net.Addr
	}
	addrReturnsOnCall map[int]struct {
		result1 net.Addr
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeListener) Accept() (net.Conn, error) {
	fake.acceptMutex.Lock()
	ret, specificReturn := fake.acceptReturnsOnCall[len(fake.acceptArgsForCall)]
	fake.acceptArgsForCall = append(fake.acceptArgsForCall, struct {
	}{})
	fake.recordInvocation("Accept", []interface{}{})
	fake.acceptMutex.Unlock()
	if fake.AcceptStub != nil {
		return fake.AcceptStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.acceptReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeListener) AcceptCallCount() int {
	fake.acceptMutex.RLock()
	defer fake.acceptMutex.RUnlock()
	return len(fake.acceptArgsForCall)
}

func (fake *FakeListener) AcceptCalls(stub func() (net.Conn, error)) {
	fake.acceptMutex.Lock()
	defer fake.acceptMutex.Unlock()
	fake.AcceptStub = stub
}

func (fake *FakeListener) AcceptReturns(result1 net.Conn, result2 error) {
	fake.acceptMutex.Lock()
	defer fake.acceptMutex.Unlock()
	fake.AcceptStub = nil
	fake.acceptReturns = struct {
		result1 net.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeListener) AcceptReturnsOnCall(i int, result1 net.Conn, result2 error) {
	fake.acceptMutex.Lock()
	defer fake.acceptMutex.Unlock()
	fake.AcceptStub = nil
	if fake.acceptReturnsOnCall == nil {
		fake.acceptReturnsOnCall = make(map[int]struct {
			result1 net.Conn
			result2 error
		})
	}
	fake.acceptReturnsOnCall[i] = struct {
		result1 net.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeListener) Addr() net.Addr {
	fake.addrMutex.Lock()
	ret, specificReturn := fake.addrReturnsOnCall[len(fake.addrArgsForCall)]
	fake.addrArgsForCall = append(fake.addrArgsForCall, struct {
	}{})
	fake.recordInvocation("Addr", []interface{}{})
	fake.addrMutex.Unlock()
	if fake.AddrStub != nil {
		return fake.AddrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addrReturns
	return fakeReturns.result1
}

func (fake *FakeListener) AddrCallCount() int {
	fake.addrMutex.RLock()
	defer fake.addrMutex.RUnlock()
	return len(fake.addrArgsForCall)
}

func (fake *FakeListener) AddrCalls(stub func() net.Addr) {
	fake.addrMutex.Lock()
	defer fake.addrMutex.Unlock()
	fake.AddrStub = stub
}

func (fake *FakeListener) AddrReturns(result1 net.Addr) {
	fake.addrMutex.Lock()
	defer fake.addrMutex.Unlock()
	fake.AddrStub = nil
	fake.addrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeListener) AddrReturnsOnCall(i int, result1 net.Addr) {
	fake.addrMutex.Lock()
	defer fake.addrMutex.Unlock()
	fake.AddrStub = nil
	if fake.addrReturnsOnCall == nil {
		fake.addrReturnsOnCall = make(map[int]struct {
			result1 net.Addr
		})
	}
	fake.addrReturnsOnCall[i] = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeListener) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeListener) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeListener) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeListener) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeListener) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeListener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acceptMutex.RLock()
	defer fake.acceptMutex.RUnlock()
	fake.addrMutex.RLock()
	defer fake.addrMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeListener) recordInvocation(key string, args []interface{}) {
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

var _ net.Listener = new(FakeListener)
