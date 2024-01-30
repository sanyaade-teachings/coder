// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coder/coder/v2/tailnet (interfaces: MultiAgentConn)
//
// Generated by this command:
//
//	mockgen -destination ./multiagentmock.go -package tailnettest github.com/coder/coder/v2/tailnet MultiAgentConn
//

// Package tailnettest is a generated GoMock package.
package tailnettest

import (
	context "context"
	reflect "reflect"

	proto "github.com/coder/coder/v2/tailnet/proto"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockMultiAgentConn is a mock of MultiAgentConn interface.
type MockMultiAgentConn struct {
	ctrl     *gomock.Controller
	recorder *MockMultiAgentConnMockRecorder
}

// MockMultiAgentConnMockRecorder is the mock recorder for MockMultiAgentConn.
type MockMultiAgentConnMockRecorder struct {
	mock *MockMultiAgentConn
}

// NewMockMultiAgentConn creates a new mock instance.
func NewMockMultiAgentConn(ctrl *gomock.Controller) *MockMultiAgentConn {
	mock := &MockMultiAgentConn{ctrl: ctrl}
	mock.recorder = &MockMultiAgentConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultiAgentConn) EXPECT() *MockMultiAgentConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMultiAgentConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMultiAgentConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMultiAgentConn)(nil).Close))
}

// IsClosed mocks base method.
func (m *MockMultiAgentConn) IsClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockMultiAgentConnMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockMultiAgentConn)(nil).IsClosed))
}

// NextUpdate mocks base method.
func (m *MockMultiAgentConn) NextUpdate(arg0 context.Context) (*proto.CoordinateResponse, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextUpdate", arg0)
	ret0, _ := ret[0].(*proto.CoordinateResponse)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// NextUpdate indicates an expected call of NextUpdate.
func (mr *MockMultiAgentConnMockRecorder) NextUpdate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextUpdate", reflect.TypeOf((*MockMultiAgentConn)(nil).NextUpdate), arg0)
}

// SubscribeAgent mocks base method.
func (m *MockMultiAgentConn) SubscribeAgent(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeAgent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeAgent indicates an expected call of SubscribeAgent.
func (mr *MockMultiAgentConnMockRecorder) SubscribeAgent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeAgent", reflect.TypeOf((*MockMultiAgentConn)(nil).SubscribeAgent), arg0)
}

// UnsubscribeAgent mocks base method.
func (m *MockMultiAgentConn) UnsubscribeAgent(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeAgent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeAgent indicates an expected call of UnsubscribeAgent.
func (mr *MockMultiAgentConnMockRecorder) UnsubscribeAgent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeAgent", reflect.TypeOf((*MockMultiAgentConn)(nil).UnsubscribeAgent), arg0)
}

// UpdateSelf mocks base method.
func (m *MockMultiAgentConn) UpdateSelf(arg0 *proto.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSelf", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSelf indicates an expected call of UpdateSelf.
func (mr *MockMultiAgentConnMockRecorder) UpdateSelf(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSelf", reflect.TypeOf((*MockMultiAgentConn)(nil).UpdateSelf), arg0)
}