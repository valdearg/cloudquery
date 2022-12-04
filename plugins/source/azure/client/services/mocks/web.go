// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: WebAppsClient,WebFunctionsClient,WebSiteAuthSettingsClient,WebSiteAuthSettingsV2Client,WebVnetConnectionsClient,WebPublishingProfilesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	web "github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	services "github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	gomock "github.com/golang/mock/gomock"
)

// MockWebAppsClient is a mock of WebAppsClient interface.
type MockWebAppsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebAppsClientMockRecorder
}

// MockWebAppsClientMockRecorder is the mock recorder for MockWebAppsClient.
type MockWebAppsClientMockRecorder struct {
	mock *MockWebAppsClient
}

// NewMockWebAppsClient creates a new mock instance.
func NewMockWebAppsClient(ctrl *gomock.Controller) *MockWebAppsClient {
	mock := &MockWebAppsClient{ctrl: ctrl}
	mock.recorder = &MockWebAppsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebAppsClient) EXPECT() *MockWebAppsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockWebAppsClient) List(arg0 context.Context) (web.AppCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(web.AppCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWebAppsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWebAppsClient)(nil).List), arg0)
}

// MockWebFunctionsClient is a mock of WebFunctionsClient interface.
type MockWebFunctionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebFunctionsClientMockRecorder
}

// MockWebFunctionsClientMockRecorder is the mock recorder for MockWebFunctionsClient.
type MockWebFunctionsClientMockRecorder struct {
	mock *MockWebFunctionsClient
}

// NewMockWebFunctionsClient creates a new mock instance.
func NewMockWebFunctionsClient(ctrl *gomock.Controller) *MockWebFunctionsClient {
	mock := &MockWebFunctionsClient{ctrl: ctrl}
	mock.recorder = &MockWebFunctionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebFunctionsClient) EXPECT() *MockWebFunctionsClientMockRecorder {
	return m.recorder
}

// ListFunctions mocks base method.
func (m *MockWebFunctionsClient) ListFunctions(arg0 context.Context, arg1, arg2 string) (web.FunctionEnvelopeCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFunctions", arg0, arg1, arg2)
	ret0, _ := ret[0].(web.FunctionEnvelopeCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctions indicates an expected call of ListFunctions.
func (mr *MockWebFunctionsClientMockRecorder) ListFunctions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctions", reflect.TypeOf((*MockWebFunctionsClient)(nil).ListFunctions), arg0, arg1, arg2)
}

// MockWebSiteAuthSettingsClient is a mock of WebSiteAuthSettingsClient interface.
type MockWebSiteAuthSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebSiteAuthSettingsClientMockRecorder
}

// MockWebSiteAuthSettingsClientMockRecorder is the mock recorder for MockWebSiteAuthSettingsClient.
type MockWebSiteAuthSettingsClientMockRecorder struct {
	mock *MockWebSiteAuthSettingsClient
}

// NewMockWebSiteAuthSettingsClient creates a new mock instance.
func NewMockWebSiteAuthSettingsClient(ctrl *gomock.Controller) *MockWebSiteAuthSettingsClient {
	mock := &MockWebSiteAuthSettingsClient{ctrl: ctrl}
	mock.recorder = &MockWebSiteAuthSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebSiteAuthSettingsClient) EXPECT() *MockWebSiteAuthSettingsClientMockRecorder {
	return m.recorder
}

