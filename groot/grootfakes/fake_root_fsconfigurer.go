// This file was generated by counterfeiter
package grootfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/groot"
	specsv1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type FakeRootFSConfigurer struct {
	ConfigureStub        func(rootFSPath string, baseImage *specsv1.Image) error
	configureMutex       sync.RWMutex
	configureArgsForCall []struct {
		rootFSPath string
		baseImage  *specsv1.Image
	}
	configureReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRootFSConfigurer) Configure(rootFSPath string, baseImage *specsv1.Image) error {
	fake.configureMutex.Lock()
	fake.configureArgsForCall = append(fake.configureArgsForCall, struct {
		rootFSPath string
		baseImage  *specsv1.Image
	}{rootFSPath, baseImage})
	fake.recordInvocation("Configure", []interface{}{rootFSPath, baseImage})
	fake.configureMutex.Unlock()
	if fake.ConfigureStub != nil {
		return fake.ConfigureStub(rootFSPath, baseImage)
	} else {
		return fake.configureReturns.result1
	}
}

func (fake *FakeRootFSConfigurer) ConfigureCallCount() int {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return len(fake.configureArgsForCall)
}

func (fake *FakeRootFSConfigurer) ConfigureArgsForCall(i int) (string, *specsv1.Image) {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return fake.configureArgsForCall[i].rootFSPath, fake.configureArgsForCall[i].baseImage
}

func (fake *FakeRootFSConfigurer) ConfigureReturns(result1 error) {
	fake.ConfigureStub = nil
	fake.configureReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRootFSConfigurer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRootFSConfigurer) recordInvocation(key string, args []interface{}) {
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

var _ groot.RootFSConfigurer = new(FakeRootFSConfigurer)
