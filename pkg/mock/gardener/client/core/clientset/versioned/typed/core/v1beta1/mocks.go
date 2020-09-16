// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/core/clientset/versioned/typed/core/v1beta1 (interfaces: CoreV1beta1Interface,ShootInterface,SeedInterface)

// Package v1beta1 is a generated GoMock package.
package v1beta1

import (
	context "context"
	reflect "reflect"

	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta10 "github.com/gardener/gardener/pkg/client/core/clientset/versioned/typed/core/v1beta1"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MockCoreV1beta1Interface is a mock of CoreV1beta1Interface interface.
type MockCoreV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockCoreV1beta1InterfaceMockRecorder
}

// MockCoreV1beta1InterfaceMockRecorder is the mock recorder for MockCoreV1beta1Interface.
type MockCoreV1beta1InterfaceMockRecorder struct {
	mock *MockCoreV1beta1Interface
}

// NewMockCoreV1beta1Interface creates a new mock instance.
func NewMockCoreV1beta1Interface(ctrl *gomock.Controller) *MockCoreV1beta1Interface {
	mock := &MockCoreV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockCoreV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoreV1beta1Interface) EXPECT() *MockCoreV1beta1InterfaceMockRecorder {
	return m.recorder
}

// BackupBuckets mocks base method.
func (m *MockCoreV1beta1Interface) BackupBuckets() v1beta10.BackupBucketInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupBuckets")
	ret0, _ := ret[0].(v1beta10.BackupBucketInterface)
	return ret0
}

// BackupBuckets indicates an expected call of BackupBuckets.
func (mr *MockCoreV1beta1InterfaceMockRecorder) BackupBuckets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupBuckets", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).BackupBuckets))
}

// BackupEntries mocks base method.
func (m *MockCoreV1beta1Interface) BackupEntries(arg0 string) v1beta10.BackupEntryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupEntries", arg0)
	ret0, _ := ret[0].(v1beta10.BackupEntryInterface)
	return ret0
}

// BackupEntries indicates an expected call of BackupEntries.
func (mr *MockCoreV1beta1InterfaceMockRecorder) BackupEntries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupEntries", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).BackupEntries), arg0)
}

// CloudProfiles mocks base method.
func (m *MockCoreV1beta1Interface) CloudProfiles() v1beta10.CloudProfileInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProfiles")
	ret0, _ := ret[0].(v1beta10.CloudProfileInterface)
	return ret0
}

// CloudProfiles indicates an expected call of CloudProfiles.
func (mr *MockCoreV1beta1InterfaceMockRecorder) CloudProfiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProfiles", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).CloudProfiles))
}

// ControllerInstallations mocks base method.
func (m *MockCoreV1beta1Interface) ControllerInstallations() v1beta10.ControllerInstallationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerInstallations")
	ret0, _ := ret[0].(v1beta10.ControllerInstallationInterface)
	return ret0
}

// ControllerInstallations indicates an expected call of ControllerInstallations.
func (mr *MockCoreV1beta1InterfaceMockRecorder) ControllerInstallations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerInstallations", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).ControllerInstallations))
}

// ControllerRegistrations mocks base method.
func (m *MockCoreV1beta1Interface) ControllerRegistrations() v1beta10.ControllerRegistrationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerRegistrations")
	ret0, _ := ret[0].(v1beta10.ControllerRegistrationInterface)
	return ret0
}

// ControllerRegistrations indicates an expected call of ControllerRegistrations.
func (mr *MockCoreV1beta1InterfaceMockRecorder) ControllerRegistrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerRegistrations", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).ControllerRegistrations))
}

// Plants mocks base method.
func (m *MockCoreV1beta1Interface) Plants(arg0 string) v1beta10.PlantInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Plants", arg0)
	ret0, _ := ret[0].(v1beta10.PlantInterface)
	return ret0
}

// Plants indicates an expected call of Plants.
func (mr *MockCoreV1beta1InterfaceMockRecorder) Plants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Plants", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).Plants), arg0)
}

// Projects mocks base method.
func (m *MockCoreV1beta1Interface) Projects() v1beta10.ProjectInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects")
	ret0, _ := ret[0].(v1beta10.ProjectInterface)
	return ret0
}

// Projects indicates an expected call of Projects.
func (mr *MockCoreV1beta1InterfaceMockRecorder) Projects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).Projects))
}

// Quotas mocks base method.
func (m *MockCoreV1beta1Interface) Quotas(arg0 string) v1beta10.QuotaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quotas", arg0)
	ret0, _ := ret[0].(v1beta10.QuotaInterface)
	return ret0
}

// Quotas indicates an expected call of Quotas.
func (mr *MockCoreV1beta1InterfaceMockRecorder) Quotas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quotas", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).Quotas), arg0)
}

// RESTClient mocks base method.
func (m *MockCoreV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockCoreV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).RESTClient))
}

// SecretBindings mocks base method.
func (m *MockCoreV1beta1Interface) SecretBindings(arg0 string) v1beta10.SecretBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretBindings", arg0)
	ret0, _ := ret[0].(v1beta10.SecretBindingInterface)
	return ret0
}

