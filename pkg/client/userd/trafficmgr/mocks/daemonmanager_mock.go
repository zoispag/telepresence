// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/telepresenceio/telepresence/v2/pkg/client/userd/trafficmgr (interfaces: DaemonManager)

// Package mock_trafficmgr is a generated GoMock package.
package mock_trafficmgr

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	daemon "github.com/telepresenceio/telepresence/rpc/v2/daemon"
)

// MockDaemonManager is a mock of DaemonManager interface.
type MockDaemonManager struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonManagerMockRecorder
}

// MockDaemonManagerMockRecorder is the mock recorder for MockDaemonManager.
type MockDaemonManagerMockRecorder struct {
	mock *MockDaemonManager
}

// NewMockDaemonManager creates a new mock instance.
func NewMockDaemonManager(ctrl *gomock.Controller) *MockDaemonManager {
	mock := &MockDaemonManager{ctrl: ctrl}
	mock.recorder = &MockDaemonManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonManager) EXPECT() *MockDaemonManagerMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockDaemonManager) Client(arg0 context.Context) (io.Closer, daemon.DaemonClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client", arg0)
	ret0, _ := ret[0].(io.Closer)
	ret1, _ := ret[1].(daemon.DaemonClient)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Client indicates an expected call of Client.
func (mr *MockDaemonManagerMockRecorder) Client(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockDaemonManager)(nil).Client), arg0)
}

// IsRunning mocks base method.
func (m *MockDaemonManager) IsRunning(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRunning indicates an expected call of IsRunning.
func (mr *MockDaemonManagerMockRecorder) IsRunning(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockDaemonManager)(nil).IsRunning), arg0)
}