// GetAuthSettings mocks base method.
func (m *MockWebSiteAuthSettingsClient) GetAuthSettings(arg0 context.Context, arg1, arg2 string) (web.SiteAuthSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthSettings", arg0, arg1, arg2)
	ret0, _ := ret[0].(web.SiteAuthSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthSettings indicates an expected call of GetAuthSettings.
func (mr *MockWebSiteAuthSettingsClientMockRecorder) GetAuthSettings(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthSettings", reflect.TypeOf((*MockWebSiteAuthSettingsClient)(nil).GetAuthSettings), arg0, arg1, arg2)
}

// MockWebSiteAuthSettingsV2Client is a mock of WebSiteAuthSettingsV2Client interface.
type MockWebSiteAuthSettingsV2Client struct {
	ctrl     *gomock.Controller
	recorder *MockWebSiteAuthSettingsV2ClientMockRecorder
}

// MockWebSiteAuthSettingsV2ClientMockRecorder is the mock recorder for MockWebSiteAuthSettingsV2Client.
type MockWebSiteAuthSettingsV2ClientMockRecorder struct {
	mock *MockWebSiteAuthSettingsV2Client
}

// NewMockWebSiteAuthSettingsV2Client creates a new mock instance.
func NewMockWebSiteAuthSettingsV2Client(ctrl *gomock.Controller) *MockWebSiteAuthSettingsV2Client {
	mock := &MockWebSiteAuthSettingsV2Client{ctrl: ctrl}
	mock.recorder = &MockWebSiteAuthSettingsV2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebSiteAuthSettingsV2Client) EXPECT() *MockWebSiteAuthSettingsV2ClientMockRecorder {
	return m.recorder
}

// GetAuthSettingsV2 mocks base method.
func (m *MockWebSiteAuthSettingsV2Client) GetAuthSettingsV2(arg0 context.Context, arg1, arg2 string) (web.SiteAuthSettingsV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthSettingsV2", arg0, arg1, arg2)
	ret0, _ := ret[0].(web.SiteAuthSettingsV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthSettingsV2 indicates an expected call of GetAuthSettingsV2.
func (mr *MockWebSiteAuthSettingsV2ClientMockRecorder) GetAuthSettingsV2(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthSettingsV2", reflect.TypeOf((*MockWebSiteAuthSettingsV2Client)(nil).GetAuthSettingsV2), arg0, arg1, arg2)
}

// MockWebVnetConnectionsClient is a mock of WebVnetConnectionsClient interface.
type MockWebVnetConnectionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebVnetConnectionsClientMockRecorder
}

// MockWebVnetConnectionsClientMockRecorder is the mock recorder for MockWebVnetConnectionsClient.
type MockWebVnetConnectionsClientMockRecorder struct {
	mock *MockWebVnetConnectionsClient
}

// NewMockWebVnetConnectionsClient creates a new mock instance.
func NewMockWebVnetConnectionsClient(ctrl *gomock.Controller) *MockWebVnetConnectionsClient {
	mock := &MockWebVnetConnectionsClient{ctrl: ctrl}
	mock.recorder = &MockWebVnetConnectionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebVnetConnectionsClient) EXPECT() *MockWebVnetConnectionsClientMockRecorder {
	return m.recorder
}

// GetVnetConnection mocks base method.
func (m *MockWebVnetConnectionsClient) GetVnetConnection(arg0 context.Context, arg1, arg2, arg3 string) (web.VnetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVnetConnection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(web.VnetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVnetConnection indicates an expected call of GetVnetConnection.
func (mr *MockWebVnetConnectionsClientMockRecorder) GetVnetConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVnetConnection", reflect.TypeOf((*MockWebVnetConnectionsClient)(nil).GetVnetConnection), arg0, arg1, arg2, arg3)
}

// MockWebPublishingProfilesClient is a mock of WebPublishingProfilesClient interface.
type MockWebPublishingProfilesClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebPublishingProfilesClientMockRecorder
}

// MockWebPublishingProfilesClientMockRecorder is the mock recorder for MockWebPublishingProfilesClient.
type MockWebPublishingProfilesClientMockRecorder struct {
	mock *MockWebPublishingProfilesClient
}

// NewMockWebPublishingProfilesClient creates a new mock instance.
func NewMockWebPublishingProfilesClient(ctrl *gomock.Controller) *MockWebPublishingProfilesClient {
	mock := &MockWebPublishingProfilesClient{ctrl: ctrl}
	mock.recorder = &MockWebPublishingProfilesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebPublishingProfilesClient) EXPECT() *MockWebPublishingProfilesClientMockRecorder {
	return m.recorder
}

// ListPublishingProfiles mocks base method.
func (m *MockWebPublishingProfilesClient) ListPublishingProfiles(arg0 context.Context, arg1, arg2 string) (services.PublishingProfiles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublishingProfiles", arg0, arg1, arg2)
	ret0, _ := ret[0].(services.PublishingProfiles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublishingProfiles indicates an expected call of ListPublishingProfiles.
func (mr *MockWebPublishingProfilesClientMockRecorder) ListPublishingProfiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublishingProfiles", reflect.TypeOf((*MockWebPublishingProfilesClient)(nil).ListPublishingProfiles), arg0, arg1, arg2)
}