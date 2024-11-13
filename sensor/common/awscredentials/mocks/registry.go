// Code generated by MockGen. DO NOT EDIT.
// Source: registry.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/registry.go -source registry.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	awscredentials "github.com/stackrox/rox/sensor/common/awscredentials"
	gomock "go.uber.org/mock/gomock"
)

// MockRegistryCredentialsManager is a mock of RegistryCredentialsManager interface.
type MockRegistryCredentialsManager struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryCredentialsManagerMockRecorder
	isgomock struct{}
}

// MockRegistryCredentialsManagerMockRecorder is the mock recorder for MockRegistryCredentialsManager.
type MockRegistryCredentialsManagerMockRecorder struct {
	mock *MockRegistryCredentialsManager
}

// NewMockRegistryCredentialsManager creates a new mock instance.
func NewMockRegistryCredentialsManager(ctrl *gomock.Controller) *MockRegistryCredentialsManager {
	mock := &MockRegistryCredentialsManager{ctrl: ctrl}
	mock.recorder = &MockRegistryCredentialsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryCredentialsManager) EXPECT() *MockRegistryCredentialsManagerMockRecorder {
	return m.recorder
}

// GetRegistryCredentials mocks base method.
func (m *MockRegistryCredentialsManager) GetRegistryCredentials(r string) *awscredentials.RegistryCredentials {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistryCredentials", r)
	ret0, _ := ret[0].(*awscredentials.RegistryCredentials)
	return ret0
}

// GetRegistryCredentials indicates an expected call of GetRegistryCredentials.
func (mr *MockRegistryCredentialsManagerMockRecorder) GetRegistryCredentials(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistryCredentials", reflect.TypeOf((*MockRegistryCredentialsManager)(nil).GetRegistryCredentials), r)
}

// Start mocks base method.
func (m *MockRegistryCredentialsManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockRegistryCredentialsManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRegistryCredentialsManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockRegistryCredentialsManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockRegistryCredentialsManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRegistryCredentialsManager)(nil).Stop))
}
