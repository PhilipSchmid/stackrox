// Code generated by MockGen. DO NOT EDIT.
// Source: types.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/types.go -source types.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	central "github.com/stackrox/rox/generated/internalapi/central"
	storage "github.com/stackrox/rox/generated/storage"
	deduperkey "github.com/stackrox/rox/pkg/deduperkey"
	clusterentities "github.com/stackrox/rox/sensor/common/clusterentities"
	rbac "github.com/stackrox/rox/sensor/common/rbac"
	registry "github.com/stackrox/rox/sensor/common/registry"
	selector "github.com/stackrox/rox/sensor/common/selector"
	service "github.com/stackrox/rox/sensor/common/service"
	store "github.com/stackrox/rox/sensor/common/store"
	gomock "go.uber.org/mock/gomock"
)

// MockDeploymentStore is a mock of DeploymentStore interface.
type MockDeploymentStore struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentStoreMockRecorder
	isgomock struct{}
}

// MockDeploymentStoreMockRecorder is the mock recorder for MockDeploymentStore.
type MockDeploymentStoreMockRecorder struct {
	mock *MockDeploymentStore
}

// NewMockDeploymentStore creates a new mock instance.
func NewMockDeploymentStore(ctrl *gomock.Controller) *MockDeploymentStore {
	mock := &MockDeploymentStore{ctrl: ctrl}
	mock.recorder = &MockDeploymentStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentStore) EXPECT() *MockDeploymentStoreMockRecorder {
	return m.recorder
}

// BuildDeploymentWithDependencies mocks base method.
func (m *MockDeploymentStore) BuildDeploymentWithDependencies(id string, dependencies store.Dependencies) (*storage.Deployment, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildDeploymentWithDependencies", id, dependencies)
	ret0, _ := ret[0].(*storage.Deployment)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BuildDeploymentWithDependencies indicates an expected call of BuildDeploymentWithDependencies.
func (mr *MockDeploymentStoreMockRecorder) BuildDeploymentWithDependencies(id, dependencies any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildDeploymentWithDependencies", reflect.TypeOf((*MockDeploymentStore)(nil).BuildDeploymentWithDependencies), id, dependencies)
}

// CountDeploymentsForNamespace mocks base method.
func (m *MockDeploymentStore) CountDeploymentsForNamespace(namespaceName string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDeploymentsForNamespace", namespaceName)
	ret0, _ := ret[0].(int)
	return ret0
}

// CountDeploymentsForNamespace indicates an expected call of CountDeploymentsForNamespace.
func (mr *MockDeploymentStoreMockRecorder) CountDeploymentsForNamespace(namespaceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDeploymentsForNamespace", reflect.TypeOf((*MockDeploymentStore)(nil).CountDeploymentsForNamespace), namespaceName)
}

// EnhanceDeploymentReadOnly mocks base method.
func (m *MockDeploymentStore) EnhanceDeploymentReadOnly(d *storage.Deployment, dependencies store.Dependencies) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnhanceDeploymentReadOnly", d, dependencies)
}

// EnhanceDeploymentReadOnly indicates an expected call of EnhanceDeploymentReadOnly.
func (mr *MockDeploymentStoreMockRecorder) EnhanceDeploymentReadOnly(d, dependencies any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnhanceDeploymentReadOnly", reflect.TypeOf((*MockDeploymentStore)(nil).EnhanceDeploymentReadOnly), d, dependencies)
}

// FindDeploymentIDsByImages mocks base method.
func (m *MockDeploymentStore) FindDeploymentIDsByImages(arg0 []*storage.Image) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentIDsByImages", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// FindDeploymentIDsByImages indicates an expected call of FindDeploymentIDsByImages.
func (mr *MockDeploymentStoreMockRecorder) FindDeploymentIDsByImages(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentIDsByImages", reflect.TypeOf((*MockDeploymentStore)(nil).FindDeploymentIDsByImages), arg0)
}

// FindDeploymentIDsByLabels mocks base method.
func (m *MockDeploymentStore) FindDeploymentIDsByLabels(namespace string, sel selector.Selector) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentIDsByLabels", namespace, sel)
	ret0, _ := ret[0].([]string)
	return ret0
}

