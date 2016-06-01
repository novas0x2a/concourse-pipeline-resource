// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeGardenWorkerDB struct {
	CreateContainerStub        func(db.Container, time.Duration, time.Duration) (db.SavedContainer, error)
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		arg1 db.Container
		arg2 time.Duration
		arg3 time.Duration
	}
	createContainerReturns struct {
		result1 db.SavedContainer
		result2 error
	}
	UpdateExpiresAtOnContainerStub        func(handle string, ttl time.Duration) error
	updateExpiresAtOnContainerMutex       sync.RWMutex
	updateExpiresAtOnContainerArgsForCall []struct {
		handle string
		ttl    time.Duration
	}
	updateExpiresAtOnContainerReturns struct {
		result1 error
	}
	InsertVolumeStub        func(db.Volume) error
	insertVolumeMutex       sync.RWMutex
	insertVolumeArgsForCall []struct {
		arg1 db.Volume
	}
	insertVolumeReturns struct {
		result1 error
	}
	SetVolumeTTLStub        func(string, time.Duration) error
	setVolumeTTLMutex       sync.RWMutex
	setVolumeTTLArgsForCall []struct {
		arg1 string
		arg2 time.Duration
	}
	setVolumeTTLReturns struct {
		result1 error
	}
	GetVolumeTTLStub        func(string) (time.Duration, bool, error)
	getVolumeTTLMutex       sync.RWMutex
	getVolumeTTLArgsForCall []struct {
		arg1 string
	}
	getVolumeTTLReturns struct {
		result1 time.Duration
		result2 bool
		result3 error
	}
	GetVolumesByIdentifierStub        func(db.VolumeIdentifier) ([]db.SavedVolume, error)
	getVolumesByIdentifierMutex       sync.RWMutex
	getVolumesByIdentifierArgsForCall []struct {
		arg1 db.VolumeIdentifier
	}
	getVolumesByIdentifierReturns struct {
		result1 []db.SavedVolume
		result2 error
	}
	ReapVolumeStub        func(string) error
	reapVolumeMutex       sync.RWMutex
	reapVolumeArgsForCall []struct {
		arg1 string
	}
	reapVolumeReturns struct {
		result1 error
	}
}

func (fake *FakeGardenWorkerDB) CreateContainer(arg1 db.Container, arg2 time.Duration, arg3 time.Duration) (db.SavedContainer, error) {
	fake.createContainerMutex.Lock()
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		arg1 db.Container
		arg2 time.Duration
		arg3 time.Duration
	}{arg1, arg2, arg3})
	fake.createContainerMutex.Unlock()
	if fake.CreateContainerStub != nil {
		return fake.CreateContainerStub(arg1, arg2, arg3)
	} else {
		return fake.createContainerReturns.result1, fake.createContainerReturns.result2
	}
}

func (fake *FakeGardenWorkerDB) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) CreateContainerArgsForCall(i int) (db.Container, time.Duration, time.Duration) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return fake.createContainerArgsForCall[i].arg1, fake.createContainerArgsForCall[i].arg2, fake.createContainerArgsForCall[i].arg3
}

