// Code generated by MockGen. DO NOT EDIT.
// Source: reprocessor.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/reprocessor.go -source reprocessor.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPlatformReprocessor is a mock of PlatformReprocessor interface.
type MockPlatformReprocessor struct {
	ctrl     *gomock.Controller
	recorder *MockPlatformReprocessorMockRecorder
}

// MockPlatformReprocessorMockRecorder is the mock recorder for MockPlatformReprocessor.
type MockPlatformReprocessorMockRecorder struct {
	mock *MockPlatformReprocessor
}

// NewMockPlatformReprocessor creates a new mock instance.
func NewMockPlatformReprocessor(ctrl *gomock.Controller) *MockPlatformReprocessor {
	mock := &MockPlatformReprocessor{ctrl: ctrl}
	mock.recorder = &MockPlatformReprocessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlatformReprocessor) EXPECT() *MockPlatformReprocessorMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockPlatformReprocessor) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockPlatformReprocessorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPlatformReprocessor)(nil).Start))
}

// Stop mocks base method.
func (m *MockPlatformReprocessor) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockPlatformReprocessorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockPlatformReprocessor)(nil).Stop))
}