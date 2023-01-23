// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/telepresenceio/telepresence/v2/pkg/client/userd (interfaces: Service)

// Package mock_userd is a generated GoMock package.
package mock_userd

import (
	context "context"
	reflect "reflect"

	rpc "github.com/datawire/go-fuseftp/rpc"
	gomock "github.com/golang/mock/gomock"
	manager "github.com/telepresenceio/telepresence/rpc/v2/manager"
	scout "github.com/telepresenceio/telepresence/v2/pkg/client/scout"
	grpc "google.golang.org/grpc"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// As mocks base method.
func (m *MockService) As(arg0 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "As", arg0)
}

// As indicates an expected call of As.
func (mr *MockServiceMockRecorder) As(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "As", reflect.TypeOf((*MockService)(nil).As), arg0)
}

// GetAPIKey mocks base method.
func (m *MockService) GetAPIKey(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIKey", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIKey indicates an expected call of GetAPIKey.
func (mr *MockServiceMockRecorder) GetAPIKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIKey", reflect.TypeOf((*MockService)(nil).GetAPIKey), arg0)
}

// GetFuseFTPClient mocks base method.
func (m *MockService) GetFuseFTPClient(arg0 context.Context) rpc.FuseFTPClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFuseFTPClient", arg0)
	ret0, _ := ret[0].(rpc.FuseFTPClient)
	return ret0
}

// GetFuseFTPClient indicates an expected call of GetFuseFTPClient.
func (mr *MockServiceMockRecorder) GetFuseFTPClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFuseFTPClient", reflect.TypeOf((*MockService)(nil).GetFuseFTPClient), arg0)
}

// Reporter mocks base method.
func (m *MockService) Reporter() *scout.Reporter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reporter")
	ret0, _ := ret[0].(*scout.Reporter)
	return ret0
}

// Reporter indicates an expected call of Reporter.
func (mr *MockServiceMockRecorder) Reporter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reporter", reflect.TypeOf((*MockService)(nil).Reporter))
}

// Server mocks base method.
func (m *MockService) Server() *grpc.Server {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Server")
	ret0, _ := ret[0].(*grpc.Server)
	return ret0
}

// Server indicates an expected call of Server.
func (mr *MockServiceMockRecorder) Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Server", reflect.TypeOf((*MockService)(nil).Server))
}

// SetManagerClient mocks base method.
func (m *MockService) SetManagerClient(arg0 manager.ManagerClient, arg1 ...grpc.CallOption) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetManagerClient", varargs...)
}

// SetManagerClient indicates an expected call of SetManagerClient.
func (mr *MockServiceMockRecorder) SetManagerClient(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagerClient", reflect.TypeOf((*MockService)(nil).SetManagerClient), varargs...)
}