// FindDeploymentIDsByLabels indicates an expected call of FindDeploymentIDsByLabels.
func (mr *MockDeploymentStoreMockRecorder) FindDeploymentIDsByLabels(namespace, sel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentIDsByLabels", reflect.TypeOf((*MockDeploymentStore)(nil).FindDeploymentIDsByLabels), namespace, sel)
}

// FindDeploymentIDsWithServiceAccount mocks base method.
func (m *MockDeploymentStore) FindDeploymentIDsWithServiceAccount(namespace, sa string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeploymentIDsWithServiceAccount", namespace, sa)
	ret0, _ := ret[0].([]string)
	return ret0
}

// FindDeploymentIDsWithServiceAccount indicates an expected call of FindDeploymentIDsWithServiceAccount.
func (mr *MockDeploymentStoreMockRecorder) FindDeploymentIDsWithServiceAccount(namespace, sa any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentIDsWithServiceAccount", reflect.TypeOf((*MockDeploymentStore)(nil).FindDeploymentIDsWithServiceAccount), namespace, sa)
}

// Get mocks base method.
func (m *MockDeploymentStore) Get(id string) *storage.Deployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*storage.Deployment)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockDeploymentStoreMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeploymentStore)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockDeploymentStore) GetAll() []*storage.Deployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*storage.Deployment)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDeploymentStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDeploymentStore)(nil).GetAll))
}

// GetBuiltDeployment mocks base method.
func (m *MockDeploymentStore) GetBuiltDeployment(id string) (*storage.Deployment, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuiltDeployment", id)
	ret0, _ := ret[0].(*storage.Deployment)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetBuiltDeployment indicates an expected call of GetBuiltDeployment.
func (mr *MockDeploymentStoreMockRecorder) GetBuiltDeployment(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuiltDeployment", reflect.TypeOf((*MockDeploymentStore)(nil).GetBuiltDeployment), id)
}

// GetSnapshot mocks base method.
func (m *MockDeploymentStore) GetSnapshot(id string) *storage.Deployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", id)
	ret0, _ := ret[0].(*storage.Deployment)
	return ret0
}

// GetSnapshot indicates an expected call of GetSnapshot.
func (mr *MockDeploymentStoreMockRecorder) GetSnapshot(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockDeploymentStore)(nil).GetSnapshot), id)
}

// MockPodStore is a mock of PodStore interface.
type MockPodStore struct {
	ctrl     *gomock.Controller
	recorder *MockPodStoreMockRecorder
	isgomock struct{}
}

// MockPodStoreMockRecorder is the mock recorder for MockPodStore.
type MockPodStoreMockRecorder struct {
	mock *MockPodStore
}

// NewMockPodStore creates a new mock instance.
func NewMockPodStore(ctrl *gomock.Controller) *MockPodStore {
	mock := &MockPodStore{ctrl: ctrl}
	mock.recorder = &MockPodStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodStore) EXPECT() *MockPodStoreMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockPodStore) GetAll() []*storage.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*storage.Pod)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPodStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPodStore)(nil).GetAll))
}

// GetByName mocks base method.
func (m *MockPodStore) GetByName(podName, namespace string) *storage.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", podName, namespace)
	ret0, _ := ret[0].(*storage.Pod)
	return ret0
}

// GetByName indicates an expected call of GetByName.
func (mr *MockPodStoreMockRecorder) GetByName(podName, namespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockPodStore)(nil).GetByName), podName, namespace)
}

// MockNetworkPolicyStore is a mock of NetworkPolicyStore interface.
type MockNetworkPolicyStore struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkPolicyStoreMockRecorder
	isgomock struct{}
}

// MockNetworkPolicyStoreMockRecorder is the mock recorder for MockNetworkPolicyStore.
type MockNetworkPolicyStoreMockRecorder struct {
	mock *MockNetworkPolicyStore
}

// NewMockNetworkPolicyStore creates a new mock instance.
func NewMockNetworkPolicyStore(ctrl *gomock.Controller) *MockNetworkPolicyStore {
	mock := &MockNetworkPolicyStore{ctrl: ctrl}
	mock.recorder = &MockNetworkPolicyStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkPolicyStore) EXPECT() *MockNetworkPolicyStoreMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockNetworkPolicyStore) All() map[string]*storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].(map[string]*storage.NetworkPolicy)
	return ret0
}

