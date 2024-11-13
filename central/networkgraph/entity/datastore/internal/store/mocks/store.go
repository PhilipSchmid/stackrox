// Code generated by MockGen. DO NOT EDIT.
// Source: store.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/store.go -source store.go
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

// MockEntityStore is a mock of EntityStore interface.
type MockEntityStore struct {
	ctrl     *gomock.Controller
	recorder *MockEntityStoreMockRecorder
	isgomock struct{}
}

// MockEntityStoreMockRecorder is the mock recorder for MockEntityStore.
type MockEntityStoreMockRecorder struct {
	mock *MockEntityStore
}

// NewMockEntityStore creates a new mock instance.
func NewMockEntityStore(ctrl *gomock.Controller) *MockEntityStore {
	mock := &MockEntityStore{ctrl: ctrl}
	mock.recorder = &MockEntityStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityStore) EXPECT() *MockEntityStoreMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockEntityStore) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEntityStoreMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEntityStore)(nil).Delete), ctx, id)
}

// DeleteMany mocks base method.
func (m *MockEntityStore) DeleteMany(ctx context.Context, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMany", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMany indicates an expected call of DeleteMany.
func (mr *MockEntityStoreMockRecorder) DeleteMany(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMany", reflect.TypeOf((*MockEntityStore)(nil).DeleteMany), ctx, ids)
}

// Exists mocks base method.
func (m *MockEntityStore) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockEntityStoreMockRecorder) Exists(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockEntityStore)(nil).Exists), ctx, id)
}

// Get mocks base method.
func (m *MockEntityStore) Get(ctx context.Context, id string) (*storage.NetworkEntity, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.NetworkEntity)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockEntityStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEntityStore)(nil).Get), ctx, id)
}

// GetByQuery mocks base method.
func (m *MockEntityStore) GetByQuery(ctx context.Context, query *v1.Query) ([]*storage.NetworkEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByQuery", ctx, query)
	ret0, _ := ret[0].([]*storage.NetworkEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByQuery indicates an expected call of GetByQuery.
func (mr *MockEntityStoreMockRecorder) GetByQuery(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByQuery", reflect.TypeOf((*MockEntityStore)(nil).GetByQuery), ctx, query)
}

// GetIDs mocks base method.
func (m *MockEntityStore) GetIDs(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDs", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDs indicates an expected call of GetIDs.
func (mr *MockEntityStoreMockRecorder) GetIDs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDs", reflect.TypeOf((*MockEntityStore)(nil).GetIDs), ctx)
}

// Upsert mocks base method.
func (m *MockEntityStore) Upsert(ctx context.Context, entity *storage.NetworkEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockEntityStoreMockRecorder) Upsert(ctx, entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockEntityStore)(nil).Upsert), ctx, entity)
}

// UpsertMany mocks base method.
func (m *MockEntityStore) UpsertMany(ctx context.Context, objs []*storage.NetworkEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMany", ctx, objs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMany indicates an expected call of UpsertMany.
func (mr *MockEntityStoreMockRecorder) UpsertMany(ctx, objs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMany", reflect.TypeOf((*MockEntityStore)(nil).UpsertMany), ctx, objs)
}

// Walk mocks base method.
func (m *MockEntityStore) Walk(ctx context.Context, fn func(*storage.NetworkEntity) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockEntityStoreMockRecorder) Walk(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockEntityStore)(nil).Walk), ctx, fn)
}

// WalkByQuery mocks base method.
func (m *MockEntityStore) WalkByQuery(ctx context.Context, query *v1.Query, fn func(*storage.NetworkEntity) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkByQuery", ctx, query, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkByQuery indicates an expected call of WalkByQuery.
func (mr *MockEntityStoreMockRecorder) WalkByQuery(ctx, query, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkByQuery", reflect.TypeOf((*MockEntityStore)(nil).WalkByQuery), ctx, query, fn)
}
