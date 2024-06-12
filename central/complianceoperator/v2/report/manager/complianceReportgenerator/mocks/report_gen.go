// Code generated by MockGen. DO NOT EDIT.
// Source: report_gen.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/report_gen.go -source report_gen.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	complianceReportgenerator "github.com/stackrox/rox/central/complianceoperator/v2/report/manager/complianceReportgenerator"
	gomock "go.uber.org/mock/gomock"
)

// MockComplianceReportGenerator is a mock of ComplianceReportGenerator interface.
type MockComplianceReportGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockComplianceReportGeneratorMockRecorder
}

// MockComplianceReportGeneratorMockRecorder is the mock recorder for MockComplianceReportGenerator.
type MockComplianceReportGeneratorMockRecorder struct {
	mock *MockComplianceReportGenerator
}

// NewMockComplianceReportGenerator creates a new mock instance.
func NewMockComplianceReportGenerator(ctrl *gomock.Controller) *MockComplianceReportGenerator {
	mock := &MockComplianceReportGenerator{ctrl: ctrl}
	mock.recorder = &MockComplianceReportGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplianceReportGenerator) EXPECT() *MockComplianceReportGeneratorMockRecorder {
	return m.recorder
}

// ProcessReportRequest mocks base method.
func (m *MockComplianceReportGenerator) ProcessReportRequest(req *complianceReportgenerator.ComplianceReportRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessReportRequest", req)
}

// ProcessReportRequest indicates an expected call of ProcessReportRequest.
func (mr *MockComplianceReportGeneratorMockRecorder) ProcessReportRequest(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessReportRequest", reflect.TypeOf((*MockComplianceReportGenerator)(nil).ProcessReportRequest), req)
}