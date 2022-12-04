// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: StorageAccountsClient,StorageBlobServicesClient,StorageContainersClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	gomock "github.com/golang/mock/gomock"
	accounts "github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
	queues "github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues"
)

// MockStorageAccountsClient is a mock of StorageAccountsClient interface.
type MockStorageAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAccountsClientMockRecorder
}

// MockStorageAccountsClientMockRecorder is the mock recorder for MockStorageAccountsClient.
type MockStorageAccountsClientMockRecorder struct {
	mock *MockStorageAccountsClient
}

// NewMockStorageAccountsClient creates a new mock instance.
func NewMockStorageAccountsClient(ctrl *gomock.Controller) *MockStorageAccountsClient {
	mock := &MockStorageAccountsClient{ctrl: ctrl}
	mock.recorder = &MockStorageAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAccountsClient) EXPECT() *MockStorageAccountsClientMockRecorder {
	return m.recorder
}

// GetBlobServiceProperties mocks base method.
func (m *MockStorageAccountsClient) GetBlobServiceProperties(arg0 context.Context, arg1, arg2 string) (accounts.GetServicePropertiesResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobServiceProperties", arg0, arg1, arg2)
	ret0, _ := ret[0].(accounts.GetServicePropertiesResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlobServiceProperties indicates an expected call of GetBlobServiceProperties.
func (mr *MockStorageAccountsClientMockRecorder) GetBlobServiceProperties(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobServiceProperties", reflect.TypeOf((*MockStorageAccountsClient)(nil).GetBlobServiceProperties), arg0, arg1, arg2)
}

// GetQueueServiceProperties mocks base method.
func (m *MockStorageAccountsClient) GetQueueServiceProperties(arg0 context.Context, arg1, arg2 string) (queues.StorageServicePropertiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueServiceProperties", arg0, arg1, arg2)
	ret0, _ := ret[0].(queues.StorageServicePropertiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueServiceProperties indicates an expected call of GetQueueServiceProperties.
func (mr *MockStorageAccountsClientMockRecorder) GetQueueServiceProperties(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueServiceProperties", reflect.TypeOf((*MockStorageAccountsClient)(nil).GetQueueServiceProperties), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockStorageAccountsClient) List(arg0 context.Context) (storage.AccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(storage.AccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageAccountsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageAccountsClient)(nil).List), arg0)
}

// MockStorageBlobServicesClient is a mock of StorageBlobServicesClient interface.
type MockStorageBlobServicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageBlobServicesClientMockRecorder
}

// MockStorageBlobServicesClientMockRecorder is the mock recorder for MockStorageBlobServicesClient.
type MockStorageBlobServicesClientMockRecorder struct {
	mock *MockStorageBlobServicesClient
}

// NewMockStorageBlobServicesClient creates a new mock instance.
func NewMockStorageBlobServicesClient(ctrl *gomock.Controller) *MockStorageBlobServicesClient {
	mock := &MockStorageBlobServicesClient{ctrl: ctrl}
	mock.recorder = &MockStorageBlobServicesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageBlobServicesClient) EXPECT() *MockStorageBlobServicesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStorageBlobServicesClient) List(arg0 context.Context, arg1, arg2 string) (storage.BlobServiceItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(storage.BlobServiceItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageBlobServicesClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageBlobServicesClient)(nil).List), arg0, arg1, arg2)
}

// MockStorageContainersClient is a mock of StorageContainersClient interface.
type MockStorageContainersClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageContainersClientMockRecorder
}

// MockStorageContainersClientMockRecorder is the mock recorder for MockStorageContainersClient.
type MockStorageContainersClientMockRecorder struct {
	mock *MockStorageContainersClient
}

// NewMockStorageContainersClient creates a new mock instance.
func NewMockStorageContainersClient(ctrl *gomock.Controller) *MockStorageContainersClient {
	mock := &MockStorageContainersClient{ctrl: ctrl}
	mock.recorder = &MockStorageContainersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageContainersClient) EXPECT() *MockStorageContainersClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStorageContainersClient) List(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 storage.ListContainersInclude) (storage.ListContainerItemsPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(storage.ListContainerItemsPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageContainersClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageContainersClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5)
}