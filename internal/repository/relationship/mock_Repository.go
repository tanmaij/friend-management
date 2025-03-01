// Code generated by mockery v2.46.3. DO NOT EDIT.

package relationship

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

// Create provides a mock function with given fields: ctx, relationship
func (_m *MockRepository) Create(ctx context.Context, relationship model.Relationship) error {
	ret := _m.Called(ctx, relationship)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Relationship) error); ok {
		r0 = rf(ctx, relationship)
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
//   - relationship model.Relationship
func (_e *MockRepository_Expecter) Create(ctx interface{}, relationship interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", ctx, relationship)}
}

func (_c *MockRepository_Create_Call) Run(run func(ctx context.Context, relationship model.Relationship)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Relationship))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, model.Relationship) error) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindEligibleRecipientEmailsWithMentioned provides a mock function with given fields: ctx, sender, mentionedEmails
func (_m *MockRepository) FindEligibleRecipientEmailsWithMentioned(ctx context.Context, sender string, mentionedEmails []string) ([]string, error) {
	ret := _m.Called(ctx, sender, mentionedEmails)

	if len(ret) == 0 {
		panic("no return value specified for FindEligibleRecipientEmailsWithMentioned")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]string, error)); ok {
		return rf(ctx, sender, mentionedEmails)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []string); ok {
		r0 = rf(ctx, sender, mentionedEmails)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, sender, mentionedEmails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindEligibleRecipientEmailsWithMentioned_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindEligibleRecipientEmailsWithMentioned'
type MockRepository_FindEligibleRecipientEmailsWithMentioned_Call struct {
	*mock.Call
}

// FindEligibleRecipientEmailsWithMentioned is a helper method to define mock.On call
//   - ctx context.Context
//   - sender string
//   - mentionedEmails []string
func (_e *MockRepository_Expecter) FindEligibleRecipientEmailsWithMentioned(ctx interface{}, sender interface{}, mentionedEmails interface{}) *MockRepository_FindEligibleRecipientEmailsWithMentioned_Call {
	return &MockRepository_FindEligibleRecipientEmailsWithMentioned_Call{Call: _e.mock.On("FindEligibleRecipientEmailsWithMentioned", ctx, sender, mentionedEmails)}
}

func (_c *MockRepository_FindEligibleRecipientEmailsWithMentioned_Call) Run(run func(ctx context.Context, sender string, mentionedEmails []string)) *MockRepository_FindEligibleRecipientEmailsWithMentioned_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockRepository_FindEligibleRecipientEmailsWithMentioned_Call) Return(_a0 []string, _a1 error) *MockRepository_FindEligibleRecipientEmailsWithMentioned_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindEligibleRecipientEmailsWithMentioned_Call) RunAndReturn(run func(context.Context, string, []string) ([]string, error)) *MockRepository_FindEligibleRecipientEmailsWithMentioned_Call {
	_c.Call.Return(run)
	return _c
}

// ListByTwoUserIDs provides a mock function with given fields: ctx, primaryUserID, secondaryUserID
func (_m *MockRepository) ListByTwoUserIDs(ctx context.Context, primaryUserID int, secondaryUserID int) ([]model.Relationship, error) {
	ret := _m.Called(ctx, primaryUserID, secondaryUserID)

	if len(ret) == 0 {
		panic("no return value specified for ListByTwoUserIDs")
	}

	var r0 []model.Relationship
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]model.Relationship, error)); ok {
		return rf(ctx, primaryUserID, secondaryUserID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []model.Relationship); ok {
		r0 = rf(ctx, primaryUserID, secondaryUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Relationship)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, primaryUserID, secondaryUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_ListByTwoUserIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByTwoUserIDs'
type MockRepository_ListByTwoUserIDs_Call struct {
	*mock.Call
}

// ListByTwoUserIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - primaryUserID int
//   - secondaryUserID int
func (_e *MockRepository_Expecter) ListByTwoUserIDs(ctx interface{}, primaryUserID interface{}, secondaryUserID interface{}) *MockRepository_ListByTwoUserIDs_Call {
	return &MockRepository_ListByTwoUserIDs_Call{Call: _e.mock.On("ListByTwoUserIDs", ctx, primaryUserID, secondaryUserID)}
}

func (_c *MockRepository_ListByTwoUserIDs_Call) Run(run func(ctx context.Context, primaryUserID int, secondaryUserID int)) *MockRepository_ListByTwoUserIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockRepository_ListByTwoUserIDs_Call) Return(_a0 []model.Relationship, _a1 error) *MockRepository_ListByTwoUserIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_ListByTwoUserIDs_Call) RunAndReturn(run func(context.Context, int, int) ([]model.Relationship, error)) *MockRepository_ListByTwoUserIDs_Call {
	_c.Call.Return(run)
	return _c
}

// ListFriendByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) ListFriendByEmail(ctx context.Context, email string) ([]model.User, int64, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for ListFriendByEmail")
	}

	var r0 []model.User
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.User, int64, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) int64); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepository_ListFriendByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFriendByEmail'
type MockRepository_ListFriendByEmail_Call struct {
	*mock.Call
}

// ListFriendByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockRepository_Expecter) ListFriendByEmail(ctx interface{}, email interface{}) *MockRepository_ListFriendByEmail_Call {
	return &MockRepository_ListFriendByEmail_Call{Call: _e.mock.On("ListFriendByEmail", ctx, email)}
}

func (_c *MockRepository_ListFriendByEmail_Call) Run(run func(ctx context.Context, email string)) *MockRepository_ListFriendByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_ListFriendByEmail_Call) Return(_a0 []model.User, _a1 int64, _a2 error) *MockRepository_ListFriendByEmail_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepository_ListFriendByEmail_Call) RunAndReturn(run func(context.Context, string) ([]model.User, int64, error)) *MockRepository_ListFriendByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// ListTwoEmailsCommonFriends provides a mock function with given fields: ctx, primaryEmail, secondadryEmail
func (_m *MockRepository) ListTwoEmailsCommonFriends(ctx context.Context, primaryEmail string, secondadryEmail string) ([]model.User, int64, error) {
	ret := _m.Called(ctx, primaryEmail, secondadryEmail)

	if len(ret) == 0 {
		panic("no return value specified for ListTwoEmailsCommonFriends")
	}

	var r0 []model.User
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]model.User, int64, error)); ok {
		return rf(ctx, primaryEmail, secondadryEmail)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []model.User); ok {
		r0 = rf(ctx, primaryEmail, secondadryEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) int64); ok {
		r1 = rf(ctx, primaryEmail, secondadryEmail)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, primaryEmail, secondadryEmail)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepository_ListTwoEmailsCommonFriends_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTwoEmailsCommonFriends'
type MockRepository_ListTwoEmailsCommonFriends_Call struct {
	*mock.Call
}

// ListTwoEmailsCommonFriends is a helper method to define mock.On call
//   - ctx context.Context
//   - primaryEmail string
//   - secondadryEmail string
func (_e *MockRepository_Expecter) ListTwoEmailsCommonFriends(ctx interface{}, primaryEmail interface{}, secondadryEmail interface{}) *MockRepository_ListTwoEmailsCommonFriends_Call {
	return &MockRepository_ListTwoEmailsCommonFriends_Call{Call: _e.mock.On("ListTwoEmailsCommonFriends", ctx, primaryEmail, secondadryEmail)}
}

func (_c *MockRepository_ListTwoEmailsCommonFriends_Call) Run(run func(ctx context.Context, primaryEmail string, secondadryEmail string)) *MockRepository_ListTwoEmailsCommonFriends_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockRepository_ListTwoEmailsCommonFriends_Call) Return(_a0 []model.User, _a1 int64, _a2 error) *MockRepository_ListTwoEmailsCommonFriends_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepository_ListTwoEmailsCommonFriends_Call) RunAndReturn(run func(context.Context, string, string) ([]model.User, int64, error)) *MockRepository_ListTwoEmailsCommonFriends_Call {
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
