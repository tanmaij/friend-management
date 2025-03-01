// Code generated by mockery v2.46.3. DO NOT EDIT.

package user

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/tanmaij/friend-management/internal/model"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, user
func (_m *MockRepository) Create(ctx context.Context, user model.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - user model.User
func (_e *MockRepository_Expecter) Create(ctx interface{}, user interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", ctx, user)}
}

func (_c *MockRepository_Create_Call) Run(run func(ctx context.Context, user model.User)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.User))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, model.User) error) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// ExistsByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for ExistsByEmail")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_ExistsByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExistsByEmail'
type MockRepository_ExistsByEmail_Call struct {
	*mock.Call
}

// ExistsByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockRepository_Expecter) ExistsByEmail(ctx interface{}, email interface{}) *MockRepository_ExistsByEmail_Call {
	return &MockRepository_ExistsByEmail_Call{Call: _e.mock.On("ExistsByEmail", ctx, email)}
}

func (_c *MockRepository_ExistsByEmail_Call) Run(run func(ctx context.Context, email string)) *MockRepository_ExistsByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_ExistsByEmail_Call) Return(_a0 bool, _a1 error) *MockRepository_ExistsByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_ExistsByEmail_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *MockRepository_ExistsByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type MockRepository_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockRepository_Expecter) GetByEmail(ctx interface{}, email interface{}) *MockRepository_GetByEmail_Call {
	return &MockRepository_GetByEmail_Call{Call: _e.mock.On("GetByEmail", ctx, email)}
}

func (_c *MockRepository_GetByEmail_Call) Run(run func(ctx context.Context, email string)) *MockRepository_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetByEmail_Call) Return(_a0 model.User, _a1 error) *MockRepository_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetByEmail_Call) RunAndReturn(run func(context.Context, string) (model.User, error)) *MockRepository_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
