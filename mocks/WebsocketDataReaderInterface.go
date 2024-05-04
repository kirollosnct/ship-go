// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// WebsocketDataReaderInterface is an autogenerated mock type for the WebsocketDataReaderInterface type
type WebsocketDataReaderInterface struct {
	mock.Mock
}

type WebsocketDataReaderInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *WebsocketDataReaderInterface) EXPECT() *WebsocketDataReaderInterface_Expecter {
	return &WebsocketDataReaderInterface_Expecter{mock: &_m.Mock}
}

// HandleIncomingWebsocketMessage provides a mock function with given fields: _a0
func (_m *WebsocketDataReaderInterface) HandleIncomingWebsocketMessage(_a0 []byte) {
	_m.Called(_a0)
}

// WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleIncomingWebsocketMessage'
type WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call struct {
	*mock.Call
}

// HandleIncomingWebsocketMessage is a helper method to define mock.On call
//   - _a0 []byte
func (_e *WebsocketDataReaderInterface_Expecter) HandleIncomingWebsocketMessage(_a0 interface{}) *WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call {
	return &WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call{Call: _e.mock.On("HandleIncomingWebsocketMessage", _a0)}
}

func (_c *WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call) Run(run func(_a0 []byte)) *WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call) Return() *WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call) RunAndReturn(run func([]byte)) *WebsocketDataReaderInterface_HandleIncomingWebsocketMessage_Call {
	_c.Call.Return(run)
	return _c
}

// ReportConnectionError provides a mock function with given fields: _a0
func (_m *WebsocketDataReaderInterface) ReportConnectionError(_a0 error) {
	_m.Called(_a0)
}

// WebsocketDataReaderInterface_ReportConnectionError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportConnectionError'
type WebsocketDataReaderInterface_ReportConnectionError_Call struct {
	*mock.Call
}

// ReportConnectionError is a helper method to define mock.On call
//   - _a0 error
func (_e *WebsocketDataReaderInterface_Expecter) ReportConnectionError(_a0 interface{}) *WebsocketDataReaderInterface_ReportConnectionError_Call {
	return &WebsocketDataReaderInterface_ReportConnectionError_Call{Call: _e.mock.On("ReportConnectionError", _a0)}
}

func (_c *WebsocketDataReaderInterface_ReportConnectionError_Call) Run(run func(_a0 error)) *WebsocketDataReaderInterface_ReportConnectionError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *WebsocketDataReaderInterface_ReportConnectionError_Call) Return() *WebsocketDataReaderInterface_ReportConnectionError_Call {
	_c.Call.Return()
	return _c
}

func (_c *WebsocketDataReaderInterface_ReportConnectionError_Call) RunAndReturn(run func(error)) *WebsocketDataReaderInterface_ReportConnectionError_Call {
	_c.Call.Return(run)
	return _c
}

// NewWebsocketDataReaderInterface creates a new instance of WebsocketDataReaderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebsocketDataReaderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *WebsocketDataReaderInterface {
	mock := &WebsocketDataReaderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
