// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// WebsocketDataProcessing is an autogenerated mock type for the WebsocketDataProcessing type
type WebsocketDataProcessing struct {
	mock.Mock
}

type WebsocketDataProcessing_Expecter struct {
	mock *mock.Mock
}

func (_m *WebsocketDataProcessing) EXPECT() *WebsocketDataProcessing_Expecter {
	return &WebsocketDataProcessing_Expecter{mock: &_m.Mock}
}

// HandleIncomingShipMessage provides a mock function with given fields: _a0
func (_m *WebsocketDataProcessing) HandleIncomingShipMessage(_a0 []byte) {
	_m.Called(_a0)
}

// WebsocketDataProcessing_HandleIncomingShipMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleIncomingShipMessage'
type WebsocketDataProcessing_HandleIncomingShipMessage_Call struct {
	*mock.Call
}

// HandleIncomingShipMessage is a helper method to define mock.On call
//   - _a0 []byte
func (_e *WebsocketDataProcessing_Expecter) HandleIncomingShipMessage(_a0 interface{}) *WebsocketDataProcessing_HandleIncomingShipMessage_Call {
	return &WebsocketDataProcessing_HandleIncomingShipMessage_Call{Call: _e.mock.On("HandleIncomingShipMessage", _a0)}
}

func (_c *WebsocketDataProcessing_HandleIncomingShipMessage_Call) Run(run func(_a0 []byte)) *WebsocketDataProcessing_HandleIncomingShipMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *WebsocketDataProcessing_HandleIncomingShipMessage_Call) Return() *WebsocketDataProcessing_HandleIncomingShipMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *WebsocketDataProcessing_HandleIncomingShipMessage_Call) RunAndReturn(run func([]byte)) *WebsocketDataProcessing_HandleIncomingShipMessage_Call {
	_c.Call.Return(run)
	return _c
}

// ReportConnectionError provides a mock function with given fields: _a0
func (_m *WebsocketDataProcessing) ReportConnectionError(_a0 error) {
	_m.Called(_a0)
}

// WebsocketDataProcessing_ReportConnectionError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportConnectionError'
type WebsocketDataProcessing_ReportConnectionError_Call struct {
	*mock.Call
}

// ReportConnectionError is a helper method to define mock.On call
//   - _a0 error
func (_e *WebsocketDataProcessing_Expecter) ReportConnectionError(_a0 interface{}) *WebsocketDataProcessing_ReportConnectionError_Call {
	return &WebsocketDataProcessing_ReportConnectionError_Call{Call: _e.mock.On("ReportConnectionError", _a0)}
}

func (_c *WebsocketDataProcessing_ReportConnectionError_Call) Run(run func(_a0 error)) *WebsocketDataProcessing_ReportConnectionError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *WebsocketDataProcessing_ReportConnectionError_Call) Return() *WebsocketDataProcessing_ReportConnectionError_Call {
	_c.Call.Return()
	return _c
}

func (_c *WebsocketDataProcessing_ReportConnectionError_Call) RunAndReturn(run func(error)) *WebsocketDataProcessing_ReportConnectionError_Call {
	_c.Call.Return(run)
	return _c
}

// NewWebsocketDataProcessing creates a new instance of WebsocketDataProcessing. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebsocketDataProcessing(t interface {
	mock.TestingT
	Cleanup(func())
}) *WebsocketDataProcessing {
	mock := &WebsocketDataProcessing{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
