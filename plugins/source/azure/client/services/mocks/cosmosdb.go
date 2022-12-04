// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: CosmosDBAccountsClient,CosmosDBSQLDatabasesClient,CosmosDBMongoDBDatabasesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	documentdb "github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	gomock "github.com/golang/mock/gomock"
)

// MockCosmosDBAccountsClient is a mock of CosmosDBAccountsClient interface.
type MockCosmosDBAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBAccountsClientMockRecorder
}

// MockCosmosDBAccountsClientMockRecorder is the mock recorder for MockCosmosDBAccountsClient.
type MockCosmosDBAccountsClientMockRecorder struct {
	mock *MockCosmosDBAccountsClient
}

// NewMockCosmosDBAccountsClient creates a new mock instance.
func NewMockCosmosDBAccountsClient(ctrl *gomock.Controller) *MockCosmosDBAccountsClient {
	mock := &MockCosmosDBAccountsClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCosmosDBAccountsClient) EXPECT() *MockCosmosDBAccountsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockCosmosDBAccountsClient) List(arg0 context.Context) (documentdb.DatabaseAccountsListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(documentdb.DatabaseAccountsListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCosmosDBAccountsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCosmosDBAccountsClient)(nil).List), arg0)
}

// MockCosmosDBSQLDatabasesClient is a mock of CosmosDBSQLDatabasesClient interface.
type MockCosmosDBSQLDatabasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBSQLDatabasesClientMockRecorder
}

// MockCosmosDBSQLDatabasesClientMockRecorder is the mock recorder for MockCosmosDBSQLDatabasesClient.
type MockCosmosDBSQLDatabasesClientMockRecorder struct {
	mock *MockCosmosDBSQLDatabasesClient
}

// NewMockCosmosDBSQLDatabasesClient creates a new mock instance.
func NewMockCosmosDBSQLDatabasesClient(ctrl *gomock.Controller) *MockCosmosDBSQLDatabasesClient {
	mock := &MockCosmosDBSQLDatabasesClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBSQLDatabasesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCosmosDBSQLDatabasesClient) EXPECT() *MockCosmosDBSQLDatabasesClientMockRecorder {
	return m.recorder
}

// ListSQLDatabases mocks base method.
func (m *MockCosmosDBSQLDatabasesClient) ListSQLDatabases(arg0 context.Context, arg1, arg2 string) (documentdb.SQLDatabaseListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSQLDatabases", arg0, arg1, arg2)
	ret0, _ := ret[0].(documentdb.SQLDatabaseListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSQLDatabases indicates an expected call of ListSQLDatabases.
func (mr *MockCosmosDBSQLDatabasesClientMockRecorder) ListSQLDatabases(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSQLDatabases", reflect.TypeOf((*MockCosmosDBSQLDatabasesClient)(nil).ListSQLDatabases), arg0, arg1, arg2)
}

// MockCosmosDBMongoDBDatabasesClient is a mock of CosmosDBMongoDBDatabasesClient interface.
type MockCosmosDBMongoDBDatabasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBMongoDBDatabasesClientMockRecorder
}

// MockCosmosDBMongoDBDatabasesClientMockRecorder is the mock recorder for MockCosmosDBMongoDBDatabasesClient.
type MockCosmosDBMongoDBDatabasesClientMockRecorder struct {
	mock *MockCosmosDBMongoDBDatabasesClient
}

// NewMockCosmosDBMongoDBDatabasesClient creates a new mock instance.
func NewMockCosmosDBMongoDBDatabasesClient(ctrl *gomock.Controller) *MockCosmosDBMongoDBDatabasesClient {
	mock := &MockCosmosDBMongoDBDatabasesClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBMongoDBDatabasesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCosmosDBMongoDBDatabasesClient) EXPECT() *MockCosmosDBMongoDBDatabasesClientMockRecorder {
	return m.recorder
}

// ListMongoDBDatabases mocks base method.
func (m *MockCosmosDBMongoDBDatabasesClient) ListMongoDBDatabases(arg0 context.Context, arg1, arg2 string) (documentdb.MongoDBDatabaseListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMongoDBDatabases", arg0, arg1, arg2)
	ret0, _ := ret[0].(documentdb.MongoDBDatabaseListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMongoDBDatabases indicates an expected call of ListMongoDBDatabases.
func (mr *MockCosmosDBMongoDBDatabasesClientMockRecorder) ListMongoDBDatabases(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMongoDBDatabases", reflect.TypeOf((*MockCosmosDBMongoDBDatabasesClient)(nil).ListMongoDBDatabases), arg0, arg1, arg2)
}