// This file was generated by counterfeiter
package managerfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/store/manager"
	"code.cloudfoundry.org/lager"
)

type FakeStoreDriver struct {
	ConfigureStoreStub        func(logger lager.Logger, storePath string, ownerUID, ownerGID int) error
	configureStoreMutex       sync.RWMutex
	configureStoreArgsForCall []struct {
		logger    lager.Logger
		storePath string
		ownerUID  int
		ownerGID  int
	}
	configureStoreReturns struct {
		result1 error
	}
	ValidateFileSystemStub        func(logger lager.Logger, path string) error
	validateFileSystemMutex       sync.RWMutex
	validateFileSystemArgsForCall []struct {
		logger lager.Logger
		path   string
	}
	validateFileSystemReturns struct {
		result1 error
	}
	InitFilesystemStub        func(logger lager.Logger, filesystemPath, storePath string) error
	initFilesystemMutex       sync.RWMutex
	initFilesystemArgsForCall []struct {
		logger         lager.Logger
		filesystemPath string
		storePath      string
	}
	initFilesystemReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStoreDriver) ConfigureStore(logger lager.Logger, storePath string, ownerUID int, ownerGID int) error {
	fake.configureStoreMutex.Lock()
	fake.configureStoreArgsForCall = append(fake.configureStoreArgsForCall, struct {
		logger    lager.Logger
		storePath string
		ownerUID  int
		ownerGID  int
	}{logger, storePath, ownerUID, ownerGID})
	fake.recordInvocation("ConfigureStore", []interface{}{logger, storePath, ownerUID, ownerGID})
	fake.configureStoreMutex.Unlock()
	if fake.ConfigureStoreStub != nil {
		return fake.ConfigureStoreStub(logger, storePath, ownerUID, ownerGID)
	} else {
		return fake.configureStoreReturns.result1
	}
}

func (fake *FakeStoreDriver) ConfigureStoreCallCount() int {
	fake.configureStoreMutex.RLock()
	defer fake.configureStoreMutex.RUnlock()
	return len(fake.configureStoreArgsForCall)
}

func (fake *FakeStoreDriver) ConfigureStoreArgsForCall(i int) (lager.Logger, string, int, int) {
	fake.configureStoreMutex.RLock()
	defer fake.configureStoreMutex.RUnlock()
	return fake.configureStoreArgsForCall[i].logger, fake.configureStoreArgsForCall[i].storePath, fake.configureStoreArgsForCall[i].ownerUID, fake.configureStoreArgsForCall[i].ownerGID
}

func (fake *FakeStoreDriver) ConfigureStoreReturns(result1 error) {
	fake.ConfigureStoreStub = nil
	fake.configureStoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) ValidateFileSystem(logger lager.Logger, path string) error {
	fake.validateFileSystemMutex.Lock()
	fake.validateFileSystemArgsForCall = append(fake.validateFileSystemArgsForCall, struct {
		logger lager.Logger
		path   string
	}{logger, path})
	fake.recordInvocation("ValidateFileSystem", []interface{}{logger, path})
	fake.validateFileSystemMutex.Unlock()
	if fake.ValidateFileSystemStub != nil {
		return fake.ValidateFileSystemStub(logger, path)
	} else {
		return fake.validateFileSystemReturns.result1
	}
}

func (fake *FakeStoreDriver) ValidateFileSystemCallCount() int {
	fake.validateFileSystemMutex.RLock()
	defer fake.validateFileSystemMutex.RUnlock()
	return len(fake.validateFileSystemArgsForCall)
}

func (fake *FakeStoreDriver) ValidateFileSystemArgsForCall(i int) (lager.Logger, string) {
	fake.validateFileSystemMutex.RLock()
	defer fake.validateFileSystemMutex.RUnlock()
	return fake.validateFileSystemArgsForCall[i].logger, fake.validateFileSystemArgsForCall[i].path
}

func (fake *FakeStoreDriver) ValidateFileSystemReturns(result1 error) {
	fake.ValidateFileSystemStub = nil
	fake.validateFileSystemReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) InitFilesystem(logger lager.Logger, filesystemPath string, storePath string) error {
	fake.initFilesystemMutex.Lock()
	fake.initFilesystemArgsForCall = append(fake.initFilesystemArgsForCall, struct {
		logger         lager.Logger
		filesystemPath string
		storePath      string
	}{logger, filesystemPath, storePath})
	fake.recordInvocation("InitFilesystem", []interface{}{logger, filesystemPath, storePath})
	fake.initFilesystemMutex.Unlock()
	if fake.InitFilesystemStub != nil {
		return fake.InitFilesystemStub(logger, filesystemPath, storePath)
	} else {
		return fake.initFilesystemReturns.result1
	}
}

func (fake *FakeStoreDriver) InitFilesystemCallCount() int {
	fake.initFilesystemMutex.RLock()
	defer fake.initFilesystemMutex.RUnlock()
	return len(fake.initFilesystemArgsForCall)
}

func (fake *FakeStoreDriver) InitFilesystemArgsForCall(i int) (lager.Logger, string, string) {
	fake.initFilesystemMutex.RLock()
	defer fake.initFilesystemMutex.RUnlock()
	return fake.initFilesystemArgsForCall[i].logger, fake.initFilesystemArgsForCall[i].filesystemPath, fake.initFilesystemArgsForCall[i].storePath
}

func (fake *FakeStoreDriver) InitFilesystemReturns(result1 error) {
	fake.InitFilesystemStub = nil
	fake.initFilesystemReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configureStoreMutex.RLock()
	defer fake.configureStoreMutex.RUnlock()
	fake.validateFileSystemMutex.RLock()
	defer fake.validateFileSystemMutex.RUnlock()
	fake.initFilesystemMutex.RLock()
	defer fake.initFilesystemMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStoreDriver) recordInvocation(key string, args []interface{}) {
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

var _ manager.StoreDriver = new(FakeStoreDriver)
