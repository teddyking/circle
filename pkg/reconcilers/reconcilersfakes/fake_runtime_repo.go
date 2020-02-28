// Code generated by counterfeiter. DO NOT EDIT.
package reconcilersfakes

import (
	"sync"

	"github.com/teddyking/circle/apis/runtime/v1alpha1"
	"github.com/teddyking/circle/pkg/reconcilers"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type FakeRuntimeRepo struct {
	GetCloudFoundryStub        func(reconcile.Request) (*v1alpha1.CloudFoundry, error)
	getCloudFoundryMutex       sync.RWMutex
	getCloudFoundryArgsForCall []struct {
		arg1 reconcile.Request
	}
	getCloudFoundryReturns struct {
		result1 *v1alpha1.CloudFoundry
		result2 error
	}
	getCloudFoundryReturnsOnCall map[int]struct {
		result1 *v1alpha1.CloudFoundry
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRuntimeRepo) GetCloudFoundry(arg1 reconcile.Request) (*v1alpha1.CloudFoundry, error) {
	fake.getCloudFoundryMutex.Lock()
	ret, specificReturn := fake.getCloudFoundryReturnsOnCall[len(fake.getCloudFoundryArgsForCall)]
	fake.getCloudFoundryArgsForCall = append(fake.getCloudFoundryArgsForCall, struct {
		arg1 reconcile.Request
	}{arg1})
	fake.recordInvocation("GetCloudFoundry", []interface{}{arg1})
	fake.getCloudFoundryMutex.Unlock()
	if fake.GetCloudFoundryStub != nil {
		return fake.GetCloudFoundryStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCloudFoundryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRuntimeRepo) GetCloudFoundryCallCount() int {
	fake.getCloudFoundryMutex.RLock()
	defer fake.getCloudFoundryMutex.RUnlock()
	return len(fake.getCloudFoundryArgsForCall)
}

func (fake *FakeRuntimeRepo) GetCloudFoundryCalls(stub func(reconcile.Request) (*v1alpha1.CloudFoundry, error)) {
	fake.getCloudFoundryMutex.Lock()
	defer fake.getCloudFoundryMutex.Unlock()
	fake.GetCloudFoundryStub = stub
}

func (fake *FakeRuntimeRepo) GetCloudFoundryArgsForCall(i int) reconcile.Request {
	fake.getCloudFoundryMutex.RLock()
	defer fake.getCloudFoundryMutex.RUnlock()
	argsForCall := fake.getCloudFoundryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRuntimeRepo) GetCloudFoundryReturns(result1 *v1alpha1.CloudFoundry, result2 error) {
	fake.getCloudFoundryMutex.Lock()
	defer fake.getCloudFoundryMutex.Unlock()
	fake.GetCloudFoundryStub = nil
	fake.getCloudFoundryReturns = struct {
		result1 *v1alpha1.CloudFoundry
		result2 error
	}{result1, result2}
}

func (fake *FakeRuntimeRepo) GetCloudFoundryReturnsOnCall(i int, result1 *v1alpha1.CloudFoundry, result2 error) {
	fake.getCloudFoundryMutex.Lock()
	defer fake.getCloudFoundryMutex.Unlock()
	fake.GetCloudFoundryStub = nil
	if fake.getCloudFoundryReturnsOnCall == nil {
		fake.getCloudFoundryReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.CloudFoundry
			result2 error
		})
	}
	fake.getCloudFoundryReturnsOnCall[i] = struct {
		result1 *v1alpha1.CloudFoundry
		result2 error
	}{result1, result2}
}

func (fake *FakeRuntimeRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCloudFoundryMutex.RLock()
	defer fake.getCloudFoundryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRuntimeRepo) recordInvocation(key string, args []interface{}) {
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

var _ reconcilers.RuntimeRepo = new(FakeRuntimeRepo)