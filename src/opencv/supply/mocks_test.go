// Code generated by MockGen. DO NOT EDIT.
// Source: supply.go

// Package supply_test is a generated GoMock package.
package supply_test

import (
	reflect "reflect"

	libbuildpack "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
)

// MockManifest is a mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *MockManifestMockRecorder
}

// MockManifestMockRecorder is the mock recorder for MockManifest
type MockManifestMockRecorder struct {
	mock *MockManifest
}

// NewMockManifest creates a new mock instance
func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &MockManifestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManifest) EXPECT() *MockManifestMockRecorder {
	return m.recorder
}

// DefaultVersion mocks base method
func (m *MockManifest) DefaultVersion(arg0 string) (libbuildpack.Dependency, error) {
	ret := m.ctrl.Call(m, "DefaultVersion", arg0)
	ret0, _ := ret[0].(libbuildpack.Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultVersion indicates an expected call of DefaultVersion
func (mr *MockManifestMockRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultVersion", reflect.TypeOf((*MockManifest)(nil).DefaultVersion), arg0)
}

// InstallDependency mocks base method
func (m *MockManifest) InstallDependency(arg0 libbuildpack.Dependency, arg1 string) error {
	ret := m.ctrl.Call(m, "InstallDependency", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependency indicates an expected call of InstallDependency
func (mr *MockManifestMockRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallDependency", reflect.TypeOf((*MockManifest)(nil).InstallDependency), arg0, arg1)
}

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStager) EXPECT() *MockStagerMockRecorder {
	return m.recorder
}

// AddBinDependencyLink mocks base method
func (m *MockStager) AddBinDependencyLink(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AddBinDependencyLink", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinDependencyLink indicates an expected call of AddBinDependencyLink
func (mr *MockStagerMockRecorder) AddBinDependencyLink(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinDependencyLink", reflect.TypeOf((*MockStager)(nil).AddBinDependencyLink), arg0, arg1)
}

// DepDir mocks base method
func (m *MockStager) DepDir() string {
	ret := m.ctrl.Call(m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepDir indicates an expected call of DepDir
func (mr *MockStagerMockRecorder) DepDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepDir", reflect.TypeOf((*MockStager)(nil).DepDir))
}
