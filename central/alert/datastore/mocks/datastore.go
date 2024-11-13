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
	search "github.com/stackrox/rox/pkg/search"
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

// Count mocks base method.
func (m *MockDataStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDataStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// CountAlerts mocks base method.
func (m *MockDataStore) CountAlerts(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAlerts", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAlerts indicates an expected call of CountAlerts.
func (mr *MockDataStoreMockRecorder) CountAlerts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAlerts", reflect.TypeOf((*MockDataStore)(nil).CountAlerts), ctx)
}

// DeleteAlerts mocks base method.
func (m *MockDataStore) DeleteAlerts(ctx context.Context, ids ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAlerts", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlerts indicates an expected call of DeleteAlerts.
func (mr *MockDataStoreMockRecorder) DeleteAlerts(ctx any, ids ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlerts", reflect.TypeOf((*MockDataStore)(nil).DeleteAlerts), varargs...)
}

// GetAlert mocks base method.
func (m *MockDataStore) GetAlert(ctx context.Context, id string) (*storage.Alert, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlert", ctx, id)
	ret0, _ := ret[0].(*storage.Alert)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlert indicates an expected call of GetAlert.
func (mr *MockDataStoreMockRecorder) GetAlert(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlert", reflect.TypeOf((*MockDataStore)(nil).GetAlert), ctx, id)
}

// GetByQuery mocks base method.
func (m *MockDataStore) GetByQuery(ctx context.Context, q *v1.Query) ([]*storage.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByQuery", ctx, q)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByQuery indicates an expected call of GetByQuery.
func (mr *MockDataStoreMockRecorder) GetByQuery(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByQuery", reflect.TypeOf((*MockDataStore)(nil).GetByQuery), ctx, q)
}

// MarkAlertsResolvedBatch mocks base method.
func (m *MockDataStore) MarkAlertsResolvedBatch(ctx context.Context, id ...string) ([]*storage.Alert, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range id {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MarkAlertsResolvedBatch", varargs...)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkAlertsResolvedBatch indicates an expected call of MarkAlertsResolvedBatch.
func (mr *MockDataStoreMockRecorder) MarkAlertsResolvedBatch(ctx any, id ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, id...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAlertsResolvedBatch", reflect.TypeOf((*MockDataStore)(nil).MarkAlertsResolvedBatch), varargs...)
}

// PruneAlerts mocks base method.
func (m *MockDataStore) PruneAlerts(ctx context.Context, ids ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PruneAlerts", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PruneAlerts indicates an expected call of PruneAlerts.
func (mr *MockDataStoreMockRecorder) PruneAlerts(ctx any, ids ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneAlerts", reflect.TypeOf((*MockDataStore)(nil).PruneAlerts), varargs...)
}

// Search mocks base method.
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDataStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchAlerts mocks base method.
func (m *MockDataStore) SearchAlerts(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAlerts", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAlerts indicates an expected call of SearchAlerts.
func (mr *MockDataStoreMockRecorder) SearchAlerts(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchAlerts), ctx, q)
}

// SearchListAlerts mocks base method.
func (m *MockDataStore) SearchListAlerts(ctx context.Context, q *v1.Query) ([]*storage.ListAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchListAlerts", ctx, q)
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchListAlerts indicates an expected call of SearchListAlerts.
func (mr *MockDataStoreMockRecorder) SearchListAlerts(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchListAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchListAlerts), ctx, q)
}

// SearchRawAlerts mocks base method.
func (m *MockDataStore) SearchRawAlerts(ctx context.Context, q *v1.Query) ([]*storage.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawAlerts", ctx, q)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawAlerts indicates an expected call of SearchRawAlerts.
func (mr *MockDataStoreMockRecorder) SearchRawAlerts(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchRawAlerts), ctx, q)
}

// UpsertAlert mocks base method.
func (m *MockDataStore) UpsertAlert(ctx context.Context, alert *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertAlert", ctx, alert)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAlert indicates an expected call of UpsertAlert.
func (mr *MockDataStoreMockRecorder) UpsertAlert(ctx, alert any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAlert", reflect.TypeOf((*MockDataStore)(nil).UpsertAlert), ctx, alert)
}

// UpsertAlerts mocks base method.
func (m *MockDataStore) UpsertAlerts(ctx context.Context, alerts []*storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertAlerts", ctx, alerts)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAlerts indicates an expected call of UpsertAlerts.
func (mr *MockDataStoreMockRecorder) UpsertAlerts(ctx, alerts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAlerts", reflect.TypeOf((*MockDataStore)(nil).UpsertAlerts), ctx, alerts)
}

// WalkAll mocks base method.
func (m *MockDataStore) WalkAll(ctx context.Context, fn func(*storage.ListAlert) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkAll", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkAll indicates an expected call of WalkAll.
func (mr *MockDataStoreMockRecorder) WalkAll(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkAll", reflect.TypeOf((*MockDataStore)(nil).WalkAll), ctx, fn)
}

// WalkByQuery mocks base method.
func (m *MockDataStore) WalkByQuery(ctx context.Context, q *v1.Query, db func(*storage.Alert) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkByQuery", ctx, q, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkByQuery indicates an expected call of WalkByQuery.
func (mr *MockDataStoreMockRecorder) WalkByQuery(ctx, q, db any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkByQuery", reflect.TypeOf((*MockDataStore)(nil).WalkByQuery), ctx, q, db)
}