// SecretBindings indicates an expected call of SecretBindings.
func (mr *MockCoreV1beta1InterfaceMockRecorder) SecretBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretBindings", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).SecretBindings), arg0)
}

// Seeds mocks base method.
func (m *MockCoreV1beta1Interface) Seeds() v1beta10.SeedInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seeds")
	ret0, _ := ret[0].(v1beta10.SeedInterface)
	return ret0
}

// Seeds indicates an expected call of Seeds.
func (mr *MockCoreV1beta1InterfaceMockRecorder) Seeds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seeds", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).Seeds))
}

// Shoots mocks base method.
func (m *MockCoreV1beta1Interface) Shoots(arg0 string) v1beta10.ShootInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shoots", arg0)
	ret0, _ := ret[0].(v1beta10.ShootInterface)
	return ret0
}

// Shoots indicates an expected call of Shoots.
func (mr *MockCoreV1beta1InterfaceMockRecorder) Shoots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shoots", reflect.TypeOf((*MockCoreV1beta1Interface)(nil).Shoots), arg0)
}

// MockShootInterface is a mock of ShootInterface interface.
type MockShootInterface struct {
	ctrl     *gomock.Controller
	recorder *MockShootInterfaceMockRecorder
}

// MockShootInterfaceMockRecorder is the mock recorder for MockShootInterface.
type MockShootInterfaceMockRecorder struct {
	mock *MockShootInterface
}

// NewMockShootInterface creates a new mock instance.
func NewMockShootInterface(ctrl *gomock.Controller) *MockShootInterface {
	mock := &MockShootInterface{ctrl: ctrl}
	mock.recorder = &MockShootInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShootInterface) EXPECT() *MockShootInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockShootInterface) Create(arg0 context.Context, arg1 *v1beta1.Shoot, arg2 v1.CreateOptions) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockShootInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShootInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockShootInterface) Delete(arg0 context.Context, arg1 string, arg2 v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockShootInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockShootInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockShootInterface) DeleteCollection(arg0 context.Context, arg1 v1.DeleteOptions, arg2 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockShootInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockShootInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockShootInterface) Get(arg0 context.Context, arg1 string, arg2 v1.GetOptions) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShootInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShootInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockShootInterface) List(arg0 context.Context, arg1 v1.ListOptions) (*v1beta1.ShootList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.ShootList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockShootInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockShootInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockShootInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v1.PatchOptions, arg5 ...string) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockShootInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockShootInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockShootInterface) Update(arg0 context.Context, arg1 *v1beta1.Shoot, arg2 v1.UpdateOptions) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockShootInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockShootInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockShootInterface) UpdateStatus(arg0 context.Context, arg1 *v1beta1.Shoot, arg2 v1.UpdateOptions) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockShootInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockShootInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockShootInterface) Watch(arg0 context.Context, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockShootInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockShootInterface)(nil).Watch), arg0, arg1)
}

// MockSeedInterface is a mock of SeedInterface interface.
type MockSeedInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSeedInterfaceMockRecorder
}

// MockSeedInterfaceMockRecorder is the mock recorder for MockSeedInterface.
type MockSeedInterfaceMockRecorder struct {
	mock *MockSeedInterface
}

// NewMockSeedInterface creates a new mock instance.
func NewMockSeedInterface(ctrl *gomock.Controller) *MockSeedInterface {
	mock := &MockSeedInterface{ctrl: ctrl}
	mock.recorder = &MockSeedInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeedInterface) EXPECT() *MockSeedInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSeedInterface) Create(arg0 context.Context, arg1 *v1beta1.Seed, arg2 v1.CreateOptions) (*v1beta1.Seed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Seed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSeedInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSeedInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockSeedInterface) Delete(arg0 context.Context, arg1 string, arg2 v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSeedInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSeedInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockSeedInterface) DeleteCollection(arg0 context.Context, arg1 v1.DeleteOptions, arg2 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockSeedInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockSeedInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockSeedInterface) Get(arg0 context.Context, arg1 string, arg2 v1.GetOptions) (*v1beta1.Seed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Seed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSeedInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSeedInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockSeedInterface) List(arg0 context.Context, arg1 v1.ListOptions) (*v1beta1.SeedList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.SeedList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSeedInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSeedInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockSeedInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v1.PatchOptions, arg5 ...string) (*v1beta1.Seed, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.Seed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockSeedInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockSeedInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockSeedInterface) Update(arg0 context.Context, arg1 *v1beta1.Seed, arg2 v1.UpdateOptions) (*v1beta1.Seed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Seed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSeedInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSeedInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockSeedInterface) UpdateStatus(arg0 context.Context, arg1 *v1beta1.Seed, arg2 v1.UpdateOptions) (*v1beta1.Seed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Seed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockSeedInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockSeedInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockSeedInterface) Watch(arg0 context.Context, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockSeedInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockSeedInterface)(nil).Watch), arg0, arg1)
}
