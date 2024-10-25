package user

import (
	"context"
	userRepo "github.com/tanmaij/friend-management/internal/repository/user"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/assert"
	"github.com/tanmaij/friend-management/internal/model"
)

func Test_impl_Create(t *testing.T) {
	unexpectedErr := errors.New("an unexpected error")

	type mockExistsByEmail struct {
		isCalled     bool
		inputEmail   string
		outputExists bool
		outputErr    error
	}

	type mockCreateUser struct {
		isCalled  bool
		input     model.User
		outputErr error
	}

	tcs := map[string]struct {
		input             CreateInput
		mockExistsByEmail mockExistsByEmail
		mockCreateUser    mockCreateUser
		expOutputErr      error
	}{
		"success": {
			input: CreateInput{
				Email: "user@example.com",
			},
			mockExistsByEmail: mockExistsByEmail{
				isCalled:     true,
				inputEmail:   "user@example.com",
				outputExists: false,
				outputErr:    nil,
			},
			mockCreateUser: mockCreateUser{
				isCalled:  true,
				input:     model.User{Email: "user@example.com"},
				outputErr: nil,
			},
			expOutputErr: nil,
		},
		"user_already_exists": {
			input: CreateInput{
				Email: "user@example.com",
			},
			mockExistsByEmail: mockExistsByEmail{
				isCalled:     true,
				inputEmail:   "user@example.com",
				outputExists: true,
				outputErr:    nil,
			},
			expOutputErr: ErrUserAlreadyExists,
		},
		"error_checking_existing_user": {
			input: CreateInput{
				Email: "user@example.com",
			},
			mockExistsByEmail: mockExistsByEmail{
				isCalled:   true,
				inputEmail: "user@example.com",
				outputErr:  unexpectedErr,
			},
			expOutputErr: unexpectedErr,
		},
		"error_creating_user": {
			input: CreateInput{
				Email: "user@example.com",
			},
			mockExistsByEmail: mockExistsByEmail{
				isCalled:     true,
				inputEmail:   "user@example.com",
				outputExists: false,
				outputErr:    nil,
			},
			mockCreateUser: mockCreateUser{
				isCalled:  true,
				input:     model.User{Email: "user@example.com"},
				outputErr: unexpectedErr,
			},
			expOutputErr: unexpectedErr,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			emptyCtx := context.Background()

			userRepoMock := new(userRepo.MockRepository)

			if tc.mockExistsByEmail.isCalled {
				userRepoMock.EXPECT().
					ExistsByEmail(emptyCtx, tc.mockExistsByEmail.inputEmail).
					Return(tc.mockExistsByEmail.outputExists, tc.mockExistsByEmail.outputErr)
			}

			if tc.mockCreateUser.isCalled {
				userRepoMock.EXPECT().
					Create(emptyCtx, tc.mockCreateUser.input).
					Return(tc.mockCreateUser.outputErr)
			}

			ctrl := New(userRepoMock)

			// WHEN
			actErr := ctrl.Create(emptyCtx, tc.input)

			// THEN
			if tc.expOutputErr != nil {
				assert.EqualError(t, actErr, tc.expOutputErr.Error())
			} else {
				assert.NoError(t, actErr)
			}

			userRepoMock.AssertExpectations(t)
		})
	}
}
