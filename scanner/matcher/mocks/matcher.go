// Code generated by MockGen. DO NOT EDIT.
// Source: matcher.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/matcher.go -source matcher.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	claircore "github.com/quay/claircore"
	gomock "go.uber.org/mock/gomock"
)

// MockMatcher is a mock of Matcher interface.
type MockMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMockRecorder
	isgomock struct{}
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher.
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance.
func NewMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMatcher) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMatcherMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMatcher)(nil).Close), ctx)
}

// GetKnownDistributions mocks base method.
func (m *MockMatcher) GetKnownDistributions(ctx context.Context) []claircore.Distribution {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKnownDistributions", ctx)
	ret0, _ := ret[0].([]claircore.Distribution)
	return ret0
}

// GetKnownDistributions indicates an expected call of GetKnownDistributions.
func (mr *MockMatcherMockRecorder) GetKnownDistributions(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnownDistributions", reflect.TypeOf((*MockMatcher)(nil).GetKnownDistributions), ctx)
}

// GetLastVulnerabilityUpdate mocks base method.
func (m *MockMatcher) GetLastVulnerabilityUpdate(ctx context.Context) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastVulnerabilityUpdate", ctx)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastVulnerabilityUpdate indicates an expected call of GetLastVulnerabilityUpdate.
func (mr *MockMatcherMockRecorder) GetLastVulnerabilityUpdate(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastVulnerabilityUpdate", reflect.TypeOf((*MockMatcher)(nil).GetLastVulnerabilityUpdate), ctx)
}

// GetVulnerabilities mocks base method.
func (m *MockMatcher) GetVulnerabilities(ctx context.Context, ir *claircore.IndexReport) (*claircore.VulnerabilityReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVulnerabilities", ctx, ir)
	ret0, _ := ret[0].(*claircore.VulnerabilityReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVulnerabilities indicates an expected call of GetVulnerabilities.
func (mr *MockMatcherMockRecorder) GetVulnerabilities(ctx, ir any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVulnerabilities", reflect.TypeOf((*MockMatcher)(nil).GetVulnerabilities), ctx, ir)
}

// Initialized mocks base method.
func (m *MockMatcher) Initialized(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialized", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialized indicates an expected call of Initialized.
func (mr *MockMatcherMockRecorder) Initialized(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialized", reflect.TypeOf((*MockMatcher)(nil).Initialized), ctx)
}

// Ready mocks base method.
func (m *MockMatcher) Ready(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ready indicates an expected call of Ready.
func (mr *MockMatcherMockRecorder) Ready(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockMatcher)(nil).Ready), ctx)
}