func (fake *FakeGardenWorkerDB) CreateContainerReturns(result1 db.SavedContainer, result2 error) {
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainer(handle string, ttl time.Duration) error {
	fake.updateExpiresAtOnContainerMutex.Lock()
	fake.updateExpiresAtOnContainerArgsForCall = append(fake.updateExpiresAtOnContainerArgsForCall, struct {
		handle string
		ttl    time.Duration
	}{handle, ttl})
	fake.updateExpiresAtOnContainerMutex.Unlock()
	if fake.UpdateExpiresAtOnContainerStub != nil {
		return fake.UpdateExpiresAtOnContainerStub(handle, ttl)
	} else {
		return fake.updateExpiresAtOnContainerReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerCallCount() int {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return len(fake.updateExpiresAtOnContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerArgsForCall(i int) (string, time.Duration) {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return fake.updateExpiresAtOnContainerArgsForCall[i].handle, fake.updateExpiresAtOnContainerArgsForCall[i].ttl
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerReturns(result1 error) {
	fake.UpdateExpiresAtOnContainerStub = nil
	fake.updateExpiresAtOnContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGardenWorkerDB) InsertVolume(arg1 db.Volume) error {
	fake.insertVolumeMutex.Lock()
	fake.insertVolumeArgsForCall = append(fake.insertVolumeArgsForCall, struct {
		arg1 db.Volume
	}{arg1})
	fake.insertVolumeMutex.Unlock()
	if fake.InsertVolumeStub != nil {
		return fake.InsertVolumeStub(arg1)
	} else {
		return fake.insertVolumeReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) InsertVolumeCallCount() int {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return len(fake.insertVolumeArgsForCall)
}

func (fake *FakeGardenWorkerDB) InsertVolumeArgsForCall(i int) db.Volume {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return fake.insertVolumeArgsForCall[i].arg1
}

func (fake *FakeGardenWorkerDB) InsertVolumeReturns(result1 error) {
	fake.InsertVolumeStub = nil
	fake.insertVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGardenWorkerDB) SetVolumeTTL(arg1 string, arg2 time.Duration) error {
	fake.setVolumeTTLMutex.Lock()
	fake.setVolumeTTLArgsForCall = append(fake.setVolumeTTLArgsForCall, struct {
		arg1 string
		arg2 time.Duration
	}{arg1, arg2})
	fake.setVolumeTTLMutex.Unlock()
	if fake.SetVolumeTTLStub != nil {
		return fake.SetVolumeTTLStub(arg1, arg2)
	} else {
		return fake.setVolumeTTLReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) SetVolumeTTLCallCount() int {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return len(fake.setVolumeTTLArgsForCall)
}

func (fake *FakeGardenWorkerDB) SetVolumeTTLArgsForCall(i int) (string, time.Duration) {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return fake.setVolumeTTLArgsForCall[i].arg1, fake.setVolumeTTLArgsForCall[i].arg2
}

func (fake *FakeGardenWorkerDB) SetVolumeTTLReturns(result1 error) {
	fake.SetVolumeTTLStub = nil
	fake.setVolumeTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGardenWorkerDB) GetVolumeTTL(arg1 string) (time.Duration, bool, error) {
	fake.getVolumeTTLMutex.Lock()
	fake.getVolumeTTLArgsForCall = append(fake.getVolumeTTLArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getVolumeTTLMutex.Unlock()
	if fake.GetVolumeTTLStub != nil {
		return fake.GetVolumeTTLStub(arg1)
	} else {
		return fake.getVolumeTTLReturns.result1, fake.getVolumeTTLReturns.result2, fake.getVolumeTTLReturns.result3
	}
}

func (fake *FakeGardenWorkerDB) GetVolumeTTLCallCount() int {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return len(fake.getVolumeTTLArgsForCall)
}

func (fake *FakeGardenWorkerDB) GetVolumeTTLArgsForCall(i int) string {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return fake.getVolumeTTLArgsForCall[i].arg1
}

func (fake *FakeGardenWorkerDB) GetVolumeTTLReturns(result1 time.Duration, result2 bool, result3 error) {
	fake.GetVolumeTTLStub = nil
	fake.getVolumeTTLReturns = struct {
		result1 time.Duration
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifier(arg1 db.VolumeIdentifier) ([]db.SavedVolume, error) {
	fake.getVolumesByIdentifierMutex.Lock()
	fake.getVolumesByIdentifierArgsForCall = append(fake.getVolumesByIdentifierArgsForCall, struct {
		arg1 db.VolumeIdentifier
	}{arg1})
	fake.getVolumesByIdentifierMutex.Unlock()
	if fake.GetVolumesByIdentifierStub != nil {
		return fake.GetVolumesByIdentifierStub(arg1)
	} else {
		return fake.getVolumesByIdentifierReturns.result1, fake.getVolumesByIdentifierReturns.result2
	}
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifierCallCount() int {
	fake.getVolumesByIdentifierMutex.RLock()
	defer fake.getVolumesByIdentifierMutex.RUnlock()
	return len(fake.getVolumesByIdentifierArgsForCall)
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifierArgsForCall(i int) db.VolumeIdentifier {
	fake.getVolumesByIdentifierMutex.RLock()
	defer fake.getVolumesByIdentifierMutex.RUnlock()
	return fake.getVolumesByIdentifierArgsForCall[i].arg1
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifierReturns(result1 []db.SavedVolume, result2 error) {
	fake.GetVolumesByIdentifierStub = nil
	fake.getVolumesByIdentifierReturns = struct {
		result1 []db.SavedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeGardenWorkerDB) ReapVolume(arg1 string) error {
	fake.reapVolumeMutex.Lock()
	fake.reapVolumeArgsForCall = append(fake.reapVolumeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.reapVolumeMutex.Unlock()
	if fake.ReapVolumeStub != nil {
		return fake.ReapVolumeStub(arg1)
	} else {
		return fake.reapVolumeReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) ReapVolumeCallCount() int {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return len(fake.reapVolumeArgsForCall)
}

func (fake *FakeGardenWorkerDB) ReapVolumeArgsForCall(i int) string {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return fake.reapVolumeArgsForCall[i].arg1
}

func (fake *FakeGardenWorkerDB) ReapVolumeReturns(result1 error) {
	fake.ReapVolumeStub = nil
	fake.reapVolumeReturns = struct {
		result1 error
	}{result1}
}

var _ worker.GardenWorkerDB = new(FakeGardenWorkerDB)