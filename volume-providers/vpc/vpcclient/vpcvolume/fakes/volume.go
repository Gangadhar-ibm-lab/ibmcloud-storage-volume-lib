// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/vpcclient/models"
	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/vpcclient/vpcvolume"
	"go.uber.org/zap"
)

type VolumeService struct {
	CheckVolumeTagStub        func(string, string, *zap.Logger) error
	checkVolumeTagMutex       sync.RWMutex
	checkVolumeTagArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *zap.Logger
	}
	checkVolumeTagReturns struct {
		result1 error
	}
	checkVolumeTagReturnsOnCall map[int]struct {
		result1 error
	}
	CreateVolumeStub        func(*models.Volume, *zap.Logger) (*models.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 *models.Volume
		arg2 *zap.Logger
	}
	createVolumeReturns struct {
		result1 *models.Volume
		result2 error
	}
	createVolumeReturnsOnCall map[int]struct {
		result1 *models.Volume
		result2 error
	}
	DeleteVolumeStub        func(string, *zap.Logger) error
	deleteVolumeMutex       sync.RWMutex
	deleteVolumeArgsForCall []struct {
		arg1 string
		arg2 *zap.Logger
	}
	deleteVolumeReturns struct {
		result1 error
	}
	deleteVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteVolumeTagStub        func(string, string, *zap.Logger) error
	deleteVolumeTagMutex       sync.RWMutex
	deleteVolumeTagArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *zap.Logger
	}
	deleteVolumeTagReturns struct {
		result1 error
	}
	deleteVolumeTagReturnsOnCall map[int]struct {
		result1 error
	}
	GetVolumeStub        func(string, *zap.Logger) (*models.Volume, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		arg1 string
		arg2 *zap.Logger
	}
	getVolumeReturns struct {
		result1 *models.Volume
		result2 error
	}
	getVolumeReturnsOnCall map[int]struct {
		result1 *models.Volume
		result2 error
	}
	GetVolumeByNameStub        func(string, *zap.Logger) (*models.Volume, error)
	getVolumeByNameMutex       sync.RWMutex
	getVolumeByNameArgsForCall []struct {
		arg1 string
		arg2 *zap.Logger
	}
	getVolumeByNameReturns struct {
		result1 *models.Volume
		result2 error
	}
	getVolumeByNameReturnsOnCall map[int]struct {
		result1 *models.Volume
		result2 error
	}
	ListVolumeTagsStub        func(string, *zap.Logger) (*[]string, error)
	listVolumeTagsMutex       sync.RWMutex
	listVolumeTagsArgsForCall []struct {
		arg1 string
		arg2 *zap.Logger
	}
	listVolumeTagsReturns struct {
		result1 *[]string
		result2 error
	}
	listVolumeTagsReturnsOnCall map[int]struct {
		result1 *[]string
		result2 error
	}
	ListVolumesStub        func(int, *models.ListVolumeFilters, *zap.Logger) (*models.VolumeList, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		arg1 int
		arg2 *models.ListVolumeFilters
		arg3 *zap.Logger
	}
	listVolumesReturns struct {
		result1 *models.VolumeList
		result2 error
	}
	listVolumesReturnsOnCall map[int]struct {
		result1 *models.VolumeList
		result2 error
	}
	SetVolumeTagStub        func(string, string, *zap.Logger) error
	setVolumeTagMutex       sync.RWMutex
	setVolumeTagArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *zap.Logger
	}
	setVolumeTagReturns struct {
		result1 error
	}
	setVolumeTagReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *VolumeService) CheckVolumeTag(arg1 string, arg2 string, arg3 *zap.Logger) error {
	fake.checkVolumeTagMutex.Lock()
	ret, specificReturn := fake.checkVolumeTagReturnsOnCall[len(fake.checkVolumeTagArgsForCall)]
	fake.checkVolumeTagArgsForCall = append(fake.checkVolumeTagArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *zap.Logger
	}{arg1, arg2, arg3})
	fake.recordInvocation("CheckVolumeTag", []interface{}{arg1, arg2, arg3})
	fake.checkVolumeTagMutex.Unlock()
	if fake.CheckVolumeTagStub != nil {
		return fake.CheckVolumeTagStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkVolumeTagReturns
	return fakeReturns.result1
}

