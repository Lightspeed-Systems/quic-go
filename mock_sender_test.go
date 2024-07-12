// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: Sender)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_sender_test.go github.com/quic-go/quic-go Sender
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/Lightspeed-Systems/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockSender is a mock of Sender interface.
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender.
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance.
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Available mocks base method.
func (m *MockSender) Available() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Available")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Available indicates an expected call of Available.
func (mr *MockSenderMockRecorder) Available() *MockSenderAvailableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Available", reflect.TypeOf((*MockSender)(nil).Available))
	return &MockSenderAvailableCall{Call: call}
}

// MockSenderAvailableCall wrap *gomock.Call
type MockSenderAvailableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSenderAvailableCall) Return(arg0 <-chan struct{}) *MockSenderAvailableCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSenderAvailableCall) Do(f func() <-chan struct{}) *MockSenderAvailableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSenderAvailableCall) DoAndReturn(f func() <-chan struct{}) *MockSenderAvailableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockSender) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSenderMockRecorder) Close() *MockSenderCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSender)(nil).Close))
	return &MockSenderCloseCall{Call: call}
}

// MockSenderCloseCall wrap *gomock.Call
type MockSenderCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSenderCloseCall) Return() *MockSenderCloseCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSenderCloseCall) Do(f func()) *MockSenderCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSenderCloseCall) DoAndReturn(f func()) *MockSenderCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Run mocks base method.
func (m *MockSender) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockSenderMockRecorder) Run() *MockSenderRunCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockSender)(nil).Run))
	return &MockSenderRunCall{Call: call}
}

// MockSenderRunCall wrap *gomock.Call
type MockSenderRunCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSenderRunCall) Return(arg0 error) *MockSenderRunCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSenderRunCall) Do(f func() error) *MockSenderRunCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSenderRunCall) DoAndReturn(f func() error) *MockSenderRunCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Send mocks base method.
func (m *MockSender) Send(arg0 *packetBuffer, arg1 uint16, arg2 protocol.ECN) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0, arg1, arg2)
}

// Send indicates an expected call of Send.
func (mr *MockSenderMockRecorder) Send(arg0, arg1, arg2 any) *MockSenderSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSender)(nil).Send), arg0, arg1, arg2)
	return &MockSenderSendCall{Call: call}
}

// MockSenderSendCall wrap *gomock.Call
type MockSenderSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSenderSendCall) Return() *MockSenderSendCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSenderSendCall) Do(f func(*packetBuffer, uint16, protocol.ECN)) *MockSenderSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSenderSendCall) DoAndReturn(f func(*packetBuffer, uint16, protocol.ECN)) *MockSenderSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WouldBlock mocks base method.
func (m *MockSender) WouldBlock() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WouldBlock")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WouldBlock indicates an expected call of WouldBlock.
func (mr *MockSenderMockRecorder) WouldBlock() *MockSenderWouldBlockCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WouldBlock", reflect.TypeOf((*MockSender)(nil).WouldBlock))
	return &MockSenderWouldBlockCall{Call: call}
}

// MockSenderWouldBlockCall wrap *gomock.Call
type MockSenderWouldBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSenderWouldBlockCall) Return(arg0 bool) *MockSenderWouldBlockCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSenderWouldBlockCall) Do(f func() bool) *MockSenderWouldBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSenderWouldBlockCall) DoAndReturn(f func() bool) *MockSenderWouldBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
