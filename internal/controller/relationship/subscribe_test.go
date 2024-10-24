package relationship

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/assert"
	"github.com/tanmaij/friend-management/internal/model"
	relationshipRepo "github.com/tanmaij/friend-management/internal/repository/relationship"
	userRepo "github.com/tanmaij/friend-management/internal/repository/user"
)

func Test_impl_Subscribe(t *testing.T) {
	unexpectedErr := errors.New("an unexpected error")

	type mockCreateRelationship struct {
		isCalled  bool
		input     model.Relationship
		outputErr error
	}

	type mockGetUserByEmail struct {
		isCalled   bool
		inputEmail string
		outputUser model.User
		outputErr  error
	}

	type mockListRelsBetweenTwoEmails struct {
		isCalled            bool
		inputPrimaryID      int
		inputSecondaryID    int
		outputRelationships []model.Relationship
		outputErr           error
	}

	tcs := map[string]struct {
		input                        SubscribeInput
		mockCreateRelationship       mockCreateRelationship
		mockGetRequestorByEmail      mockGetUserByEmail
		mockGetTargetUserByEmail     mockGetUserByEmail
		mockListRelsBetweenTwoEmails mockListRelsBetweenTwoEmails
		expOutputErr                 error
	}{
		"success": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  nil,
			},
			mockGetTargetUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "target@example.com",
				outputUser: model.User{ID: 2},
				outputErr:  nil,
			},
			mockListRelsBetweenTwoEmails: mockListRelsBetweenTwoEmails{
				isCalled:            true,
				inputPrimaryID:      1,
				inputSecondaryID:    2,
				outputRelationships: []model.Relationship{},
			},
			mockCreateRelationship: mockCreateRelationship{
				isCalled:  true,
				input:     model.Relationship{RequesterID: 1, TargetID: 2, Type: model.RelationshipTypeSubscribe},
				outputErr: nil,
			},
			expOutputErr: nil,
		},
		"requestor_not_found": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputErr:  userRepo.ErrUserNotFound,
			},
			expOutputErr: ErrUserNotFoundWithGivenEmail,
		},
		"target_user_not_found": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputUser: model.User{ID: 1},
			},
			mockGetTargetUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "target@example.com",
				outputErr:  userRepo.ErrUserNotFound,
			},
			expOutputErr: ErrUserNotFoundWithGivenEmail,
		},
		"already_subscribed": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputUser: model.User{ID: 1},
			},
			mockGetTargetUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "target@example.com",
				outputUser: model.User{ID: 2},
			},
			mockListRelsBetweenTwoEmails: mockListRelsBetweenTwoEmails{
				isCalled:         true,
				inputPrimaryID:   1,
				inputSecondaryID: 2,
				outputRelationships: []model.Relationship{
					{RequesterID: 1, TargetID: 2, Type: model.RelationshipTypeSubscribe},
				},
			},
			expOutputErr: ErrAlreadySubscribed,
		},
		"unexpected_error_from_getting_requestor": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputErr:  unexpectedErr,
			},
			expOutputErr: unexpectedErr,
		},
		"unexpected_error_from_getting_target_user": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputUser: model.User{ID: 1},
			},
			mockGetTargetUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "target@example.com",
				outputErr:  unexpectedErr,
			},
			expOutputErr: unexpectedErr,
		},
		"unexpected_error_from_listing_rels": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputUser: model.User{ID: 1},
			},
			mockGetTargetUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "target@example.com",
				outputUser: model.User{ID: 2},
			},
			mockListRelsBetweenTwoEmails: mockListRelsBetweenTwoEmails{
				isCalled:         true,
				inputPrimaryID:   1,
				inputSecondaryID: 2,
				outputErr:        unexpectedErr,
			},
			expOutputErr: unexpectedErr,
		},
		"unexpected_error_from_creating_rel": {
			input: SubscribeInput{
				RequestorEmail: "requestor@example.com",
				TargetEmail:    "target@example.com",
			},
			mockGetRequestorByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "requestor@example.com",
				outputUser: model.User{ID: 1},
			},
			mockGetTargetUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "target@example.com",
				outputUser: model.User{ID: 2},
			},
			mockListRelsBetweenTwoEmails: mockListRelsBetweenTwoEmails{
				isCalled:            true,
				inputPrimaryID:      1,
				inputSecondaryID:    2,
				outputRelationships: []model.Relationship{},
			},
			mockCreateRelationship: mockCreateRelationship{
				isCalled:  true,
				input:     model.Relationship{RequesterID: 1, TargetID: 2, Type: model.RelationshipTypeSubscribe},
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
			relationshipRepoMock := new(relationshipRepo.MockRepository)

			if tc.mockCreateRelationship.isCalled {
				relationshipRepoMock.EXPECT().
					Create(emptyCtx, tc.mockCreateRelationship.input).
					Return(tc.mockCreateRelationship.outputErr)
			}

			if tc.mockGetRequestorByEmail.isCalled {
				userRepoMock.EXPECT().
					GetByEmail(emptyCtx, tc.mockGetRequestorByEmail.inputEmail).
					Return(tc.mockGetRequestorByEmail.outputUser, tc.mockGetRequestorByEmail.outputErr)
			}

			if tc.mockGetTargetUserByEmail.isCalled {
				userRepoMock.EXPECT().
					GetByEmail(emptyCtx, tc.mockGetTargetUserByEmail.inputEmail).
					Return(tc.mockGetTargetUserByEmail.outputUser, tc.mockGetTargetUserByEmail.outputErr)
			}

			if tc.mockListRelsBetweenTwoEmails.isCalled {
				relationshipRepoMock.EXPECT().
					ListByTwoUserIDs(
						emptyCtx,
						tc.mockListRelsBetweenTwoEmails.inputPrimaryID,
						tc.mockListRelsBetweenTwoEmails.inputSecondaryID).
					Return(
						tc.mockListRelsBetweenTwoEmails.outputRelationships,
						tc.mockListRelsBetweenTwoEmails.outputErr,
					)
			}

			subscriptionService := New(relationshipRepoMock, userRepoMock)

			// WHEN
			outputErr := subscriptionService.Subscribe(emptyCtx, tc.input)

			// THEN
			assert.Equal(t, tc.expOutputErr, outputErr)
		})
	}
}