func (fake *VolumeService) CheckVolumeTagCallCount() int {
	fake.checkVolumeTagMutex.RLock()
	defer fake.checkVolumeTagMutex.RUnlock()
	return len(fake.checkVolumeTagArgsForCall)
}

func (fake *VolumeService) CheckVolumeTagCalls(stub func(string, string, *zap.Logger) error) {
	fake.checkVolumeTagMutex.Lock()
	defer fake.checkVolumeTagMutex.Unlock()
	fake.CheckVolumeTagStub = stub
}

func (fake *VolumeService) CheckVolumeTagArgsForCall(i int) (string, string, *zap.Logger) {
	fake.checkVolumeTagMutex.RLock()
	defer fake.checkVolumeTagMutex.RUnlock()
	argsForCall := fake.checkVolumeTagArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *VolumeService) CheckVolumeTagReturns(result1 error) {
	fake.checkVolumeTagMutex.Lock()
	defer fake.checkVolumeTagMutex.Unlock()
	fake.CheckVolumeTagStub = nil
	fake.checkVolumeTagReturns = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) CheckVolumeTagReturnsOnCall(i int, result1 error) {
	fake.checkVolumeTagMutex.Lock()
	defer fake.checkVolumeTagMutex.Unlock()
	fake.CheckVolumeTagStub = nil
	if fake.checkVolumeTagReturnsOnCall == nil {
		fake.checkVolumeTagReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkVolumeTagReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) CreateVolume(arg1 *models.Volume, arg2 *zap.Logger) (*models.Volume, error) {
	fake.createVolumeMutex.Lock()
	ret, specificReturn := fake.createVolumeReturnsOnCall[len(fake.createVolumeArgsForCall)]
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 *models.Volume
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("CreateVolume", []interface{}{arg1, arg2})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createVolumeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeService) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *VolumeService) CreateVolumeCalls(stub func(*models.Volume, *zap.Logger) (*models.Volume, error)) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = stub
}

