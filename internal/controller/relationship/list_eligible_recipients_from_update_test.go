package relationship

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/assert"
	relationshipRepo "github.com/tanmaij/friend-management/internal/repository/relationship"
)

func Test_impl_ListEligibleRecipientEmailsFromUpdate(t *testing.T) {
	type mockFindEligibleRecipientEmails struct {
		isCalled        bool
		inputEmail      string
		mentionedEmails []string
		output          []string
		outputErr       error
	}

	tcs := map[string]struct {
		input            ListEligibleRecipientEmailsFromUpdateInput
		mockFindEligible mockFindEligibleRecipientEmails
		expOutput        ListEligibleRecipientEmailsFromUpdateOutput
		expOutputErr     error
	}{
		"success": {
			input: ListEligibleRecipientEmailsFromUpdateInput{
				SenderEmail: "user@example.com",
				Text:        "Hello friend1@example.com and friend2@example.com",
			},
			mockFindEligible: mockFindEligibleRecipientEmails{
				isCalled:        true,
				inputEmail:      "user@example.com",
				mentionedEmails: []string{"friend1@example.com", "friend2@example.com"},
				output:          []string{"friend1@example.com", "friend2@example.com"},
				outputErr:       nil,
			},
			expOutput: ListEligibleRecipientEmailsFromUpdateOutput{
				Recipients: []string{"friend1@example.com", "friend2@example.com"},
			},
			expOutputErr: nil,
		},
		"error": {
			input: ListEligibleRecipientEmailsFromUpdateInput{
				SenderEmail: "user@example.com",
				Text:        "Hello",
			},
			mockFindEligible: mockFindEligibleRecipientEmails{
				isCalled:        true,
				inputEmail:      "user@example.com",
				mentionedEmails: []string{},
				output:          nil,
				outputErr:       errors.New("database error"),
			},
			expOutput:    ListEligibleRecipientEmailsFromUpdateOutput{},
			expOutputErr: errors.New("database error"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			emptyCtx := context.Background()
			relationshipRepoMock := new(relationshipRepo.MockRepository)

			if tc.mockFindEligible.isCalled {
				relationshipRepoMock.EXPECT().
					FindEligibleRecipientEmailsWithMentioned(emptyCtx, tc.mockFindEligible.inputEmail, tc.mockFindEligible.mentionedEmails).
					Return(tc.mockFindEligible.output, tc.mockFindEligible.outputErr)
			}

			ctrl := New(relationshipRepoMock, nil)

			// WHEN
			actOutput, actErr := ctrl.ListEligibleRecipientEmailsFromUpdate(emptyCtx, tc.input)

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
