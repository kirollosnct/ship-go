// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"
)

// MdnsProviderInterface is an autogenerated mock type for the MdnsProviderInterface type
type MdnsProviderInterface struct {
	mock.Mock
}

type MdnsProviderInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MdnsProviderInterface) EXPECT() *MdnsProviderInterface_Expecter {
	return &MdnsProviderInterface_Expecter{mock: &_m.Mock}
}

// Announce provides a mock function with given fields: serviceName, port, txt
func (_m *MdnsProviderInterface) Announce(serviceName string, port int, txt []string) error {
	ret := _m.Called(serviceName, port, txt)

	if len(ret) == 0 {
		panic("no return value specified for Announce")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, []string) error); ok {
		r0 = rf(serviceName, port, txt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MdnsProviderInterface_Announce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Announce'
type MdnsProviderInterface_Announce_Call struct {
	*mock.Call
}

// Announce is a helper method to define mock.On call
//   - serviceName string
//   - port int
//   - txt []string
func (_e *MdnsProviderInterface_Expecter) Announce(serviceName interface{}, port interface{}, txt interface{}) *MdnsProviderInterface_Announce_Call {
	return &MdnsProviderInterface_Announce_Call{Call: _e.mock.On("Announce", serviceName, port, txt)}
}

func (_c *MdnsProviderInterface_Announce_Call) Run(run func(serviceName string, port int, txt []string)) *MdnsProviderInterface_Announce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int), args[2].([]string))
	})
	return _c
}

func (_c *MdnsProviderInterface_Announce_Call) Return(_a0 error) *MdnsProviderInterface_Announce_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MdnsProviderInterface_Announce_Call) RunAndReturn(run func(string, int, []string) error) *MdnsProviderInterface_Announce_Call {
	_c.Call.Return(run)
	return _c
}

// CheckAvailability provides a mock function with given fields:
func (_m *MdnsProviderInterface) CheckAvailability() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CheckAvailability")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MdnsProviderInterface_CheckAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckAvailability'
type MdnsProviderInterface_CheckAvailability_Call struct {
	*mock.Call
}

// CheckAvailability is a helper method to define mock.On call
func (_e *MdnsProviderInterface_Expecter) CheckAvailability() *MdnsProviderInterface_CheckAvailability_Call {
	return &MdnsProviderInterface_CheckAvailability_Call{Call: _e.mock.On("CheckAvailability")}
}

func (_c *MdnsProviderInterface_CheckAvailability_Call) Run(run func()) *MdnsProviderInterface_CheckAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MdnsProviderInterface_CheckAvailability_Call) Return(_a0 bool) *MdnsProviderInterface_CheckAvailability_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MdnsProviderInterface_CheckAvailability_Call) RunAndReturn(run func() bool) *MdnsProviderInterface_CheckAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// ResolveEntries provides a mock function with given fields: cb
func (_m *MdnsProviderInterface) ResolveEntries(cb api.MdnsResolveCB) {
	_m.Called(cb)
}

// MdnsProviderInterface_ResolveEntries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResolveEntries'
type MdnsProviderInterface_ResolveEntries_Call struct {
	*mock.Call
}

// ResolveEntries is a helper method to define mock.On call
//   - cb api.MdnsResolveCB
func (_e *MdnsProviderInterface_Expecter) ResolveEntries(cb interface{}) *MdnsProviderInterface_ResolveEntries_Call {
	return &MdnsProviderInterface_ResolveEntries_Call{Call: _e.mock.On("ResolveEntries", cb)}
}

func (_c *MdnsProviderInterface_ResolveEntries_Call) Run(run func(cb api.MdnsResolveCB)) *MdnsProviderInterface_ResolveEntries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.MdnsResolveCB))
	})
	return _c
}

func (_c *MdnsProviderInterface_ResolveEntries_Call) Return() *MdnsProviderInterface_ResolveEntries_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsProviderInterface_ResolveEntries_Call) RunAndReturn(run func(api.MdnsResolveCB)) *MdnsProviderInterface_ResolveEntries_Call {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with given fields:
func (_m *MdnsProviderInterface) Shutdown() {
	_m.Called()
}

// MdnsProviderInterface_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type MdnsProviderInterface_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
func (_e *MdnsProviderInterface_Expecter) Shutdown() *MdnsProviderInterface_Shutdown_Call {
	return &MdnsProviderInterface_Shutdown_Call{Call: _e.mock.On("Shutdown")}
}

func (_c *MdnsProviderInterface_Shutdown_Call) Run(run func()) *MdnsProviderInterface_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MdnsProviderInterface_Shutdown_Call) Return() *MdnsProviderInterface_Shutdown_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsProviderInterface_Shutdown_Call) RunAndReturn(run func()) *MdnsProviderInterface_Shutdown_Call {
	_c.Call.Return(run)
	return _c
}

// Unannounce provides a mock function with given fields:
func (_m *MdnsProviderInterface) Unannounce() {
	_m.Called()
}

// MdnsProviderInterface_Unannounce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unannounce'
type MdnsProviderInterface_Unannounce_Call struct {
	*mock.Call
}

// Unannounce is a helper method to define mock.On call
func (_e *MdnsProviderInterface_Expecter) Unannounce() *MdnsProviderInterface_Unannounce_Call {
	return &MdnsProviderInterface_Unannounce_Call{Call: _e.mock.On("Unannounce")}
}

func (_c *MdnsProviderInterface_Unannounce_Call) Run(run func()) *MdnsProviderInterface_Unannounce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MdnsProviderInterface_Unannounce_Call) Return() *MdnsProviderInterface_Unannounce_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsProviderInterface_Unannounce_Call) RunAndReturn(run func()) *MdnsProviderInterface_Unannounce_Call {
	_c.Call.Return(run)
	return _c
}

// NewMdnsProviderInterface creates a new instance of MdnsProviderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMdnsProviderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MdnsProviderInterface {
	mock := &MdnsProviderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
