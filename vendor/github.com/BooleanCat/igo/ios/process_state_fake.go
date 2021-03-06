// This file was generated by counterfeiter
// counterfeiter -o ios/process_state_fake.go --fake-name ProcessStateFake ios/process_state.go ProcessState

package ios

import (
	"sync"
	"time"
)

//ProcessStateFake ...
type ProcessStateFake struct {
	ExitedStub        func() bool
	exitedMutex       sync.RWMutex
	exitedArgsForCall []struct{}
	exitedReturns     struct {
		result1 bool
	}
	PidStub        func() int
	pidMutex       sync.RWMutex
	pidArgsForCall []struct{}
	pidReturns     struct {
		result1 int
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct{}
	stringReturns     struct {
		result1 string
	}
	SuccessStub        func() bool
	successMutex       sync.RWMutex
	successArgsForCall []struct{}
	successReturns     struct {
		result1 bool
	}
	SysStub        func() interface{}
	sysMutex       sync.RWMutex
	sysArgsForCall []struct{}
	sysReturns     struct {
		result1 interface{}
	}
	SysUsageStub        func() interface{}
	sysUsageMutex       sync.RWMutex
	sysUsageArgsForCall []struct{}
	sysUsageReturns     struct {
		result1 interface{}
	}
	SystemTimeStub        func() time.Duration
	systemTimeMutex       sync.RWMutex
	systemTimeArgsForCall []struct{}
	systemTimeReturns     struct {
		result1 time.Duration
	}
	UserTimeStub        func() time.Duration
	userTimeMutex       sync.RWMutex
	userTimeArgsForCall []struct{}
	userTimeReturns     struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//NewProcessStateFake is the preferred way to initialise a ProcessStateFake
func NewProcessStateFake() *ProcessStateFake {
	return new(ProcessStateFake)
}

//Exited ...
func (fake *ProcessStateFake) Exited() bool {
	fake.exitedMutex.Lock()
	fake.exitedArgsForCall = append(fake.exitedArgsForCall, struct{}{})
	fake.recordInvocation("Exited", []interface{}{})
	fake.exitedMutex.Unlock()
	if fake.ExitedStub != nil {
		return fake.ExitedStub()
	}
	return fake.exitedReturns.result1
}

//ExitedCallCount ...
func (fake *ProcessStateFake) ExitedCallCount() int {
	fake.exitedMutex.RLock()
	defer fake.exitedMutex.RUnlock()
	return len(fake.exitedArgsForCall)
}

//ExitedReturns ...
func (fake *ProcessStateFake) ExitedReturns(result1 bool) {
	fake.ExitedStub = nil
	fake.exitedReturns = struct {
		result1 bool
	}{result1}
}

//Pid ...
func (fake *ProcessStateFake) Pid() int {
	fake.pidMutex.Lock()
	fake.pidArgsForCall = append(fake.pidArgsForCall, struct{}{})
	fake.recordInvocation("Pid", []interface{}{})
	fake.pidMutex.Unlock()
	if fake.PidStub != nil {
		return fake.PidStub()
	}
	return fake.pidReturns.result1
}

//PidCallCount ...
func (fake *ProcessStateFake) PidCallCount() int {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return len(fake.pidArgsForCall)
}

//PidReturns ...
func (fake *ProcessStateFake) PidReturns(result1 int) {
	fake.PidStub = nil
	fake.pidReturns = struct {
		result1 int
	}{result1}
}

//String ...
func (fake *ProcessStateFake) String() string {
	fake.stringMutex.Lock()
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct{}{})
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if fake.StringStub != nil {
		return fake.StringStub()
	}
	return fake.stringReturns.result1
}

//StringCallCount ...
func (fake *ProcessStateFake) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

//StringReturns ...
func (fake *ProcessStateFake) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

//Success ...
func (fake *ProcessStateFake) Success() bool {
	fake.successMutex.Lock()
	fake.successArgsForCall = append(fake.successArgsForCall, struct{}{})
	fake.recordInvocation("Success", []interface{}{})
	fake.successMutex.Unlock()
	if fake.SuccessStub != nil {
		return fake.SuccessStub()
	}
	return fake.successReturns.result1
}

//SuccessCallCount ...
func (fake *ProcessStateFake) SuccessCallCount() int {
	fake.successMutex.RLock()
	defer fake.successMutex.RUnlock()
	return len(fake.successArgsForCall)
}

//SuccessReturns ...
func (fake *ProcessStateFake) SuccessReturns(result1 bool) {
	fake.SuccessStub = nil
	fake.successReturns = struct {
		result1 bool
	}{result1}
}

//Sys ...
func (fake *ProcessStateFake) Sys() interface{} {
	fake.sysMutex.Lock()
	fake.sysArgsForCall = append(fake.sysArgsForCall, struct{}{})
	fake.recordInvocation("Sys", []interface{}{})
	fake.sysMutex.Unlock()
	if fake.SysStub != nil {
		return fake.SysStub()
	}
	return fake.sysReturns.result1
}

//SysCallCount ...
func (fake *ProcessStateFake) SysCallCount() int {
	fake.sysMutex.RLock()
	defer fake.sysMutex.RUnlock()
	return len(fake.sysArgsForCall)
}

//SysReturns ...
func (fake *ProcessStateFake) SysReturns(result1 interface{}) {
	fake.SysStub = nil
	fake.sysReturns = struct {
		result1 interface{}
	}{result1}
}

//SysUsage ...
func (fake *ProcessStateFake) SysUsage() interface{} {
	fake.sysUsageMutex.Lock()
	fake.sysUsageArgsForCall = append(fake.sysUsageArgsForCall, struct{}{})
	fake.recordInvocation("SysUsage", []interface{}{})
	fake.sysUsageMutex.Unlock()
	if fake.SysUsageStub != nil {
		return fake.SysUsageStub()
	}
	return fake.sysUsageReturns.result1
}

//SysUsageCallCount ...
func (fake *ProcessStateFake) SysUsageCallCount() int {
	fake.sysUsageMutex.RLock()
	defer fake.sysUsageMutex.RUnlock()
	return len(fake.sysUsageArgsForCall)
}

//SysUsageReturns ...
func (fake *ProcessStateFake) SysUsageReturns(result1 interface{}) {
	fake.SysUsageStub = nil
	fake.sysUsageReturns = struct {
		result1 interface{}
	}{result1}
}

//SystemTime ...
func (fake *ProcessStateFake) SystemTime() time.Duration {
	fake.systemTimeMutex.Lock()
	fake.systemTimeArgsForCall = append(fake.systemTimeArgsForCall, struct{}{})
	fake.recordInvocation("SystemTime", []interface{}{})
	fake.systemTimeMutex.Unlock()
	if fake.SystemTimeStub != nil {
		return fake.SystemTimeStub()
	}
	return fake.systemTimeReturns.result1
}

//SystemTimeCallCount ...
func (fake *ProcessStateFake) SystemTimeCallCount() int {
	fake.systemTimeMutex.RLock()
	defer fake.systemTimeMutex.RUnlock()
	return len(fake.systemTimeArgsForCall)
}

//SystemTimeReturns ...
func (fake *ProcessStateFake) SystemTimeReturns(result1 time.Duration) {
	fake.SystemTimeStub = nil
	fake.systemTimeReturns = struct {
		result1 time.Duration
	}{result1}
}

//UserTime ...
func (fake *ProcessStateFake) UserTime() time.Duration {
	fake.userTimeMutex.Lock()
	fake.userTimeArgsForCall = append(fake.userTimeArgsForCall, struct{}{})
	fake.recordInvocation("UserTime", []interface{}{})
	fake.userTimeMutex.Unlock()
	if fake.UserTimeStub != nil {
		return fake.UserTimeStub()
	}
	return fake.userTimeReturns.result1
}

//UserTimeCallCount ...
func (fake *ProcessStateFake) UserTimeCallCount() int {
	fake.userTimeMutex.RLock()
	defer fake.userTimeMutex.RUnlock()
	return len(fake.userTimeArgsForCall)
}

//UserTimeReturns ...
func (fake *ProcessStateFake) UserTimeReturns(result1 time.Duration) {
	fake.UserTimeStub = nil
	fake.userTimeReturns = struct {
		result1 time.Duration
	}{result1}
}

//Invocations ...
func (fake *ProcessStateFake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.exitedMutex.RLock()
	defer fake.exitedMutex.RUnlock()
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	fake.successMutex.RLock()
	defer fake.successMutex.RUnlock()
	fake.sysMutex.RLock()
	defer fake.sysMutex.RUnlock()
	fake.sysUsageMutex.RLock()
	defer fake.sysUsageMutex.RUnlock()
	fake.systemTimeMutex.RLock()
	defer fake.systemTimeMutex.RUnlock()
	fake.userTimeMutex.RLock()
	defer fake.userTimeMutex.RUnlock()
	return fake.invocations
}

func (fake *ProcessStateFake) recordInvocation(key string, args []interface{}) {
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

var _ ProcessState = new(ProcessStateFake)