// All indicates an expected call of All.
func (mr *MockNetworkPolicyStoreMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockNetworkPolicyStore)(nil).All))
}

// Delete mocks base method.
func (m *MockNetworkPolicyStore) Delete(ID, ns string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ID, ns)
}

// Delete indicates an expected call of Delete.
func (mr *MockNetworkPolicyStoreMockRecorder) Delete(ID, ns any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Delete), ID, ns)
}

// Find mocks base method.
func (m *MockNetworkPolicyStore) Find(namespace string, labels map[string]string) map[string]*storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", namespace, labels)
	ret0, _ := ret[0].(map[string]*storage.NetworkPolicy)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockNetworkPolicyStoreMockRecorder) Find(namespace, labels any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Find), namespace, labels)
}

// Get mocks base method.
func (m *MockNetworkPolicyStore) Get(id string) *storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*storage.NetworkPolicy)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockNetworkPolicyStoreMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Get), id)
}

// Size mocks base method.
func (m *MockNetworkPolicyStore) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockNetworkPolicyStoreMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Size))
}

// Upsert mocks base method.
func (m *MockNetworkPolicyStore) Upsert(ns *storage.NetworkPolicy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Upsert", ns)
}

// Upsert indicates an expected call of Upsert.
func (mr *MockNetworkPolicyStoreMockRecorder) Upsert(ns any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Upsert), ns)
}

// MockServiceAccountStore is a mock of ServiceAccountStore interface.
type MockServiceAccountStore struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountStoreMockRecorder
	isgomock struct{}
}

// MockServiceAccountStoreMockRecorder is the mock recorder for MockServiceAccountStore.
type MockServiceAccountStoreMockRecorder struct {
	mock *MockServiceAccountStore
}

// NewMockServiceAccountStore creates a new mock instance.
func NewMockServiceAccountStore(ctrl *gomock.Controller) *MockServiceAccountStore {
	mock := &MockServiceAccountStore{ctrl: ctrl}
	mock.recorder = &MockServiceAccountStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountStore) EXPECT() *MockServiceAccountStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockServiceAccountStore) Add(sa *storage.ServiceAccount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", sa)
}

// Add indicates an expected call of Add.
func (mr *MockServiceAccountStoreMockRecorder) Add(sa any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockServiceAccountStore)(nil).Add), sa)
}

// GetImagePullSecrets mocks base method.
func (m *MockServiceAccountStore) GetImagePullSecrets(namespace, name string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagePullSecrets", namespace, name)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetImagePullSecrets indicates an expected call of GetImagePullSecrets.
func (mr *MockServiceAccountStoreMockRecorder) GetImagePullSecrets(namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagePullSecrets", reflect.TypeOf((*MockServiceAccountStore)(nil).GetImagePullSecrets), namespace, name)
}

// Remove mocks base method.
func (m *MockServiceAccountStore) Remove(sa *storage.ServiceAccount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", sa)
}

// Remove indicates an expected call of Remove.
func (mr *MockServiceAccountStoreMockRecorder) Remove(sa any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockServiceAccountStore)(nil).Remove), sa)
}

// MockServiceStore is a mock of ServiceStore interface.
type MockServiceStore struct {
	ctrl     *gomock.Controller
	recorder *MockServiceStoreMockRecorder
	isgomock struct{}
}

// MockServiceStoreMockRecorder is the mock recorder for MockServiceStore.
type MockServiceStoreMockRecorder struct {
	mock *MockServiceStore
}

// NewMockServiceStore creates a new mock instance.
func NewMockServiceStore(ctrl *gomock.Controller) *MockServiceStore {
	mock := &MockServiceStore{ctrl: ctrl}
	mock.recorder = &MockServiceStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceStore) EXPECT() *MockServiceStoreMockRecorder {
	return m.recorder
}

// GetExposureInfos mocks base method.
func (m *MockServiceStore) GetExposureInfos(namespace string, labels map[string]string) []map[service.PortRef][]*storage.PortConfig_ExposureInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExposureInfos", namespace, labels)
	ret0, _ := ret[0].([]map[service.PortRef][]*storage.PortConfig_ExposureInfo)
	return ret0
}

