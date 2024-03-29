// Code generated by mockery v2.40.1. DO NOT EDIT.

package loader

import mock "github.com/stretchr/testify/mock"

// MockLoadedConfig is an autogenerated mock type for the LoadedConfig type
type MockLoadedConfig struct {
	mock.Mock
}

type MockLoadedConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLoadedConfig) EXPECT() *MockLoadedConfig_Expecter {
	return &MockLoadedConfig_Expecter{mock: &_m.Mock}
}

// Data provides a mock function with given fields:
func (_m *MockLoadedConfig) Data() map[string]interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// MockLoadedConfig_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type MockLoadedConfig_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *MockLoadedConfig_Expecter) Data() *MockLoadedConfig_Data_Call {
	return &MockLoadedConfig_Data_Call{Call: _e.mock.On("Data")}
}

func (_c *MockLoadedConfig_Data_Call) Run(run func()) *MockLoadedConfig_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLoadedConfig_Data_Call) Return(_a0 map[string]interface{}) *MockLoadedConfig_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLoadedConfig_Data_Call) RunAndReturn(run func() map[string]interface{}) *MockLoadedConfig_Data_Call {
	_c.Call.Return(run)
	return _c
}

// Path provides a mock function with given fields:
func (_m *MockLoadedConfig) Path() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Path")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockLoadedConfig_Path_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Path'
type MockLoadedConfig_Path_Call struct {
	*mock.Call
}

// Path is a helper method to define mock.On call
func (_e *MockLoadedConfig_Expecter) Path() *MockLoadedConfig_Path_Call {
	return &MockLoadedConfig_Path_Call{Call: _e.mock.On("Path")}
}

func (_c *MockLoadedConfig_Path_Call) Run(run func()) *MockLoadedConfig_Path_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLoadedConfig_Path_Call) Return(_a0 string) *MockLoadedConfig_Path_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLoadedConfig_Path_Call) RunAndReturn(run func() string) *MockLoadedConfig_Path_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLoadedConfig creates a new instance of MockLoadedConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLoadedConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLoadedConfig {
	mock := &MockLoadedConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
