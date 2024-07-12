// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go/internal/flowcontrol (interfaces: StreamFlowController)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mocks -destination stream_flow_controller.go github.com/quic-go/quic-go/internal/flowcontrol StreamFlowController
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	protocol "github.com/Lightspeed-Systems/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockStreamFlowController is a mock of StreamFlowController interface.
type MockStreamFlowController struct {
	ctrl     *gomock.Controller
	recorder *MockStreamFlowControllerMockRecorder
}

// MockStreamFlowControllerMockRecorder is the mock recorder for MockStreamFlowController.
type MockStreamFlowControllerMockRecorder struct {
	mock *MockStreamFlowController
}

// NewMockStreamFlowController creates a new mock instance.
func NewMockStreamFlowController(ctrl *gomock.Controller) *MockStreamFlowController {
	mock := &MockStreamFlowController{ctrl: ctrl}
	mock.recorder = &MockStreamFlowControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamFlowController) EXPECT() *MockStreamFlowControllerMockRecorder {
	return m.recorder
}

// Abandon mocks base method.
func (m *MockStreamFlowController) Abandon() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abandon")
}

// Abandon indicates an expected call of Abandon.
func (mr *MockStreamFlowControllerMockRecorder) Abandon() *MockStreamFlowControllerAbandonCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abandon", reflect.TypeOf((*MockStreamFlowController)(nil).Abandon))
	return &MockStreamFlowControllerAbandonCall{Call: call}
}

// MockStreamFlowControllerAbandonCall wrap *gomock.Call
type MockStreamFlowControllerAbandonCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerAbandonCall) Return() *MockStreamFlowControllerAbandonCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerAbandonCall) Do(f func()) *MockStreamFlowControllerAbandonCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerAbandonCall) DoAndReturn(f func()) *MockStreamFlowControllerAbandonCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddBytesRead mocks base method.
func (m *MockStreamFlowController) AddBytesRead(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBytesRead", arg0)
}

// AddBytesRead indicates an expected call of AddBytesRead.
func (mr *MockStreamFlowControllerMockRecorder) AddBytesRead(arg0 any) *MockStreamFlowControllerAddBytesReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytesRead", reflect.TypeOf((*MockStreamFlowController)(nil).AddBytesRead), arg0)
	return &MockStreamFlowControllerAddBytesReadCall{Call: call}
}

// MockStreamFlowControllerAddBytesReadCall wrap *gomock.Call
type MockStreamFlowControllerAddBytesReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerAddBytesReadCall) Return() *MockStreamFlowControllerAddBytesReadCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerAddBytesReadCall) Do(f func(protocol.ByteCount)) *MockStreamFlowControllerAddBytesReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerAddBytesReadCall) DoAndReturn(f func(protocol.ByteCount)) *MockStreamFlowControllerAddBytesReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddBytesSent mocks base method.
func (m *MockStreamFlowController) AddBytesSent(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBytesSent", arg0)
}

// AddBytesSent indicates an expected call of AddBytesSent.
func (mr *MockStreamFlowControllerMockRecorder) AddBytesSent(arg0 any) *MockStreamFlowControllerAddBytesSentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytesSent", reflect.TypeOf((*MockStreamFlowController)(nil).AddBytesSent), arg0)
	return &MockStreamFlowControllerAddBytesSentCall{Call: call}
}

// MockStreamFlowControllerAddBytesSentCall wrap *gomock.Call
type MockStreamFlowControllerAddBytesSentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerAddBytesSentCall) Return() *MockStreamFlowControllerAddBytesSentCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerAddBytesSentCall) Do(f func(protocol.ByteCount)) *MockStreamFlowControllerAddBytesSentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerAddBytesSentCall) DoAndReturn(f func(protocol.ByteCount)) *MockStreamFlowControllerAddBytesSentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetWindowUpdate mocks base method.
func (m *MockStreamFlowController) GetWindowUpdate() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetWindowUpdate indicates an expected call of GetWindowUpdate.
func (mr *MockStreamFlowControllerMockRecorder) GetWindowUpdate() *MockStreamFlowControllerGetWindowUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWindowUpdate", reflect.TypeOf((*MockStreamFlowController)(nil).GetWindowUpdate))
	return &MockStreamFlowControllerGetWindowUpdateCall{Call: call}
}

