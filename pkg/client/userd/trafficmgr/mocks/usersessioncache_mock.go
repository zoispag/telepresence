// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/telepresenceio/telepresence/v2/pkg/client/userd/trafficmgr (interfaces: UserSessionCache)

// Package mock_trafficmgr is a generated GoMock package.
package mock_trafficmgr

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	manager "github.com/telepresenceio/telepresence/rpc/v2/manager"
)

// MockUserSessionCache is a mock of UserSessionCache interface.
type MockUserSessionCache struct {
	ctrl     *gomock.Controller
	recorder *MockUserSessionCacheMockRecorder
}

// MockUserSessionCacheMockRecorder is the mock recorder for MockUserSessionCache.
type MockUserSessionCacheMockRecorder struct {
	mock *MockUserSessionCache
}

// NewMockUserSessionCache creates a new mock instance.
func NewMockUserSessionCache(ctrl *gomock.Controller) *MockUserSessionCache {
	mock := &MockUserSessionCache{ctrl: ctrl}
	mock.recorder = &MockUserSessionCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSessionCache) EXPECT() *MockUserSessionCacheMockRecorder {
	return m.recorder
}

// DeleteSessionInfoFromUserCache mocks base method.
func (m *MockUserSessionCache) DeleteSessionInfoFromUserCache(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionInfoFromUserCache", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionInfoFromUserCache indicates an expected call of DeleteSessionInfoFromUserCache.
func (mr *MockUserSessionCacheMockRecorder) DeleteSessionInfoFromUserCache(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionInfoFromUserCache", reflect.TypeOf((*MockUserSessionCache)(nil).DeleteSessionInfoFromUserCache), arg0)
}

// LoadSessionInfoFromUserCache mocks base method.
func (m *MockUserSessionCache) LoadSessionInfoFromUserCache(arg0 context.Context, arg1 string) (*manager.SessionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSessionInfoFromUserCache", arg0, arg1)
	ret0, _ := ret[0].(*manager.SessionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadSessionInfoFromUserCache indicates an expected call of LoadSessionInfoFromUserCache.
func (mr *MockUserSessionCacheMockRecorder) LoadSessionInfoFromUserCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSessionInfoFromUserCache", reflect.TypeOf((*MockUserSessionCache)(nil).LoadSessionInfoFromUserCache), arg0, arg1)
}

// SaveSessionInfoToUserCache mocks base method.
func (m *MockUserSessionCache) SaveSessionInfoToUserCache(arg0 context.Context, arg1 string, arg2 *manager.SessionInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSessionInfoToUserCache", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSessionInfoToUserCache indicates an expected call of SaveSessionInfoToUserCache.
func (mr *MockUserSessionCacheMockRecorder) SaveSessionInfoToUserCache(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSessionInfoToUserCache", reflect.TypeOf((*MockUserSessionCache)(nil).SaveSessionInfoToUserCache), arg0, arg1, arg2)
}
