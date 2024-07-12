// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: CryptoStream)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_crypto_stream_test.go github.com/quic-go/quic-go CryptoStream
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/Lightspeed-Systems/quic-go/internal/protocol"
	wire "github.com/Lightspeed-Systems/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockCryptoStream is a mock of CryptoStream interface.
type MockCryptoStream struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoStreamMockRecorder
}

// MockCryptoStreamMockRecorder is the mock recorder for MockCryptoStream.
type MockCryptoStreamMockRecorder struct {
	mock *MockCryptoStream
}

// NewMockCryptoStream creates a new mock instance.
func NewMockCryptoStream(ctrl *gomock.Controller) *MockCryptoStream {
	mock := &MockCryptoStream{ctrl: ctrl}
	mock.recorder = &MockCryptoStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoStream) EXPECT() *MockCryptoStreamMockRecorder {
	return m.recorder
}

// Finish mocks base method.
func (m *MockCryptoStream) Finish() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finish")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finish indicates an expected call of Finish.
func (mr *MockCryptoStreamMockRecorder) Finish() *MockCryptoStreamFinishCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockCryptoStream)(nil).Finish))
	return &MockCryptoStreamFinishCall{Call: call}
}

// MockCryptoStreamFinishCall wrap *gomock.Call
type MockCryptoStreamFinishCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoStreamFinishCall) Return(arg0 error) *MockCryptoStreamFinishCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoStreamFinishCall) Do(f func() error) *MockCryptoStreamFinishCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoStreamFinishCall) DoAndReturn(f func() error) *MockCryptoStreamFinishCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCryptoData mocks base method.
func (m *MockCryptoStream) GetCryptoData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCryptoData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCryptoData indicates an expected call of GetCryptoData.
func (mr *MockCryptoStreamMockRecorder) GetCryptoData() *MockCryptoStreamGetCryptoDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCryptoData", reflect.TypeOf((*MockCryptoStream)(nil).GetCryptoData))
	return &MockCryptoStreamGetCryptoDataCall{Call: call}
}

// MockCryptoStreamGetCryptoDataCall wrap *gomock.Call
type MockCryptoStreamGetCryptoDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoStreamGetCryptoDataCall) Return(arg0 []byte) *MockCryptoStreamGetCryptoDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoStreamGetCryptoDataCall) Do(f func() []byte) *MockCryptoStreamGetCryptoDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoStreamGetCryptoDataCall) DoAndReturn(f func() []byte) *MockCryptoStreamGetCryptoDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandleCryptoFrame mocks base method.
func (m *MockCryptoStream) HandleCryptoFrame(arg0 *wire.CryptoFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleCryptoFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleCryptoFrame indicates an expected call of HandleCryptoFrame.
func (mr *MockCryptoStreamMockRecorder) HandleCryptoFrame(arg0 any) *MockCryptoStreamHandleCryptoFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCryptoFrame", reflect.TypeOf((*MockCryptoStream)(nil).HandleCryptoFrame), arg0)
	return &MockCryptoStreamHandleCryptoFrameCall{Call: call}
}

// MockCryptoStreamHandleCryptoFrameCall wrap *gomock.Call
type MockCryptoStreamHandleCryptoFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoStreamHandleCryptoFrameCall) Return(arg0 error) *MockCryptoStreamHandleCryptoFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoStreamHandleCryptoFrameCall) Do(f func(*wire.CryptoFrame) error) *MockCryptoStreamHandleCryptoFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoStreamHandleCryptoFrameCall) DoAndReturn(f func(*wire.CryptoFrame) error) *MockCryptoStreamHandleCryptoFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HasData mocks base method.
func (m *MockCryptoStream) HasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasData indicates an expected call of HasData.
func (mr *MockCryptoStreamMockRecorder) HasData() *MockCryptoStreamHasDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasData", reflect.TypeOf((*MockCryptoStream)(nil).HasData))
	return &MockCryptoStreamHasDataCall{Call: call}
}

// MockCryptoStreamHasDataCall wrap *gomock.Call
type MockCryptoStreamHasDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoStreamHasDataCall) Return(arg0 bool) *MockCryptoStreamHasDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoStreamHasDataCall) Do(f func() bool) *MockCryptoStreamHasDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoStreamHasDataCall) DoAndReturn(f func() bool) *MockCryptoStreamHasDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PopCryptoFrame mocks base method.
func (m *MockCryptoStream) PopCryptoFrame(arg0 protocol.ByteCount) *wire.CryptoFrame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopCryptoFrame", arg0)
	ret0, _ := ret[0].(*wire.CryptoFrame)
	return ret0
}

// PopCryptoFrame indicates an expected call of PopCryptoFrame.
func (mr *MockCryptoStreamMockRecorder) PopCryptoFrame(arg0 any) *MockCryptoStreamPopCryptoFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopCryptoFrame", reflect.TypeOf((*MockCryptoStream)(nil).PopCryptoFrame), arg0)
	return &MockCryptoStreamPopCryptoFrameCall{Call: call}
}

// MockCryptoStreamPopCryptoFrameCall wrap *gomock.Call
type MockCryptoStreamPopCryptoFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoStreamPopCryptoFrameCall) Return(arg0 *wire.CryptoFrame) *MockCryptoStreamPopCryptoFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoStreamPopCryptoFrameCall) Do(f func(protocol.ByteCount) *wire.CryptoFrame) *MockCryptoStreamPopCryptoFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoStreamPopCryptoFrameCall) DoAndReturn(f func(protocol.ByteCount) *wire.CryptoFrame) *MockCryptoStreamPopCryptoFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockCryptoStream) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockCryptoStreamMockRecorder) Write(arg0 any) *MockCryptoStreamWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCryptoStream)(nil).Write), arg0)
	return &MockCryptoStreamWriteCall{Call: call}
}

// MockCryptoStreamWriteCall wrap *gomock.Call
type MockCryptoStreamWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCryptoStreamWriteCall) Return(arg0 int, arg1 error) *MockCryptoStreamWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCryptoStreamWriteCall) Do(f func([]byte) (int, error)) *MockCryptoStreamWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCryptoStreamWriteCall) DoAndReturn(f func([]byte) (int, error)) *MockCryptoStreamWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