// MockStreamFlowControllerGetWindowUpdateCall wrap *gomock.Call
type MockStreamFlowControllerGetWindowUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerGetWindowUpdateCall) Return(arg0 protocol.ByteCount) *MockStreamFlowControllerGetWindowUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerGetWindowUpdateCall) Do(f func() protocol.ByteCount) *MockStreamFlowControllerGetWindowUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerGetWindowUpdateCall) DoAndReturn(f func() protocol.ByteCount) *MockStreamFlowControllerGetWindowUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNewlyBlocked mocks base method.
func (m *MockStreamFlowController) IsNewlyBlocked() (bool, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNewlyBlocked")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// IsNewlyBlocked indicates an expected call of IsNewlyBlocked.
func (mr *MockStreamFlowControllerMockRecorder) IsNewlyBlocked() *MockStreamFlowControllerIsNewlyBlockedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNewlyBlocked", reflect.TypeOf((*MockStreamFlowController)(nil).IsNewlyBlocked))
	return &MockStreamFlowControllerIsNewlyBlockedCall{Call: call}
}

// MockStreamFlowControllerIsNewlyBlockedCall wrap *gomock.Call
type MockStreamFlowControllerIsNewlyBlockedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerIsNewlyBlockedCall) Return(arg0 bool, arg1 protocol.ByteCount) *MockStreamFlowControllerIsNewlyBlockedCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerIsNewlyBlockedCall) Do(f func() (bool, protocol.ByteCount)) *MockStreamFlowControllerIsNewlyBlockedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerIsNewlyBlockedCall) DoAndReturn(f func() (bool, protocol.ByteCount)) *MockStreamFlowControllerIsNewlyBlockedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendWindowSize mocks base method.
func (m *MockStreamFlowController) SendWindowSize() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendWindowSize")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// SendWindowSize indicates an expected call of SendWindowSize.
func (mr *MockStreamFlowControllerMockRecorder) SendWindowSize() *MockStreamFlowControllerSendWindowSizeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendWindowSize", reflect.TypeOf((*MockStreamFlowController)(nil).SendWindowSize))
	return &MockStreamFlowControllerSendWindowSizeCall{Call: call}
}

// MockStreamFlowControllerSendWindowSizeCall wrap *gomock.Call
type MockStreamFlowControllerSendWindowSizeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerSendWindowSizeCall) Return(arg0 protocol.ByteCount) *MockStreamFlowControllerSendWindowSizeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerSendWindowSizeCall) Do(f func() protocol.ByteCount) *MockStreamFlowControllerSendWindowSizeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerSendWindowSizeCall) DoAndReturn(f func() protocol.ByteCount) *MockStreamFlowControllerSendWindowSizeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateHighestReceived mocks base method.
func (m *MockStreamFlowController) UpdateHighestReceived(arg0 protocol.ByteCount, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHighestReceived", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHighestReceived indicates an expected call of UpdateHighestReceived.
func (mr *MockStreamFlowControllerMockRecorder) UpdateHighestReceived(arg0, arg1 any) *MockStreamFlowControllerUpdateHighestReceivedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHighestReceived", reflect.TypeOf((*MockStreamFlowController)(nil).UpdateHighestReceived), arg0, arg1)
	return &MockStreamFlowControllerUpdateHighestReceivedCall{Call: call}
}

// MockStreamFlowControllerUpdateHighestReceivedCall wrap *gomock.Call
type MockStreamFlowControllerUpdateHighestReceivedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerUpdateHighestReceivedCall) Return(arg0 error) *MockStreamFlowControllerUpdateHighestReceivedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerUpdateHighestReceivedCall) Do(f func(protocol.ByteCount, bool) error) *MockStreamFlowControllerUpdateHighestReceivedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerUpdateHighestReceivedCall) DoAndReturn(f func(protocol.ByteCount, bool) error) *MockStreamFlowControllerUpdateHighestReceivedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateSendWindow mocks base method.
func (m *MockStreamFlowController) UpdateSendWindow(arg0 protocol.ByteCount) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSendWindow", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateSendWindow indicates an expected call of UpdateSendWindow.
func (mr *MockStreamFlowControllerMockRecorder) UpdateSendWindow(arg0 any) *MockStreamFlowControllerUpdateSendWindowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSendWindow", reflect.TypeOf((*MockStreamFlowController)(nil).UpdateSendWindow), arg0)
	return &MockStreamFlowControllerUpdateSendWindowCall{Call: call}
}

// MockStreamFlowControllerUpdateSendWindowCall wrap *gomock.Call
type MockStreamFlowControllerUpdateSendWindowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamFlowControllerUpdateSendWindowCall) Return(arg0 bool) *MockStreamFlowControllerUpdateSendWindowCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamFlowControllerUpdateSendWindowCall) Do(f func(protocol.ByteCount) bool) *MockStreamFlowControllerUpdateSendWindowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamFlowControllerUpdateSendWindowCall) DoAndReturn(f func(protocol.ByteCount) bool) *MockStreamFlowControllerUpdateSendWindowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
