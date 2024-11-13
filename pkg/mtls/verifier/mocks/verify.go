// Code generated by MockGen. DO NOT EDIT.
// Source: verify.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/verify.go -source verify.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	tls "crypto/tls"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTLSConfigurer is a mock of TLSConfigurer interface.
type MockTLSConfigurer struct {
	ctrl     *gomock.Controller
	recorder *MockTLSConfigurerMockRecorder
	isgomock struct{}
}

// MockTLSConfigurerMockRecorder is the mock recorder for MockTLSConfigurer.
type MockTLSConfigurerMockRecorder struct {
	mock *MockTLSConfigurer
}

// NewMockTLSConfigurer creates a new mock instance.
func NewMockTLSConfigurer(ctrl *gomock.Controller) *MockTLSConfigurer {
	mock := &MockTLSConfigurer{ctrl: ctrl}
	mock.recorder = &MockTLSConfigurerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTLSConfigurer) EXPECT() *MockTLSConfigurerMockRecorder {
	return m.recorder
}

// TLSConfig mocks base method.
func (m *MockTLSConfigurer) TLSConfig() (*tls.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLSConfig")
	ret0, _ := ret[0].(*tls.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TLSConfig indicates an expected call of TLSConfig.
func (mr *MockTLSConfigurerMockRecorder) TLSConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSConfig", reflect.TypeOf((*MockTLSConfigurer)(nil).TLSConfig))
}
