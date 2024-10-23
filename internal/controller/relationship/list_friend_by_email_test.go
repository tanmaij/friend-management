package relationship

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/assert"
	"github.com/tanmaij/friend-management/internal/model"
	relationshipRepo "github.com/tanmaij/friend-management/internal/repository/relationship"
)

func Test_impl_ListFriendByEmail(t *testing.T) {
	type mockListFriendByEmail struct {
		isCalled      bool
		inputEmail    string
		outputFriends []model.User
		outputCount   int64
		outputErr     error
	}

	tcs := map[string]struct {
		input                 ListFriendByEmailInput
		mockListFriendByEmail mockListFriendByEmail
		expOutput             ListFriendByEmailOutput
		expOutputErr          error
	}{
		"success": {
			input: ListFriendByEmailInput{Email: "user@example.com"},
			mockListFriendByEmail: mockListFriendByEmail{
				isCalled:   true,
				inputEmail: "user@example.com",
				outputFriends: []model.User{
					{ID: 1, Email: "friend1@example.com"},
					{ID: 2, Email: "friend2@example.com"},
				},
				outputCount: 2,
				outputErr:   nil,
			},
			expOutput: ListFriendByEmailOutput{
				Friends: []model.User{
					{ID: 1, Email: "friend1@example.com"},
					{ID: 2, Email: "friend2@example.com"},
				},
				Count: 2,
			},
			expOutputErr: nil,
		},
		"error": {
			input: ListFriendByEmailInput{Email: "user@example.com"},
			mockListFriendByEmail: mockListFriendByEmail{
				isCalled:      true,
				inputEmail:    "user@example.com",
				outputFriends: nil,
				outputCount:   0,
				outputErr:     errors.New("database error"),
			},
			expOutput:    ListFriendByEmailOutput{},
			expOutputErr: errors.New("database error"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			emptyCtx := context.Background()
			relationshipRepoMock := new(relationshipRepo.MockRepository)

			if tc.mockListFriendByEmail.isCalled {
				relationshipRepoMock.EXPECT().
					ListFriendByEmail(emptyCtx, tc.mockListFriendByEmail.inputEmail).
					Return(tc.mockListFriendByEmail.outputFriends, tc.mockListFriendByEmail.outputCount, tc.mockListFriendByEmail.outputErr)
			}
			ctrl := New(relationshipRepoMock, nil)

			// WHEN
			actOutput, actErr := ctrl.ListFriendByEmail(emptyCtx, tc.input)

			// THEN
			if tc.expOutputErr != nil {
				assert.EqualError(t, actErr, tc.expOutputErr.Error())
			} else {
				assert.NoError(t, actErr)
				assert.Equal(t, tc.expOutput, actOutput)
			}

			relationshipRepoMock.AssertExpectations(t)
		})
	}
}
