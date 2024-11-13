// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	datastore "github.com/stackrox/rox/central/complianceoperator/v2/integration/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
	isgomock struct{}
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AddComplianceIntegration mocks base method.
func (m *MockDataStore) AddComplianceIntegration(ctx context.Context, integration *storage.ComplianceIntegration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComplianceIntegration", ctx, integration)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComplianceIntegration indicates an expected call of AddComplianceIntegration.
func (mr *MockDataStoreMockRecorder) AddComplianceIntegration(ctx, integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComplianceIntegration", reflect.TypeOf((*MockDataStore)(nil).AddComplianceIntegration), ctx, integration)
}

// CountIntegrations mocks base method.
func (m *MockDataStore) CountIntegrations(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountIntegrations", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountIntegrations indicates an expected call of CountIntegrations.
func (mr *MockDataStoreMockRecorder) CountIntegrations(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountIntegrations", reflect.TypeOf((*MockDataStore)(nil).CountIntegrations), ctx, q)
}

// GetComplianceIntegration mocks base method.
func (m *MockDataStore) GetComplianceIntegration(ctx context.Context, id string) (*storage.ComplianceIntegration, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceIntegration", ctx, id)
	ret0, _ := ret[0].(*storage.ComplianceIntegration)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetComplianceIntegration indicates an expected call of GetComplianceIntegration.
func (mr *MockDataStoreMockRecorder) GetComplianceIntegration(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceIntegration", reflect.TypeOf((*MockDataStore)(nil).GetComplianceIntegration), ctx, id)
}

// GetComplianceIntegrationByCluster mocks base method.
func (m *MockDataStore) GetComplianceIntegrationByCluster(ctx context.Context, clusterID string) ([]*storage.ComplianceIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceIntegrationByCluster", ctx, clusterID)
	ret0, _ := ret[0].([]*storage.ComplianceIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComplianceIntegrationByCluster indicates an expected call of GetComplianceIntegrationByCluster.
func (mr *MockDataStoreMockRecorder) GetComplianceIntegrationByCluster(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceIntegrationByCluster", reflect.TypeOf((*MockDataStore)(nil).GetComplianceIntegrationByCluster), ctx, clusterID)
}

// GetComplianceIntegrations mocks base method.
func (m *MockDataStore) GetComplianceIntegrations(ctx context.Context, query *v1.Query) ([]*storage.ComplianceIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceIntegrations", ctx, query)
	ret0, _ := ret[0].([]*storage.ComplianceIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComplianceIntegrations indicates an expected call of GetComplianceIntegrations.
func (mr *MockDataStoreMockRecorder) GetComplianceIntegrations(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceIntegrations", reflect.TypeOf((*MockDataStore)(nil).GetComplianceIntegrations), ctx, query)
}

// GetComplianceIntegrationsView mocks base method.
func (m *MockDataStore) GetComplianceIntegrationsView(ctx context.Context, query *v1.Query) ([]*datastore.IntegrationDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceIntegrationsView", ctx, query)
	ret0, _ := ret[0].([]*datastore.IntegrationDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComplianceIntegrationsView indicates an expected call of GetComplianceIntegrationsView.
func (mr *MockDataStoreMockRecorder) GetComplianceIntegrationsView(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceIntegrationsView", reflect.TypeOf((*MockDataStore)(nil).GetComplianceIntegrationsView), ctx, query)
}

// RemoveComplianceIntegration mocks base method.
func (m *MockDataStore) RemoveComplianceIntegration(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveComplianceIntegration", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveComplianceIntegration indicates an expected call of RemoveComplianceIntegration.
func (mr *MockDataStoreMockRecorder) RemoveComplianceIntegration(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveComplianceIntegration", reflect.TypeOf((*MockDataStore)(nil).RemoveComplianceIntegration), ctx, id)
}

// RemoveComplianceIntegrationByCluster mocks base method.
func (m *MockDataStore) RemoveComplianceIntegrationByCluster(ctx context.Context, clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveComplianceIntegrationByCluster", ctx, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveComplianceIntegrationByCluster indicates an expected call of RemoveComplianceIntegrationByCluster.
func (mr *MockDataStoreMockRecorder) RemoveComplianceIntegrationByCluster(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveComplianceIntegrationByCluster", reflect.TypeOf((*MockDataStore)(nil).RemoveComplianceIntegrationByCluster), ctx, clusterID)
}

// UpdateComplianceIntegration mocks base method.
func (m *MockDataStore) UpdateComplianceIntegration(ctx context.Context, integration *storage.ComplianceIntegration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComplianceIntegration", ctx, integration)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateComplianceIntegration indicates an expected call of UpdateComplianceIntegration.
func (mr *MockDataStoreMockRecorder) UpdateComplianceIntegration(ctx, integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComplianceIntegration", reflect.TypeOf((*MockDataStore)(nil).UpdateComplianceIntegration), ctx, integration)
}
