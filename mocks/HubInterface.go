// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"
)

// HubInterface is an autogenerated mock type for the HubInterface type
type HubInterface struct {
	mock.Mock
}

type HubInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *HubInterface) EXPECT() *HubInterface_Expecter {
	return &HubInterface_Expecter{mock: &_m.Mock}
}

// CancelPairingWithSKI provides a mock function with given fields: ski
func (_m *HubInterface) CancelPairingWithSKI(ski string) {
	_m.Called(ski)
}

// HubInterface_CancelPairingWithSKI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelPairingWithSKI'
type HubInterface_CancelPairingWithSKI_Call struct {
	*mock.Call
}

// CancelPairingWithSKI is a helper method to define mock.On call
//   - ski string
func (_e *HubInterface_Expecter) CancelPairingWithSKI(ski interface{}) *HubInterface_CancelPairingWithSKI_Call {
	return &HubInterface_CancelPairingWithSKI_Call{Call: _e.mock.On("CancelPairingWithSKI", ski)}
}

func (_c *HubInterface_CancelPairingWithSKI_Call) Run(run func(ski string)) *HubInterface_CancelPairingWithSKI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubInterface_CancelPairingWithSKI_Call) Return() *HubInterface_CancelPairingWithSKI_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubInterface_CancelPairingWithSKI_Call) RunAndReturn(run func(string)) *HubInterface_CancelPairingWithSKI_Call {
	_c.Call.Return(run)
	return _c
}

// DisconnectSKI provides a mock function with given fields: ski, reason
func (_m *HubInterface) DisconnectSKI(ski string, reason string) {
	_m.Called(ski, reason)
}

// HubInterface_DisconnectSKI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisconnectSKI'
type HubInterface_DisconnectSKI_Call struct {
	*mock.Call
}

// DisconnectSKI is a helper method to define mock.On call
//   - ski string
//   - reason string
func (_e *HubInterface_Expecter) DisconnectSKI(ski interface{}, reason interface{}) *HubInterface_DisconnectSKI_Call {
	return &HubInterface_DisconnectSKI_Call{Call: _e.mock.On("DisconnectSKI", ski, reason)}
}

func (_c *HubInterface_DisconnectSKI_Call) Run(run func(ski string, reason string)) *HubInterface_DisconnectSKI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *HubInterface_DisconnectSKI_Call) Return() *HubInterface_DisconnectSKI_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubInterface_DisconnectSKI_Call) RunAndReturn(run func(string, string)) *HubInterface_DisconnectSKI_Call {
	_c.Call.Return(run)
	return _c
}

// InitiateOrApprovePairingWithSKI provides a mock function with given fields: ski
func (_m *HubInterface) InitiateOrApprovePairingWithSKI(ski string) {
	_m.Called(ski)
}

// HubInterface_InitiateOrApprovePairingWithSKI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitiateOrApprovePairingWithSKI'
type HubInterface_InitiateOrApprovePairingWithSKI_Call struct {
	*mock.Call
}

// InitiateOrApprovePairingWithSKI is a helper method to define mock.On call
//   - ski string
func (_e *HubInterface_Expecter) InitiateOrApprovePairingWithSKI(ski interface{}) *HubInterface_InitiateOrApprovePairingWithSKI_Call {
	return &HubInterface_InitiateOrApprovePairingWithSKI_Call{Call: _e.mock.On("InitiateOrApprovePairingWithSKI", ski)}
}

func (_c *HubInterface_InitiateOrApprovePairingWithSKI_Call) Run(run func(ski string)) *HubInterface_InitiateOrApprovePairingWithSKI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubInterface_InitiateOrApprovePairingWithSKI_Call) Return() *HubInterface_InitiateOrApprovePairingWithSKI_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubInterface_InitiateOrApprovePairingWithSKI_Call) RunAndReturn(run func(string)) *HubInterface_InitiateOrApprovePairingWithSKI_Call {
	_c.Call.Return(run)
	return _c
}

// PairingDetailForSki provides a mock function with given fields: ski
func (_m *HubInterface) PairingDetailForSki(ski string) *api.ConnectionStateDetail {
	ret := _m.Called(ski)

	if len(ret) == 0 {
		panic("no return value specified for PairingDetailForSki")
	}

	var r0 *api.ConnectionStateDetail
	if rf, ok := ret.Get(0).(func(string) *api.ConnectionStateDetail); ok {
		r0 = rf(ski)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ConnectionStateDetail)
		}
	}

	return r0
}

// HubInterface_PairingDetailForSki_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PairingDetailForSki'
type HubInterface_PairingDetailForSki_Call struct {
	*mock.Call
}

// PairingDetailForSki is a helper method to define mock.On call
//   - ski string
func (_e *HubInterface_Expecter) PairingDetailForSki(ski interface{}) *HubInterface_PairingDetailForSki_Call {
	return &HubInterface_PairingDetailForSki_Call{Call: _e.mock.On("PairingDetailForSki", ski)}
}

