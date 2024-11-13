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

// DeleteSnapshot mocks base method.
func (m *MockDataStore) DeleteSnapshot(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockDataStoreMockRecorder) DeleteSnapshot(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockDataStore)(nil).DeleteSnapshot), ctx, id)
}

// GetSnapshot mocks base method.
func (m *MockDataStore) GetSnapshot(ctx context.Context, id string) (*storage.ComplianceOperatorReportSnapshotV2, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", ctx, id)
	ret0, _ := ret[0].(*storage.ComplianceOperatorReportSnapshotV2)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnapshot indicates an expected call of GetSnapshot.
func (mr *MockDataStoreMockRecorder) GetSnapshot(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockDataStore)(nil).GetSnapshot), ctx, id)
}

// SearchSnapshots mocks base method.
func (m *MockDataStore) SearchSnapshots(ctx context.Context, query *v1.Query) ([]*storage.ComplianceOperatorReportSnapshotV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchSnapshots", ctx, query)
	ret0, _ := ret[0].([]*storage.ComplianceOperatorReportSnapshotV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchSnapshots indicates an expected call of SearchSnapshots.
func (mr *MockDataStoreMockRecorder) SearchSnapshots(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSnapshots", reflect.TypeOf((*MockDataStore)(nil).SearchSnapshots), ctx, query)
}

// UpsertSnapshot mocks base method.
func (m *MockDataStore) UpsertSnapshot(ctx context.Context, result *storage.ComplianceOperatorReportSnapshotV2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSnapshot", ctx, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSnapshot indicates an expected call of UpsertSnapshot.
func (mr *MockDataStoreMockRecorder) UpsertSnapshot(ctx, result any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSnapshot", reflect.TypeOf((*MockDataStore)(nil).UpsertSnapshot), ctx, result)
}