// GetExposureInfos indicates an expected call of GetExposureInfos.
func (mr *MockServiceStoreMockRecorder) GetExposureInfos(namespace, labels any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExposureInfos", reflect.TypeOf((*MockServiceStore)(nil).GetExposureInfos), namespace, labels)
}

// MockRBACStore is a mock of RBACStore interface.
type MockRBACStore struct {
	ctrl     *gomock.Controller
	recorder *MockRBACStoreMockRecorder
	isgomock struct{}
}

// MockRBACStoreMockRecorder is the mock recorder for MockRBACStore.
type MockRBACStoreMockRecorder struct {
	mock *MockRBACStore
}

// NewMockRBACStore creates a new mock instance.
func NewMockRBACStore(ctrl *gomock.Controller) *MockRBACStore {
	mock := &MockRBACStore{ctrl: ctrl}
	mock.recorder = &MockRBACStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRBACStore) EXPECT() *MockRBACStoreMockRecorder {
	return m.recorder
}

// GetPermissionLevelForDeployment mocks base method.
func (m *MockRBACStore) GetPermissionLevelForDeployment(deployment rbac.NamespacedServiceAccount) storage.PermissionLevel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissionLevelForDeployment", deployment)
	ret0, _ := ret[0].(storage.PermissionLevel)
	return ret0
}

// GetPermissionLevelForDeployment indicates an expected call of GetPermissionLevelForDeployment.
func (mr *MockRBACStoreMockRecorder) GetPermissionLevelForDeployment(deployment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionLevelForDeployment", reflect.TypeOf((*MockRBACStore)(nil).GetPermissionLevelForDeployment), deployment)
}

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
	isgomock struct{}
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Deployments mocks base method.
func (m *MockProvider) Deployments() store.DeploymentStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployments")
	ret0, _ := ret[0].(store.DeploymentStore)
	return ret0
}

// Deployments indicates an expected call of Deployments.
func (mr *MockProviderMockRecorder) Deployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockProvider)(nil).Deployments))
}

// EndpointManager mocks base method.
func (m *MockProvider) EndpointManager() store.EndpointManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointManager")
	ret0, _ := ret[0].(store.EndpointManager)
	return ret0
}

// EndpointManager indicates an expected call of EndpointManager.
func (mr *MockProviderMockRecorder) EndpointManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointManager", reflect.TypeOf((*MockProvider)(nil).EndpointManager))
}

// Entities mocks base method.
func (m *MockProvider) Entities() *clusterentities.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entities")
	ret0, _ := ret[0].(*clusterentities.Store)
	return ret0
}

// Entities indicates an expected call of Entities.
func (mr *MockProviderMockRecorder) Entities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entities", reflect.TypeOf((*MockProvider)(nil).Entities))
}

// NetworkPolicies mocks base method.
func (m *MockProvider) NetworkPolicies() store.NetworkPolicyStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkPolicies")
	ret0, _ := ret[0].(store.NetworkPolicyStore)
	return ret0
}

// NetworkPolicies indicates an expected call of NetworkPolicies.
func (mr *MockProviderMockRecorder) NetworkPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkPolicies", reflect.TypeOf((*MockProvider)(nil).NetworkPolicies))
}

// Pods mocks base method.
func (m *MockProvider) Pods() store.PodStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pods")
	ret0, _ := ret[0].(store.PodStore)
	return ret0
}

// Pods indicates an expected call of Pods.
func (mr *MockProviderMockRecorder) Pods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pods", reflect.TypeOf((*MockProvider)(nil).Pods))
}

// RBAC mocks base method.
func (m *MockProvider) RBAC() store.RBACStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RBAC")
	ret0, _ := ret[0].(store.RBACStore)
	return ret0
}

// RBAC indicates an expected call of RBAC.
func (mr *MockProviderMockRecorder) RBAC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RBAC", reflect.TypeOf((*MockProvider)(nil).RBAC))
}

// Registries mocks base method.
func (m *MockProvider) Registries() *registry.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registries")
	ret0, _ := ret[0].(*registry.Store)
	return ret0
}

// Registries indicates an expected call of Registries.
func (mr *MockProviderMockRecorder) Registries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registries", reflect.TypeOf((*MockProvider)(nil).Registries))
}

