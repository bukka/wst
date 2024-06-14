// Code generated by mockery v2.40.1. DO NOT EDIT.

package clients

import (
	context "context"

	clients "github.com/bukka/wst/run/environments/environment/providers/kubernetes/clients"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/apps/v1"
)

// MockDeploymentClient is an autogenerated mock type for the DeploymentClient type
type MockDeploymentClient struct {
	mock.Mock
}

type MockDeploymentClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeploymentClient) EXPECT() *MockDeploymentClient_Expecter {
	return &MockDeploymentClient_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, deployment, opts
func (_m *MockDeploymentClient) Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (*v1.Deployment, error) {
	ret := _m.Called(ctx, deployment, opts)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *v1.Deployment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Deployment, metav1.CreateOptions) (*v1.Deployment, error)); ok {
		return rf(ctx, deployment, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Deployment, metav1.CreateOptions) *v1.Deployment); ok {
		r0 = rf(ctx, deployment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.Deployment, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, deployment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDeploymentClient_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockDeploymentClient_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - deployment *v1.Deployment
//   - opts metav1.CreateOptions
func (_e *MockDeploymentClient_Expecter) Create(ctx interface{}, deployment interface{}, opts interface{}) *MockDeploymentClient_Create_Call {
	return &MockDeploymentClient_Create_Call{Call: _e.mock.On("Create", ctx, deployment, opts)}
}

func (_c *MockDeploymentClient_Create_Call) Run(run func(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions)) *MockDeploymentClient_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.Deployment), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *MockDeploymentClient_Create_Call) Return(_a0 *v1.Deployment, _a1 error) *MockDeploymentClient_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDeploymentClient_Create_Call) RunAndReturn(run func(context.Context, *v1.Deployment, metav1.CreateOptions) (*v1.Deployment, error)) *MockDeploymentClient_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *MockDeploymentClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDeploymentClient_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDeploymentClient_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *MockDeploymentClient_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *MockDeploymentClient_Delete_Call {
	return &MockDeploymentClient_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *MockDeploymentClient_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *MockDeploymentClient_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *MockDeploymentClient_Delete_Call) Return(_a0 error) *MockDeploymentClient_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeploymentClient_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *MockDeploymentClient_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *MockDeploymentClient) Watch(ctx context.Context, opts metav1.ListOptions) (clients.WatchResult, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for Watch")
	}

	var r0 clients.WatchResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (clients.WatchResult, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) clients.WatchResult); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clients.WatchResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDeploymentClient_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockDeploymentClient_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *MockDeploymentClient_Expecter) Watch(ctx interface{}, opts interface{}) *MockDeploymentClient_Watch_Call {
	return &MockDeploymentClient_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *MockDeploymentClient_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *MockDeploymentClient_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *MockDeploymentClient_Watch_Call) Return(_a0 clients.WatchResult, _a1 error) *MockDeploymentClient_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDeploymentClient_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (clients.WatchResult, error)) *MockDeploymentClient_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDeploymentClient creates a new instance of MockDeploymentClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeploymentClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeploymentClient {
	mock := &MockDeploymentClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