func (_c *HubInterface_PairingDetailForSki_Call) Run(run func(ski string)) *HubInterface_PairingDetailForSki_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubInterface_PairingDetailForSki_Call) Return(_a0 *api.ConnectionStateDetail) *HubInterface_PairingDetailForSki_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HubInterface_PairingDetailForSki_Call) RunAndReturn(run func(string) *api.ConnectionStateDetail) *HubInterface_PairingDetailForSki_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterRemoteSKI provides a mock function with given fields: ski
func (_m *HubInterface) RegisterRemoteSKI(ski string) {
	_m.Called(ski)
}

// HubInterface_RegisterRemoteSKI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterRemoteSKI'
type HubInterface_RegisterRemoteSKI_Call struct {
	*mock.Call
}

// RegisterRemoteSKI is a helper method to define mock.On call
//   - ski string
func (_e *HubInterface_Expecter) RegisterRemoteSKI(ski interface{}) *HubInterface_RegisterRemoteSKI_Call {
	return &HubInterface_RegisterRemoteSKI_Call{Call: _e.mock.On("RegisterRemoteSKI", ski)}
}

func (_c *HubInterface_RegisterRemoteSKI_Call) Run(run func(ski string)) *HubInterface_RegisterRemoteSKI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubInterface_RegisterRemoteSKI_Call) Return() *HubInterface_RegisterRemoteSKI_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubInterface_RegisterRemoteSKI_Call) RunAndReturn(run func(string)) *HubInterface_RegisterRemoteSKI_Call {
	_c.Call.Return(run)
	return _c
}

// ServiceForSKI provides a mock function with given fields: ski
func (_m *HubInterface) ServiceForSKI(ski string) *api.ServiceDetails {
	ret := _m.Called(ski)

	if len(ret) == 0 {
		panic("no return value specified for ServiceForSKI")
	}

	var r0 *api.ServiceDetails
	if rf, ok := ret.Get(0).(func(string) *api.ServiceDetails); ok {
		r0 = rf(ski)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ServiceDetails)
		}
	}

	return r0
}

// HubInterface_ServiceForSKI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServiceForSKI'
type HubInterface_ServiceForSKI_Call struct {
	*mock.Call
}

// ServiceForSKI is a helper method to define mock.On call
//   - ski string
func (_e *HubInterface_Expecter) ServiceForSKI(ski interface{}) *HubInterface_ServiceForSKI_Call {
	return &HubInterface_ServiceForSKI_Call{Call: _e.mock.On("ServiceForSKI", ski)}
}

func (_c *HubInterface_ServiceForSKI_Call) Run(run func(ski string)) *HubInterface_ServiceForSKI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubInterface_ServiceForSKI_Call) Return(_a0 *api.ServiceDetails) *HubInterface_ServiceForSKI_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HubInterface_ServiceForSKI_Call) RunAndReturn(run func(string) *api.ServiceDetails) *HubInterface_ServiceForSKI_Call {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with given fields:
func (_m *HubInterface) Shutdown() {
	_m.Called()
}

// HubInterface_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type HubInterface_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
func (_e *HubInterface_Expecter) Shutdown() *HubInterface_Shutdown_Call {
	return &HubInterface_Shutdown_Call{Call: _e.mock.On("Shutdown")}
}

func (_c *HubInterface_Shutdown_Call) Run(run func()) *HubInterface_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HubInterface_Shutdown_Call) Return() *HubInterface_Shutdown_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubInterface_Shutdown_Call) RunAndReturn(run func()) *HubInterface_Shutdown_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *HubInterface) Start() {
	_m.Called()
}

// HubInterface_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type HubInterface_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *HubInterface_Expecter) Start() *HubInterface_Start_Call {
	return &HubInterface_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *HubInterface_Start_Call) Run(run func()) *HubInterface_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HubInterface_Start_Call) Return() *HubInterface_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubInterface_Start_Call) RunAndReturn(run func()) *HubInterface_Start_Call {
	_c.Call.Return(run)
	return _c
}

// UnregisterRemoteSKI provides a mock function with given fields: ski
func (_m *HubInterface) UnregisterRemoteSKI(ski string) {
	_m.Called(ski)
}

// HubInterface_UnregisterRemoteSKI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnregisterRemoteSKI'
type HubInterface_UnregisterRemoteSKI_Call struct {
	*mock.Call
}

// UnregisterRemoteSKI is a helper method to define mock.On call
//   - ski string
func (_e *HubInterface_Expecter) UnregisterRemoteSKI(ski interface{}) *HubInterface_UnregisterRemoteSKI_Call {
	return &HubInterface_UnregisterRemoteSKI_Call{Call: _e.mock.On("UnregisterRemoteSKI", ski)}
}

func (_c *HubInterface_UnregisterRemoteSKI_Call) Run(run func(ski string)) *HubInterface_UnregisterRemoteSKI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubInterface_UnregisterRemoteSKI_Call) Return() *HubInterface_UnregisterRemoteSKI_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubInterface_UnregisterRemoteSKI_Call) RunAndReturn(run func(string)) *HubInterface_UnregisterRemoteSKI_Call {
	_c.Call.Return(run)
	return _c
}

// NewHubInterface creates a new instance of HubInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHubInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *HubInterface {
	mock := &HubInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