// ServiceAccounts mocks base method.
func (m *MockProvider) ServiceAccounts() store.ServiceAccountStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceAccounts")
	ret0, _ := ret[0].(store.ServiceAccountStore)
	return ret0
}

// ServiceAccounts indicates an expected call of ServiceAccounts.
func (mr *MockProviderMockRecorder) ServiceAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceAccounts", reflect.TypeOf((*MockProvider)(nil).ServiceAccounts))
}

// Services mocks base method.
func (m *MockProvider) Services() store.ServiceStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Services")
	ret0, _ := ret[0].(store.ServiceStore)
	return ret0
}

// Services indicates an expected call of Services.
func (mr *MockProviderMockRecorder) Services() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockProvider)(nil).Services))
}

// MockEndpointManager is a mock of EndpointManager interface.
type MockEndpointManager struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointManagerMockRecorder
	isgomock struct{}
}

// MockEndpointManagerMockRecorder is the mock recorder for MockEndpointManager.
type MockEndpointManagerMockRecorder struct {
	mock *MockEndpointManager
}

// NewMockEndpointManager creates a new mock instance.
func NewMockEndpointManager(ctrl *gomock.Controller) *MockEndpointManager {
	mock := &MockEndpointManager{ctrl: ctrl}
	mock.recorder = &MockEndpointManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpointManager) EXPECT() *MockEndpointManagerMockRecorder {
	return m.recorder
}

// OnDeploymentCreateOrUpdateByID mocks base method.
func (m *MockEndpointManager) OnDeploymentCreateOrUpdateByID(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnDeploymentCreateOrUpdateByID", id)
}

// OnDeploymentCreateOrUpdateByID indicates an expected call of OnDeploymentCreateOrUpdateByID.
func (mr *MockEndpointManagerMockRecorder) OnDeploymentCreateOrUpdateByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDeploymentCreateOrUpdateByID", reflect.TypeOf((*MockEndpointManager)(nil).OnDeploymentCreateOrUpdateByID), id)
}

// MockNodeStore is a mock of NodeStore interface.
type MockNodeStore struct {
	ctrl     *gomock.Controller
	recorder *MockNodeStoreMockRecorder
	isgomock struct{}
}

// MockNodeStoreMockRecorder is the mock recorder for MockNodeStore.
type MockNodeStoreMockRecorder struct {
	mock *MockNodeStore
}

// NewMockNodeStore creates a new mock instance.
func NewMockNodeStore(ctrl *gomock.Controller) *MockNodeStore {
	mock := &MockNodeStore{ctrl: ctrl}
	mock.recorder = &MockNodeStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeStore) EXPECT() *MockNodeStoreMockRecorder {
	return m.recorder
}

// GetNode mocks base method.
func (m *MockNodeStore) GetNode(nodeName string) *storage.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode", nodeName)
	ret0, _ := ret[0].(*storage.Node)
	return ret0
}

// GetNode indicates an expected call of GetNode.
func (mr *MockNodeStoreMockRecorder) GetNode(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockNodeStore)(nil).GetNode), nodeName)
}

// MockHashReconciler is a mock of HashReconciler interface.
type MockHashReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockHashReconcilerMockRecorder
	isgomock struct{}
}

// MockHashReconcilerMockRecorder is the mock recorder for MockHashReconciler.
type MockHashReconcilerMockRecorder struct {
	mock *MockHashReconciler
}

// NewMockHashReconciler creates a new mock instance.
func NewMockHashReconciler(ctrl *gomock.Controller) *MockHashReconciler {
	mock := &MockHashReconciler{ctrl: ctrl}
	mock.recorder = &MockHashReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashReconciler) EXPECT() *MockHashReconcilerMockRecorder {
	return m.recorder
}

// ProcessHashes mocks base method.
func (m *MockHashReconciler) ProcessHashes(h map[deduperkey.Key]uint64) []central.MsgFromSensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessHashes", h)
	ret0, _ := ret[0].([]central.MsgFromSensor)
	return ret0
}

// ProcessHashes indicates an expected call of ProcessHashes.
func (mr *MockHashReconcilerMockRecorder) ProcessHashes(h any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessHashes", reflect.TypeOf((*MockHashReconciler)(nil).ProcessHashes), h)
}
