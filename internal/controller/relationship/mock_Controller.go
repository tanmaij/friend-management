// Code generated by mockery v2.46.3. DO NOT EDIT.

package relationship

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

type MockController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockController) EXPECT() *MockController_Expecter {
	return &MockController_Expecter{mock: &_m.Mock}
}

// CreateFriendConn provides a mock function with given fields: ctx, inp
func (_m *MockController) CreateFriendConn(ctx context.Context, inp CreateFriendConnInp) error {
	ret := _m.Called(ctx, inp)

	if len(ret) == 0 {
		panic("no return value specified for CreateFriendConn")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CreateFriendConnInp) error); ok {
		r0 = rf(ctx, inp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockController_CreateFriendConn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateFriendConn'
type MockController_CreateFriendConn_Call struct {
	*mock.Call
}

// CreateFriendConn is a helper method to define mock.On call
//   - ctx context.Context
//   - inp CreateFriendConnInp
func (_e *MockController_Expecter) CreateFriendConn(ctx interface{}, inp interface{}) *MockController_CreateFriendConn_Call {
	return &MockController_CreateFriendConn_Call{Call: _e.mock.On("CreateFriendConn", ctx, inp)}
}

func (_c *MockController_CreateFriendConn_Call) Run(run func(ctx context.Context, inp CreateFriendConnInp)) *MockController_CreateFriendConn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(CreateFriendConnInp))
	})
	return _c
}

func (_c *MockController_CreateFriendConn_Call) Return(_a0 error) *MockController_CreateFriendConn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockController_CreateFriendConn_Call) RunAndReturn(run func(context.Context, CreateFriendConnInp) error) *MockController_CreateFriendConn_Call {
	_c.Call.Return(run)
	return _c
}

// ListFriendByEmail provides a mock function with given fields: ctx, inp
func (_m *MockController) ListFriendByEmail(ctx context.Context, inp ListFriendByEmailInput) (ListFriendByEmailOutput, error) {
	ret := _m.Called(ctx, inp)

	if len(ret) == 0 {
		panic("no return value specified for ListFriendByEmail")
	}

	var r0 ListFriendByEmailOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ListFriendByEmailInput) (ListFriendByEmailOutput, error)); ok {
		return rf(ctx, inp)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ListFriendByEmailInput) ListFriendByEmailOutput); ok {
		r0 = rf(ctx, inp)
	} else {
		r0 = ret.Get(0).(ListFriendByEmailOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ListFriendByEmailInput) error); ok {
		r1 = rf(ctx, inp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockController_ListFriendByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFriendByEmail'
type MockController_ListFriendByEmail_Call struct {
	*mock.Call
}

// ListFriendByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - inp ListFriendByEmailInput
func (_e *MockController_Expecter) ListFriendByEmail(ctx interface{}, inp interface{}) *MockController_ListFriendByEmail_Call {
	return &MockController_ListFriendByEmail_Call{Call: _e.mock.On("ListFriendByEmail", ctx, inp)}
}

func (_c *MockController_ListFriendByEmail_Call) Run(run func(ctx context.Context, inp ListFriendByEmailInput)) *MockController_ListFriendByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ListFriendByEmailInput))
	})
	return _c
}

func (_c *MockController_ListFriendByEmail_Call) Return(_a0 ListFriendByEmailOutput, _a1 error) *MockController_ListFriendByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockController_ListFriendByEmail_Call) RunAndReturn(run func(context.Context, ListFriendByEmailInput) (ListFriendByEmailOutput, error)) *MockController_ListFriendByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// ListTwoEmailCommonFriends provides a mock function with given fields: ctx, inp
func (_m *MockController) ListTwoEmailCommonFriends(ctx context.Context, inp ListTwoEmailCommonFriendsInput) (ListTwoEmailCommonFriendsOutput, error) {
	ret := _m.Called(ctx, inp)

	if len(ret) == 0 {
		panic("no return value specified for ListTwoEmailCommonFriends")
	}

	var r0 ListTwoEmailCommonFriendsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ListTwoEmailCommonFriendsInput) (ListTwoEmailCommonFriendsOutput, error)); ok {
		return rf(ctx, inp)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ListTwoEmailCommonFriendsInput) ListTwoEmailCommonFriendsOutput); ok {
		r0 = rf(ctx, inp)
	} else {
		r0 = ret.Get(0).(ListTwoEmailCommonFriendsOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ListTwoEmailCommonFriendsInput) error); ok {
		r1 = rf(ctx, inp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockController_ListTwoEmailCommonFriends_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTwoEmailCommonFriends'
type MockController_ListTwoEmailCommonFriends_Call struct {
	*mock.Call
}

// ListTwoEmailCommonFriends is a helper method to define mock.On call
//   - ctx context.Context
//   - inp ListTwoEmailCommonFriendsInput
func (_e *MockController_Expecter) ListTwoEmailCommonFriends(ctx interface{}, inp interface{}) *MockController_ListTwoEmailCommonFriends_Call {
	return &MockController_ListTwoEmailCommonFriends_Call{Call: _e.mock.On("ListTwoEmailCommonFriends", ctx, inp)}
}

func (_c *MockController_ListTwoEmailCommonFriends_Call) Run(run func(ctx context.Context, inp ListTwoEmailCommonFriendsInput)) *MockController_ListTwoEmailCommonFriends_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ListTwoEmailCommonFriendsInput))
	})
	return _c
}

func (_c *MockController_ListTwoEmailCommonFriends_Call) Return(_a0 ListTwoEmailCommonFriendsOutput, _a1 error) *MockController_ListTwoEmailCommonFriends_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockController_ListTwoEmailCommonFriends_Call) RunAndReturn(run func(context.Context, ListTwoEmailCommonFriendsInput) (ListTwoEmailCommonFriendsOutput, error)) *MockController_ListTwoEmailCommonFriends_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockController creates a new instance of MockController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockController {
	mock := &MockController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
