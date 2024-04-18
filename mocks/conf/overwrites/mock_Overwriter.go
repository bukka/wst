// Code generated by mockery v2.40.1. DO NOT EDIT.

package overwrites

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/bukka/wst/conf/types"
)

// MockOverwriter is an autogenerated mock type for the Overwriter type
type MockOverwriter struct {
	mock.Mock
}

type MockOverwriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOverwriter) EXPECT() *MockOverwriter_Expecter {
	return &MockOverwriter_Expecter{mock: &_m.Mock}
}

// Overwrite provides a mock function with given fields: config, _a1
func (_m *MockOverwriter) Overwrite(config *types.Config, _a1 map[string]string) error {
	ret := _m.Called(config, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Overwrite")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Config, map[string]string) error); ok {
		r0 = rf(config, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOverwriter_Overwrite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Overwrite'
type MockOverwriter_Overwrite_Call struct {
	*mock.Call
}

// Overwrite is a helper method to define mock.On call
//   - config *types.Config
//   - _a1 map[string]string
func (_e *MockOverwriter_Expecter) Overwrite(config interface{}, _a1 interface{}) *MockOverwriter_Overwrite_Call {
	return &MockOverwriter_Overwrite_Call{Call: _e.mock.On("Overwrite", config, _a1)}
}

func (_c *MockOverwriter_Overwrite_Call) Run(run func(config *types.Config, _a1 map[string]string)) *MockOverwriter_Overwrite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Config), args[1].(map[string]string))
	})
	return _c
}

func (_c *MockOverwriter_Overwrite_Call) Return(_a0 error) *MockOverwriter_Overwrite_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOverwriter_Overwrite_Call) RunAndReturn(run func(*types.Config, map[string]string) error) *MockOverwriter_Overwrite_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOverwriter creates a new instance of MockOverwriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOverwriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOverwriter {
	mock := &MockOverwriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}