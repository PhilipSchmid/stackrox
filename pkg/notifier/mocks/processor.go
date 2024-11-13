// Code generated by MockGen. DO NOT EDIT.
// Source: processor.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/processor.go -source processor.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	notifiers "github.com/stackrox/rox/pkg/notifiers"
	gomock "go.uber.org/mock/gomock"
)

// MockProcessor is a mock of Processor interface.
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
	isgomock struct{}
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor.
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance.
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// GetNotifier mocks base method.
func (m *MockProcessor) GetNotifier(ctx context.Context, id string) notifiers.Notifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifier", ctx, id)
	ret0, _ := ret[0].(notifiers.Notifier)
	return ret0
}

// GetNotifier indicates an expected call of GetNotifier.
func (mr *MockProcessorMockRecorder) GetNotifier(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifier", reflect.TypeOf((*MockProcessor)(nil).GetNotifier), ctx, id)
}

// HasEnabledAuditNotifiers mocks base method.
func (m *MockProcessor) HasEnabledAuditNotifiers() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasEnabledAuditNotifiers")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasEnabledAuditNotifiers indicates an expected call of HasEnabledAuditNotifiers.
func (mr *MockProcessorMockRecorder) HasEnabledAuditNotifiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasEnabledAuditNotifiers", reflect.TypeOf((*MockProcessor)(nil).HasEnabledAuditNotifiers))
}

// HasNotifiers mocks base method.
func (m *MockProcessor) HasNotifiers() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNotifiers")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNotifiers indicates an expected call of HasNotifiers.
func (mr *MockProcessorMockRecorder) HasNotifiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNotifiers", reflect.TypeOf((*MockProcessor)(nil).HasNotifiers))
}

// ProcessAlert mocks base method.
func (m *MockProcessor) ProcessAlert(ctx context.Context, alert *storage.Alert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessAlert", ctx, alert)
}

// ProcessAlert indicates an expected call of ProcessAlert.
func (mr *MockProcessorMockRecorder) ProcessAlert(ctx, alert any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessAlert", reflect.TypeOf((*MockProcessor)(nil).ProcessAlert), ctx, alert)
}

// ProcessAuditMessage mocks base method.
func (m *MockProcessor) ProcessAuditMessage(ctx context.Context, msg *v1.Audit_Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessAuditMessage", ctx, msg)
}

// ProcessAuditMessage indicates an expected call of ProcessAuditMessage.
func (mr *MockProcessorMockRecorder) ProcessAuditMessage(ctx, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessAuditMessage", reflect.TypeOf((*MockProcessor)(nil).ProcessAuditMessage), ctx, msg)
}

// RemoveNotifier mocks base method.
func (m *MockProcessor) RemoveNotifier(ctx context.Context, id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveNotifier", ctx, id)
}

// RemoveNotifier indicates an expected call of RemoveNotifier.
func (mr *MockProcessorMockRecorder) RemoveNotifier(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNotifier", reflect.TypeOf((*MockProcessor)(nil).RemoveNotifier), ctx, id)
}

// UpdateNotifier mocks base method.
func (m *MockProcessor) UpdateNotifier(ctx context.Context, notifier notifiers.Notifier) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateNotifier", ctx, notifier)
}

// UpdateNotifier indicates an expected call of UpdateNotifier.
func (mr *MockProcessorMockRecorder) UpdateNotifier(ctx, notifier any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotifier", reflect.TypeOf((*MockProcessor)(nil).UpdateNotifier), ctx, notifier)
}

// UpdateNotifierHealthStatus mocks base method.
func (m *MockProcessor) UpdateNotifierHealthStatus(notifier notifiers.Notifier, healthStatus storage.IntegrationHealth_Status, errMessage string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateNotifierHealthStatus", notifier, healthStatus, errMessage)
}

// UpdateNotifierHealthStatus indicates an expected call of UpdateNotifierHealthStatus.
func (mr *MockProcessorMockRecorder) UpdateNotifierHealthStatus(notifier, healthStatus, errMessage any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotifierHealthStatus", reflect.TypeOf((*MockProcessor)(nil).UpdateNotifierHealthStatus), notifier, healthStatus, errMessage)
}
