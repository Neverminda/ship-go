// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/ship-go/model"
)

// ShipConnection is an autogenerated mock type for the ShipConnection type
type ShipConnection struct {
	mock.Mock
}

type ShipConnection_Expecter struct {
	mock *mock.Mock
}

func (_m *ShipConnection) EXPECT() *ShipConnection_Expecter {
	return &ShipConnection_Expecter{mock: &_m.Mock}
}

// AbortPendingHandshake provides a mock function with given fields:
func (_m *ShipConnection) AbortPendingHandshake() {
	_m.Called()
}

// ShipConnection_AbortPendingHandshake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbortPendingHandshake'
type ShipConnection_AbortPendingHandshake_Call struct {
	*mock.Call
}

// AbortPendingHandshake is a helper method to define mock.On call
func (_e *ShipConnection_Expecter) AbortPendingHandshake() *ShipConnection_AbortPendingHandshake_Call {
	return &ShipConnection_AbortPendingHandshake_Call{Call: _e.mock.On("AbortPendingHandshake")}
}

func (_c *ShipConnection_AbortPendingHandshake_Call) Run(run func()) *ShipConnection_AbortPendingHandshake_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ShipConnection_AbortPendingHandshake_Call) Return() *ShipConnection_AbortPendingHandshake_Call {
	_c.Call.Return()
	return _c
}

func (_c *ShipConnection_AbortPendingHandshake_Call) RunAndReturn(run func()) *ShipConnection_AbortPendingHandshake_Call {
	_c.Call.Return(run)
	return _c
}

// ApprovePendingHandshake provides a mock function with given fields:
func (_m *ShipConnection) ApprovePendingHandshake() {
	_m.Called()
}

// ShipConnection_ApprovePendingHandshake_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApprovePendingHandshake'
type ShipConnection_ApprovePendingHandshake_Call struct {
	*mock.Call
}

// ApprovePendingHandshake is a helper method to define mock.On call
func (_e *ShipConnection_Expecter) ApprovePendingHandshake() *ShipConnection_ApprovePendingHandshake_Call {
	return &ShipConnection_ApprovePendingHandshake_Call{Call: _e.mock.On("ApprovePendingHandshake")}
}

func (_c *ShipConnection_ApprovePendingHandshake_Call) Run(run func()) *ShipConnection_ApprovePendingHandshake_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ShipConnection_ApprovePendingHandshake_Call) Return() *ShipConnection_ApprovePendingHandshake_Call {
	_c.Call.Return()
	return _c
}

func (_c *ShipConnection_ApprovePendingHandshake_Call) RunAndReturn(run func()) *ShipConnection_ApprovePendingHandshake_Call {
	_c.Call.Return(run)
	return _c
}

// CloseConnection provides a mock function with given fields: safe, code, reason
func (_m *ShipConnection) CloseConnection(safe bool, code int, reason string) {
	_m.Called(safe, code, reason)
}

// ShipConnection_CloseConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseConnection'
type ShipConnection_CloseConnection_Call struct {
	*mock.Call
}

// CloseConnection is a helper method to define mock.On call
//   - safe bool
//   - code int
//   - reason string
func (_e *ShipConnection_Expecter) CloseConnection(safe interface{}, code interface{}, reason interface{}) *ShipConnection_CloseConnection_Call {
	return &ShipConnection_CloseConnection_Call{Call: _e.mock.On("CloseConnection", safe, code, reason)}
}

func (_c *ShipConnection_CloseConnection_Call) Run(run func(safe bool, code int, reason string)) *ShipConnection_CloseConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(int), args[2].(string))
	})
	return _c
}

func (_c *ShipConnection_CloseConnection_Call) Return() *ShipConnection_CloseConnection_Call {
	_c.Call.Return()
	return _c
}

func (_c *ShipConnection_CloseConnection_Call) RunAndReturn(run func(bool, int, string)) *ShipConnection_CloseConnection_Call {
	_c.Call.Return(run)
	return _c
}

// DataHandler provides a mock function with given fields:
func (_m *ShipConnection) DataHandler() api.WebsocketDataConnection {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DataHandler")
	}

	var r0 api.WebsocketDataConnection
	if rf, ok := ret.Get(0).(func() api.WebsocketDataConnection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.WebsocketDataConnection)
		}
	}

	return r0
}

// ShipConnection_DataHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DataHandler'
type ShipConnection_DataHandler_Call struct {
	*mock.Call
}

// DataHandler is a helper method to define mock.On call
func (_e *ShipConnection_Expecter) DataHandler() *ShipConnection_DataHandler_Call {
	return &ShipConnection_DataHandler_Call{Call: _e.mock.On("DataHandler")}
}

func (_c *ShipConnection_DataHandler_Call) Run(run func()) *ShipConnection_DataHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ShipConnection_DataHandler_Call) Return(_a0 api.WebsocketDataConnection) *ShipConnection_DataHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ShipConnection_DataHandler_Call) RunAndReturn(run func() api.WebsocketDataConnection) *ShipConnection_DataHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteSKI provides a mock function with given fields:
func (_m *ShipConnection) RemoteSKI() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteSKI")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ShipConnection_RemoteSKI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteSKI'
type ShipConnection_RemoteSKI_Call struct {
	*mock.Call
}

// RemoteSKI is a helper method to define mock.On call
func (_e *ShipConnection_Expecter) RemoteSKI() *ShipConnection_RemoteSKI_Call {
	return &ShipConnection_RemoteSKI_Call{Call: _e.mock.On("RemoteSKI")}
}

func (_c *ShipConnection_RemoteSKI_Call) Run(run func()) *ShipConnection_RemoteSKI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ShipConnection_RemoteSKI_Call) Return(_a0 string) *ShipConnection_RemoteSKI_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ShipConnection_RemoteSKI_Call) RunAndReturn(run func() string) *ShipConnection_RemoteSKI_Call {
	_c.Call.Return(run)
	return _c
}

// ShipHandshakeState provides a mock function with given fields:
func (_m *ShipConnection) ShipHandshakeState() (model.ShipMessageExchangeState, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ShipHandshakeState")
	}

	var r0 model.ShipMessageExchangeState
	var r1 error
	if rf, ok := ret.Get(0).(func() (model.ShipMessageExchangeState, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() model.ShipMessageExchangeState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.ShipMessageExchangeState)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShipConnection_ShipHandshakeState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShipHandshakeState'
type ShipConnection_ShipHandshakeState_Call struct {
	*mock.Call
}

// ShipHandshakeState is a helper method to define mock.On call
func (_e *ShipConnection_Expecter) ShipHandshakeState() *ShipConnection_ShipHandshakeState_Call {
	return &ShipConnection_ShipHandshakeState_Call{Call: _e.mock.On("ShipHandshakeState")}
}

func (_c *ShipConnection_ShipHandshakeState_Call) Run(run func()) *ShipConnection_ShipHandshakeState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ShipConnection_ShipHandshakeState_Call) Return(_a0 model.ShipMessageExchangeState, _a1 error) *ShipConnection_ShipHandshakeState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ShipConnection_ShipHandshakeState_Call) RunAndReturn(run func() (model.ShipMessageExchangeState, error)) *ShipConnection_ShipHandshakeState_Call {
	_c.Call.Return(run)
	return _c
}

// NewShipConnection creates a new instance of ShipConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShipConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShipConnection {
	mock := &ShipConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
