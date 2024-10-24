package relationship

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/assert"
	"github.com/tanmaij/friend-management/internal/model"
	relationshipRepo "github.com/tanmaij/friend-management/internal/repository/relationship"
)

func Test_impl_ListTwoEmailCommonFriends(t *testing.T) {
	var dbErr = errors.New("db err")
	type mockListTwoEmailsCommonFriends struct {
		isCalled       bool
		inputPrimary   string
		inputSecondary string
		outputFriends  []model.User
		outputCount    int64
		outputErr      error
	}

	tcs := map[string]struct {
		input                          ListTwoEmailCommonFriendsInput
		mockListTwoEmailsCommonFriends mockListTwoEmailsCommonFriends
		expOutput                      ListTwoEmailCommonFriendsOutput
		expOutputErr                   error
	}{
		"success": {
			input: ListTwoEmailCommonFriendsInput{
				PrimaryEmail:   "user1@example.com",
				SecondaryEmail: "user2@example.com",
			},
			mockListTwoEmailsCommonFriends: mockListTwoEmailsCommonFriends{
				isCalled:       true,
				inputPrimary:   "user1@example.com",
				inputSecondary: "user2@example.com",
				outputFriends: []model.User{
					{ID: 1, Email: "commonfriend1@example.com"},
					{ID: 2, Email: "commonfriend2@example.com"},
				},
				outputCount: 2,
				outputErr:   nil,
			},
			expOutput: ListTwoEmailCommonFriendsOutput{
				Friends: []model.User{
					{ID: 1, Email: "commonfriend1@example.com"},
					{ID: 2, Email: "commonfriend2@example.com"},
				},
				Count: 2,
			},
			expOutputErr: nil,
		},
		"error": {
			input: ListTwoEmailCommonFriendsInput{
				PrimaryEmail:   "user1@example.com",
				SecondaryEmail: "user2@example.com",
			},
			mockListTwoEmailsCommonFriends: mockListTwoEmailsCommonFriends{
				isCalled:       true,
				inputPrimary:   "user1@example.com",
				inputSecondary: "user2@example.com",
				outputFriends:  nil,
				outputCount:    0,
				outputErr:      dbErr,
			},
			expOutput:    ListTwoEmailCommonFriendsOutput{},
			expOutputErr: dbErr,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			emptyCtx := context.Background()
			relationshipRepoMock := new(relationshipRepo.MockRepository)

			if tc.mockListTwoEmailsCommonFriends.isCalled {
				relationshipRepoMock.EXPECT().
					ListTwoEmailsCommonFriends(emptyCtx, tc.mockListTwoEmailsCommonFriends.inputPrimary, tc.mockListTwoEmailsCommonFriends.inputSecondary).
					Return(tc.mockListTwoEmailsCommonFriends.outputFriends, tc.mockListTwoEmailsCommonFriends.outputCount, tc.mockListTwoEmailsCommonFriends.outputErr)
			}

			ctrl := New(relationshipRepoMock, nil)

			// WHEN
			actOutput, actErr := ctrl.ListTwoEmailCommonFriends(emptyCtx, tc.input)

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
