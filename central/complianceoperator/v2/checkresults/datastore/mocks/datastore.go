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

	datastore "github.com/stackrox/rox/central/complianceoperator/v2/checkresults/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
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

// ComplianceCheckResultStats mocks base method.
func (m *MockDataStore) ComplianceCheckResultStats(ctx context.Context, query *v1.Query) ([]*datastore.ResourceResultCountByClusterScan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceCheckResultStats", ctx, query)
	ret0, _ := ret[0].([]*datastore.ResourceResultCountByClusterScan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceCheckResultStats indicates an expected call of ComplianceCheckResultStats.
func (mr *MockDataStoreMockRecorder) ComplianceCheckResultStats(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceCheckResultStats", reflect.TypeOf((*MockDataStore)(nil).ComplianceCheckResultStats), ctx, query)
}

// ComplianceClusterStats mocks base method.
func (m *MockDataStore) ComplianceClusterStats(ctx context.Context, query *v1.Query) ([]*datastore.ResultStatusCountByCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceClusterStats", ctx, query)
	ret0, _ := ret[0].([]*datastore.ResultStatusCountByCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceClusterStats indicates an expected call of ComplianceClusterStats.
func (mr *MockDataStoreMockRecorder) ComplianceClusterStats(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceClusterStats", reflect.TypeOf((*MockDataStore)(nil).ComplianceClusterStats), ctx, query)
}

// ComplianceProfileResultStats mocks base method.
func (m *MockDataStore) ComplianceProfileResultStats(ctx context.Context, query *v1.Query) ([]*datastore.ResourceResultCountByProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceProfileResultStats", ctx, query)
	ret0, _ := ret[0].([]*datastore.ResourceResultCountByProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceProfileResultStats indicates an expected call of ComplianceProfileResultStats.
func (mr *MockDataStoreMockRecorder) ComplianceProfileResultStats(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceProfileResultStats", reflect.TypeOf((*MockDataStore)(nil).ComplianceProfileResultStats), ctx, query)
}

// ComplianceProfileResults mocks base method.
func (m *MockDataStore) ComplianceProfileResults(ctx context.Context, query *v1.Query) ([]*datastore.ResourceResultsByProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComplianceProfileResults", ctx, query)
	ret0, _ := ret[0].([]*datastore.ResourceResultsByProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComplianceProfileResults indicates an expected call of ComplianceProfileResults.
func (mr *MockDataStoreMockRecorder) ComplianceProfileResults(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplianceProfileResults", reflect.TypeOf((*MockDataStore)(nil).ComplianceProfileResults), ctx, query)
}

// CountByField mocks base method.
func (m *MockDataStore) CountByField(ctx context.Context, query *v1.Query, field search.FieldLabel) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByField", ctx, query, field)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByField indicates an expected call of CountByField.
func (mr *MockDataStoreMockRecorder) CountByField(ctx, query, field any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByField", reflect.TypeOf((*MockDataStore)(nil).CountByField), ctx, query, field)
}

// CountCheckResults mocks base method.
func (m *MockDataStore) CountCheckResults(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountCheckResults", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountCheckResults indicates an expected call of CountCheckResults.
func (mr *MockDataStoreMockRecorder) CountCheckResults(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountCheckResults", reflect.TypeOf((*MockDataStore)(nil).CountCheckResults), ctx, q)
}

// DeleteResult mocks base method.
func (m *MockDataStore) DeleteResult(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResult", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResult indicates an expected call of DeleteResult.
func (mr *MockDataStoreMockRecorder) DeleteResult(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResult", reflect.TypeOf((*MockDataStore)(nil).DeleteResult), ctx, id)
}

// DeleteResultsByCluster mocks base method.
func (m *MockDataStore) DeleteResultsByCluster(ctx context.Context, clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResultsByCluster", ctx, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResultsByCluster indicates an expected call of DeleteResultsByCluster.
func (mr *MockDataStoreMockRecorder) DeleteResultsByCluster(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResultsByCluster", reflect.TypeOf((*MockDataStore)(nil).DeleteResultsByCluster), ctx, clusterID)
}

// GetComplianceCheckResult mocks base method.
func (m *MockDataStore) GetComplianceCheckResult(ctx context.Context, complianceResultID string) (*storage.ComplianceOperatorCheckResultV2, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceCheckResult", ctx, complianceResultID)
	ret0, _ := ret[0].(*storage.ComplianceOperatorCheckResultV2)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetComplianceCheckResult indicates an expected call of GetComplianceCheckResult.
func (mr *MockDataStoreMockRecorder) GetComplianceCheckResult(ctx, complianceResultID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceCheckResult", reflect.TypeOf((*MockDataStore)(nil).GetComplianceCheckResult), ctx, complianceResultID)
}

// SearchComplianceCheckResults mocks base method.
func (m *MockDataStore) SearchComplianceCheckResults(ctx context.Context, query *v1.Query) ([]*storage.ComplianceOperatorCheckResultV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchComplianceCheckResults", ctx, query)
	ret0, _ := ret[0].([]*storage.ComplianceOperatorCheckResultV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchComplianceCheckResults indicates an expected call of SearchComplianceCheckResults.
func (mr *MockDataStoreMockRecorder) SearchComplianceCheckResults(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchComplianceCheckResults", reflect.TypeOf((*MockDataStore)(nil).SearchComplianceCheckResults), ctx, query)
}

// UpsertResult mocks base method.
func (m *MockDataStore) UpsertResult(ctx context.Context, result *storage.ComplianceOperatorCheckResultV2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertResult", ctx, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertResult indicates an expected call of UpsertResult.
func (mr *MockDataStoreMockRecorder) UpsertResult(ctx, result any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertResult", reflect.TypeOf((*MockDataStore)(nil).UpsertResult), ctx, result)
}

// WalkByQuery mocks base method.
func (m *MockDataStore) WalkByQuery(ctx context.Context, query *v1.Query, fn func(*storage.ComplianceOperatorCheckResultV2) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkByQuery", ctx, query, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkByQuery indicates an expected call of WalkByQuery.
func (mr *MockDataStoreMockRecorder) WalkByQuery(ctx, query, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkByQuery", reflect.TypeOf((*MockDataStore)(nil).WalkByQuery), ctx, query, fn)
}