func (fake *VolumeService) CreateVolumeArgsForCall(i int) (*models.Volume, *zap.Logger) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	argsForCall := fake.createVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeService) CreateVolumeReturns(result1 *models.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 *models.Volume
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) CreateVolumeReturnsOnCall(i int, result1 *models.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	if fake.createVolumeReturnsOnCall == nil {
		fake.createVolumeReturnsOnCall = make(map[int]struct {
			result1 *models.Volume
			result2 error
		})
	}
	fake.createVolumeReturnsOnCall[i] = struct {
		result1 *models.Volume
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) DeleteVolume(arg1 string, arg2 *zap.Logger) error {
	fake.deleteVolumeMutex.Lock()
	ret, specificReturn := fake.deleteVolumeReturnsOnCall[len(fake.deleteVolumeArgsForCall)]
	fake.deleteVolumeArgsForCall = append(fake.deleteVolumeArgsForCall, struct {
		arg1 string
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("DeleteVolume", []interface{}{arg1, arg2})
	fake.deleteVolumeMutex.Unlock()
	if fake.DeleteVolumeStub != nil {
		return fake.DeleteVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteVolumeReturns
	return fakeReturns.result1
}

func (fake *VolumeService) DeleteVolumeCallCount() int {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return len(fake.deleteVolumeArgsForCall)
}

func (fake *VolumeService) DeleteVolumeCalls(stub func(string, *zap.Logger) error) {
	fake.deleteVolumeMutex.Lock()
	defer fake.deleteVolumeMutex.Unlock()
	fake.DeleteVolumeStub = stub
}

func (fake *VolumeService) DeleteVolumeArgsForCall(i int) (string, *zap.Logger) {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	argsForCall := fake.deleteVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeService) DeleteVolumeReturns(result1 error) {
	fake.deleteVolumeMutex.Lock()
	defer fake.deleteVolumeMutex.Unlock()
	fake.DeleteVolumeStub = nil
	fake.deleteVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) DeleteVolumeReturnsOnCall(i int, result1 error) {
	fake.deleteVolumeMutex.Lock()
	defer fake.deleteVolumeMutex.Unlock()
	fake.DeleteVolumeStub = nil
	if fake.deleteVolumeReturnsOnCall == nil {
		fake.deleteVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) DeleteVolumeTag(arg1 string, arg2 string, arg3 *zap.Logger) error {
	fake.deleteVolumeTagMutex.Lock()
	ret, specificReturn := fake.deleteVolumeTagReturnsOnCall[len(fake.deleteVolumeTagArgsForCall)]
	fake.deleteVolumeTagArgsForCall = append(fake.deleteVolumeTagArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *zap.Logger
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteVolumeTag", []interface{}{arg1, arg2, arg3})
	fake.deleteVolumeTagMutex.Unlock()
	if fake.DeleteVolumeTagStub != nil {
		return fake.DeleteVolumeTagStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteVolumeTagReturns
	return fakeReturns.result1
}

func (fake *VolumeService) DeleteVolumeTagCallCount() int {
	fake.deleteVolumeTagMutex.RLock()
	defer fake.deleteVolumeTagMutex.RUnlock()
	return len(fake.deleteVolumeTagArgsForCall)
}

func (fake *VolumeService) DeleteVolumeTagCalls(stub func(string, string, *zap.Logger) error) {
	fake.deleteVolumeTagMutex.Lock()
	defer fake.deleteVolumeTagMutex.Unlock()
	fake.DeleteVolumeTagStub = stub
}

func (fake *VolumeService) DeleteVolumeTagArgsForCall(i int) (string, string, *zap.Logger) {
	fake.deleteVolumeTagMutex.RLock()
	defer fake.deleteVolumeTagMutex.RUnlock()
	argsForCall := fake.deleteVolumeTagArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *VolumeService) DeleteVolumeTagReturns(result1 error) {
	fake.deleteVolumeTagMutex.Lock()
	defer fake.deleteVolumeTagMutex.Unlock()
	fake.DeleteVolumeTagStub = nil
	fake.deleteVolumeTagReturns = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) DeleteVolumeTagReturnsOnCall(i int, result1 error) {
	fake.deleteVolumeTagMutex.Lock()
	defer fake.deleteVolumeTagMutex.Unlock()
	fake.DeleteVolumeTagStub = nil
	if fake.deleteVolumeTagReturnsOnCall == nil {
		fake.deleteVolumeTagReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteVolumeTagReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) GetVolume(arg1 string, arg2 *zap.Logger) (*models.Volume, error) {
	fake.getVolumeMutex.Lock()
	ret, specificReturn := fake.getVolumeReturnsOnCall[len(fake.getVolumeArgsForCall)]
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		arg1 string
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetVolume", []interface{}{arg1, arg2})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getVolumeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeService) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *VolumeService) GetVolumeCalls(stub func(string, *zap.Logger) (*models.Volume, error)) {
	fake.getVolumeMutex.Lock()
	defer fake.getVolumeMutex.Unlock()
	fake.GetVolumeStub = stub
}

func (fake *VolumeService) GetVolumeArgsForCall(i int) (string, *zap.Logger) {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	argsForCall := fake.getVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeService) GetVolumeReturns(result1 *models.Volume, result2 error) {
	fake.getVolumeMutex.Lock()
	defer fake.getVolumeMutex.Unlock()
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 *models.Volume
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) GetVolumeReturnsOnCall(i int, result1 *models.Volume, result2 error) {
	fake.getVolumeMutex.Lock()
	defer fake.getVolumeMutex.Unlock()
	fake.GetVolumeStub = nil
	if fake.getVolumeReturnsOnCall == nil {
		fake.getVolumeReturnsOnCall = make(map[int]struct {
			result1 *models.Volume
			result2 error
		})
	}
	fake.getVolumeReturnsOnCall[i] = struct {
		result1 *models.Volume
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) GetVolumeByName(arg1 string, arg2 *zap.Logger) (*models.Volume, error) {
	fake.getVolumeByNameMutex.Lock()
	ret, specificReturn := fake.getVolumeByNameReturnsOnCall[len(fake.getVolumeByNameArgsForCall)]
	fake.getVolumeByNameArgsForCall = append(fake.getVolumeByNameArgsForCall, struct {
		arg1 string
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetVolumeByName", []interface{}{arg1, arg2})
	fake.getVolumeByNameMutex.Unlock()
	if fake.GetVolumeByNameStub != nil {
		return fake.GetVolumeByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getVolumeByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeService) GetVolumeByNameCallCount() int {
	fake.getVolumeByNameMutex.RLock()
	defer fake.getVolumeByNameMutex.RUnlock()
	return len(fake.getVolumeByNameArgsForCall)
}

func (fake *VolumeService) GetVolumeByNameCalls(stub func(string, *zap.Logger) (*models.Volume, error)) {
	fake.getVolumeByNameMutex.Lock()
	defer fake.getVolumeByNameMutex.Unlock()
	fake.GetVolumeByNameStub = stub
}

func (fake *VolumeService) GetVolumeByNameArgsForCall(i int) (string, *zap.Logger) {
	fake.getVolumeByNameMutex.RLock()
	defer fake.getVolumeByNameMutex.RUnlock()
	argsForCall := fake.getVolumeByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeService) GetVolumeByNameReturns(result1 *models.Volume, result2 error) {
	fake.getVolumeByNameMutex.Lock()
	defer fake.getVolumeByNameMutex.Unlock()
	fake.GetVolumeByNameStub = nil
	fake.getVolumeByNameReturns = struct {
		result1 *models.Volume
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) GetVolumeByNameReturnsOnCall(i int, result1 *models.Volume, result2 error) {
	fake.getVolumeByNameMutex.Lock()
	defer fake.getVolumeByNameMutex.Unlock()
	fake.GetVolumeByNameStub = nil
	if fake.getVolumeByNameReturnsOnCall == nil {
		fake.getVolumeByNameReturnsOnCall = make(map[int]struct {
			result1 *models.Volume
			result2 error
		})
	}
	fake.getVolumeByNameReturnsOnCall[i] = struct {
		result1 *models.Volume
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) ListVolumeTags(arg1 string, arg2 *zap.Logger) (*[]string, error) {
	fake.listVolumeTagsMutex.Lock()
	ret, specificReturn := fake.listVolumeTagsReturnsOnCall[len(fake.listVolumeTagsArgsForCall)]
	fake.listVolumeTagsArgsForCall = append(fake.listVolumeTagsArgsForCall, struct {
		arg1 string
		arg2 *zap.Logger
	}{arg1, arg2})
	fake.recordInvocation("ListVolumeTags", []interface{}{arg1, arg2})
	fake.listVolumeTagsMutex.Unlock()
	if fake.ListVolumeTagsStub != nil {
		return fake.ListVolumeTagsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listVolumeTagsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeService) ListVolumeTagsCallCount() int {
	fake.listVolumeTagsMutex.RLock()
	defer fake.listVolumeTagsMutex.RUnlock()
	return len(fake.listVolumeTagsArgsForCall)
}

func (fake *VolumeService) ListVolumeTagsCalls(stub func(string, *zap.Logger) (*[]string, error)) {
	fake.listVolumeTagsMutex.Lock()
	defer fake.listVolumeTagsMutex.Unlock()
	fake.ListVolumeTagsStub = stub
}

func (fake *VolumeService) ListVolumeTagsArgsForCall(i int) (string, *zap.Logger) {
	fake.listVolumeTagsMutex.RLock()
	defer fake.listVolumeTagsMutex.RUnlock()
	argsForCall := fake.listVolumeTagsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VolumeService) ListVolumeTagsReturns(result1 *[]string, result2 error) {
	fake.listVolumeTagsMutex.Lock()
	defer fake.listVolumeTagsMutex.Unlock()
	fake.ListVolumeTagsStub = nil
	fake.listVolumeTagsReturns = struct {
		result1 *[]string
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) ListVolumeTagsReturnsOnCall(i int, result1 *[]string, result2 error) {
	fake.listVolumeTagsMutex.Lock()
	defer fake.listVolumeTagsMutex.Unlock()
	fake.ListVolumeTagsStub = nil
	if fake.listVolumeTagsReturnsOnCall == nil {
		fake.listVolumeTagsReturnsOnCall = make(map[int]struct {
			result1 *[]string
			result2 error
		})
	}
	fake.listVolumeTagsReturnsOnCall[i] = struct {
		result1 *[]string
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) ListVolumes(arg1 int, arg2 *models.ListVolumeFilters, arg3 *zap.Logger) (*models.VolumeList, error) {
	fake.listVolumesMutex.Lock()
	ret, specificReturn := fake.listVolumesReturnsOnCall[len(fake.listVolumesArgsForCall)]
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		arg1 int
		arg2 *models.ListVolumeFilters
		arg3 *zap.Logger
	}{arg1, arg2, arg3})
	fake.recordInvocation("ListVolumes", []interface{}{arg1, arg2, arg3})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listVolumesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *VolumeService) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *VolumeService) ListVolumesCalls(stub func(int, *models.ListVolumeFilters, *zap.Logger) (*models.VolumeList, error)) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = stub
}

func (fake *VolumeService) ListVolumesArgsForCall(i int) (int, *models.ListVolumeFilters, *zap.Logger) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	argsForCall := fake.listVolumesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *VolumeService) ListVolumesReturns(result1 *models.VolumeList, result2 error) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 *models.VolumeList
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) ListVolumesReturnsOnCall(i int, result1 *models.VolumeList, result2 error) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = nil
	if fake.listVolumesReturnsOnCall == nil {
		fake.listVolumesReturnsOnCall = make(map[int]struct {
			result1 *models.VolumeList
			result2 error
		})
	}
	fake.listVolumesReturnsOnCall[i] = struct {
		result1 *models.VolumeList
		result2 error
	}{result1, result2}
}

func (fake *VolumeService) SetVolumeTag(arg1 string, arg2 string, arg3 *zap.Logger) error {
	fake.setVolumeTagMutex.Lock()
	ret, specificReturn := fake.setVolumeTagReturnsOnCall[len(fake.setVolumeTagArgsForCall)]
	fake.setVolumeTagArgsForCall = append(fake.setVolumeTagArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *zap.Logger
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetVolumeTag", []interface{}{arg1, arg2, arg3})
	fake.setVolumeTagMutex.Unlock()
	if fake.SetVolumeTagStub != nil {
		return fake.SetVolumeTagStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setVolumeTagReturns
	return fakeReturns.result1
}

func (fake *VolumeService) SetVolumeTagCallCount() int {
	fake.setVolumeTagMutex.RLock()
	defer fake.setVolumeTagMutex.RUnlock()
	return len(fake.setVolumeTagArgsForCall)
}

func (fake *VolumeService) SetVolumeTagCalls(stub func(string, string, *zap.Logger) error) {
	fake.setVolumeTagMutex.Lock()
	defer fake.setVolumeTagMutex.Unlock()
	fake.SetVolumeTagStub = stub
}

func (fake *VolumeService) SetVolumeTagArgsForCall(i int) (string, string, *zap.Logger) {
	fake.setVolumeTagMutex.RLock()
	defer fake.setVolumeTagMutex.RUnlock()
	argsForCall := fake.setVolumeTagArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *VolumeService) SetVolumeTagReturns(result1 error) {
	fake.setVolumeTagMutex.Lock()
	defer fake.setVolumeTagMutex.Unlock()
	fake.SetVolumeTagStub = nil
	fake.setVolumeTagReturns = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) SetVolumeTagReturnsOnCall(i int, result1 error) {
	fake.setVolumeTagMutex.Lock()
	defer fake.setVolumeTagMutex.Unlock()
	fake.SetVolumeTagStub = nil
	if fake.setVolumeTagReturnsOnCall == nil {
		fake.setVolumeTagReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setVolumeTagReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VolumeService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkVolumeTagMutex.RLock()
	defer fake.checkVolumeTagMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	fake.deleteVolumeTagMutex.RLock()
	defer fake.deleteVolumeTagMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.getVolumeByNameMutex.RLock()
	defer fake.getVolumeByNameMutex.RUnlock()
	fake.listVolumeTagsMutex.RLock()
	defer fake.listVolumeTagsMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.setVolumeTagMutex.RLock()
	defer fake.setVolumeTagMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *VolumeService) recordInvocation(key string, args []interface{}) {
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

var _ vpcvolume.VolumeManager = new(VolumeService)
