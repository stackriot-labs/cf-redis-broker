// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/cf-redis-broker/redis"
)

type FakeInstanceInformer struct {
	InstancePidStub        func(string) (int, error)
	instancePidMutex       sync.RWMutex
	instancePidArgsForCall []struct {
		arg1 string
	}
	instancePidReturns struct {
		result1 int
		result2 error
	}
	instancePidReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstanceInformer) InstancePid(arg1 string) (int, error) {
	fake.instancePidMutex.Lock()
	ret, specificReturn := fake.instancePidReturnsOnCall[len(fake.instancePidArgsForCall)]
	fake.instancePidArgsForCall = append(fake.instancePidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InstancePid", []interface{}{arg1})
	fake.instancePidMutex.Unlock()
	if fake.InstancePidStub != nil {
		return fake.InstancePidStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instancePidReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstanceInformer) InstancePidCallCount() int {
	fake.instancePidMutex.RLock()
	defer fake.instancePidMutex.RUnlock()
	return len(fake.instancePidArgsForCall)
}

func (fake *FakeInstanceInformer) InstancePidCalls(stub func(string) (int, error)) {
	fake.instancePidMutex.Lock()
	defer fake.instancePidMutex.Unlock()
	fake.InstancePidStub = stub
}

func (fake *FakeInstanceInformer) InstancePidArgsForCall(i int) string {
	fake.instancePidMutex.RLock()
	defer fake.instancePidMutex.RUnlock()
	argsForCall := fake.instancePidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstanceInformer) InstancePidReturns(result1 int, result2 error) {
	fake.instancePidMutex.Lock()
	defer fake.instancePidMutex.Unlock()
	fake.InstancePidStub = nil
	fake.instancePidReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceInformer) InstancePidReturnsOnCall(i int, result1 int, result2 error) {
	fake.instancePidMutex.Lock()
	defer fake.instancePidMutex.Unlock()
	fake.InstancePidStub = nil
	if fake.instancePidReturnsOnCall == nil {
		fake.instancePidReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.instancePidReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceInformer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.instancePidMutex.RLock()
	defer fake.instancePidMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstanceInformer) recordInvocation(key string, args []interface{}) {
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

var _ redis.InstanceInformer = new(FakeInstanceInformer)
