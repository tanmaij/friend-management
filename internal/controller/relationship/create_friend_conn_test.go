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

func Test_impl_CreateFriendConn(t *testing.T) {
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
		input                        CreateFriendConnInp
		mockCreateRelationship       mockCreateRelationship
		mockGetPrimaryUserByEmail    mockGetUserByEmail
		mockGetSecondaryUserByEmail  mockGetUserByEmail
		mockListRelsBetweenTwoEmails mockListRelsBetweenTwoEmails
		expOutputErr                 error
	}{
		"success": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  nil,
			},
			mockGetSecondaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user2@example.com",
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
				input:     model.Relationship{RequesterID: 1, TargetID: 2, Type: model.RelationshipTypeFriend},
				outputErr: nil,
			},
			expOutputErr: nil,
		},
		"primary_user_not_found_with_given_email": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputErr:  userRepo.ErrUserNotFound,
			},
			expOutputErr: ErrUserNotFoundWithGivenEmail,
		},
		"secondary_user_not_found_with_given_email": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
			},
			mockGetSecondaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user2@example.com",
				outputErr:  userRepo.ErrUserNotFound,
			},
			expOutputErr: ErrUserNotFoundWithGivenEmail,
		},
		"can_not_be_friends_due_to_already_friends": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  nil,
			},
			mockGetSecondaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user2@example.com",
				outputUser: model.User{ID: 2},
				outputErr:  nil,
			},
			mockListRelsBetweenTwoEmails: mockListRelsBetweenTwoEmails{
				isCalled:         true,
				inputPrimaryID:   1,
				inputSecondaryID: 2,
				outputRelationships: []model.Relationship{
					{Type: model.RelationshipTypeFriend},
				},
			},
			expOutputErr: ErrAlreadyFriends,
		},
		"can_not_be_friends_due_to_already_blocked": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  nil,
			},
			mockGetSecondaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user2@example.com",
				outputUser: model.User{ID: 2},
				outputErr:  nil,
			},
			mockListRelsBetweenTwoEmails: mockListRelsBetweenTwoEmails{
				isCalled:         true,
				inputPrimaryID:   1,
				inputSecondaryID: 2,
				outputRelationships: []model.Relationship{
					{Type: model.RelationshipTypeBlock},
				},
			},
			expOutputErr: ErrAlreadyBlocked,
		},
		"unexpected_error_from_getting_primary_user": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  unexpectedErr,
			},
			expOutputErr: unexpectedErr,
		},
		"unexpected_error_from_getting_secondary_user": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  nil,
			},
			mockGetSecondaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user2@example.com",
				outputErr:  unexpectedErr,
			},
			expOutputErr: unexpectedErr,
		},
		"unexpected_error_from_listing_rels": {
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  nil,
			},
			mockGetSecondaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user2@example.com",
				outputUser: model.User{ID: 2},
				outputErr:  nil,
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
			input: CreateFriendConnInp{
				RequesterEmail: "user1@example.com",
				TargetEmail:    "user2@example.com",
			},
			mockGetPrimaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user1@example.com",
				outputUser: model.User{ID: 1},
				outputErr:  nil,
			},
			mockGetSecondaryUserByEmail: mockGetUserByEmail{
				isCalled:   true,
				inputEmail: "user2@example.com",
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
				input:     model.Relationship{RequesterID: 1, TargetID: 2, Type: model.RelationshipTypeFriend},
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

			if tc.mockGetPrimaryUserByEmail.isCalled {
				userRepoMock.EXPECT().
					GetByEmail(emptyCtx, tc.mockGetPrimaryUserByEmail.inputEmail).
					Return(tc.mockGetPrimaryUserByEmail.outputUser, tc.mockGetPrimaryUserByEmail.outputErr)
			}

			if tc.mockGetSecondaryUserByEmail.isCalled {
				userRepoMock.EXPECT().
					GetByEmail(emptyCtx, tc.mockGetSecondaryUserByEmail.inputEmail).
					Return(tc.mockGetSecondaryUserByEmail.outputUser, tc.mockGetSecondaryUserByEmail.outputErr)
			}

			if tc.mockListRelsBetweenTwoEmails.isCalled {
				relationshipRepoMock.EXPECT().
					ListByTwoUserIDs(
						emptyCtx,
						tc.mockListRelsBetweenTwoEmails.inputPrimaryID,
						tc.mockListRelsBetweenTwoEmails.inputSecondaryID).
					Return(
						tc.mockListRelsBetweenTwoEmails.outputRelationships,
						tc.mockListRelsBetweenTwoEmails.outputErr)
			}

			ctrl := New(relationshipRepoMock, userRepoMock)

			// WHEN
			actErr := ctrl.CreateFriendConn(emptyCtx, tc.input)

			// THEN
			if tc.expOutputErr != nil {
				assert.EqualError(t, actErr, tc.expOutputErr.Error())
			} else {
				assert.NoError(t, actErr)
			}

			userRepoMock.AssertExpectations(t)
			relationshipRepoMock.AssertExpectations(t)
		})
	}
}
