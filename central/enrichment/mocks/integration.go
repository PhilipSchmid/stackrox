// Code generated by MockGen. DO NOT EDIT.
// Source: integration.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/integration.go -source integration.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
	isgomock struct{}
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Remove mocks base method.
func (m *MockManager) Remove(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockManagerMockRecorder) Remove(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockManager)(nil).Remove), id)
}

// Upsert mocks base method.
func (m *MockManager) Upsert(integration *storage.ImageIntegration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", integration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockManagerMockRecorder) Upsert(integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockManager)(nil).Upsert), integration)
}
