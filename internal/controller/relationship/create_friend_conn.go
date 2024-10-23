package relationship

import (
	"context"
	"log"

	"github.com/friendsofgo/errors"
	"github.com/tanmaij/friend-management/internal/model"
	userRepo "github.com/tanmaij/friend-management/internal/repository/user"
)

// CreateFriendConnInp holds the input data for creating a friend connection, including the requester's email and the target's email
type CreateFriendConnInp struct {
	PrimaryEmail   string
	SecondaryEmail string
}

// CreateFriendConn handles the logic for creating a friend connection
func (i *impl) CreateFriendConn(ctx context.Context, inp CreateFriendConnInp) error {
	foundPrimaryUser, err := i.userRepo.GetByEmail(ctx, inp.PrimaryEmail)
	if err != nil {
		if errors.Is(err, userRepo.ErrUserNotFound) {
			return ErrUserNotFoundWithGivenEmail
		}

		log.Printf("error retrieving primary user: %v", err)
		return err
	}

	foundSecondaryUser, err := i.userRepo.GetByEmail(ctx, inp.SecondaryEmail)
	if err != nil {
		if errors.Is(err, userRepo.ErrUserNotFound) {
			return ErrUserNotFoundWithGivenEmail
		}

		log.Printf("error retrieving secondary user: %v", err)
		return err
	}

	existingRels, err := i.relationshipRepo.ListByTwoUserIDs(ctx, foundPrimaryUser.ID, foundSecondaryUser.ID)
	if err != nil {
		log.Printf("error listing relationships: %v", err)
		return err
	}

	if err := checkFriendship(existingRels); err != nil {
		return err
	}

	return i.relationshipRepo.Create(ctx, model.Relationship{
		RequesterID: foundPrimaryUser.ID,
		TargetID:    foundSecondaryUser.ID,
		Type:        model.RelationshipTypeFriend,
	})
}

func checkFriendship(rels []model.Relationship) error {
	for _, rel := range rels {
		switch rel.Type {
		case model.RelationshipTypeFriend:
			return ErrAlreadyFriends

		case model.RelationshipTypeBlock:
			return ErrAlreadyBlocked

		default:
		}
	}

	return nil
}
