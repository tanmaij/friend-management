package relationship

import (
	"context"
	"log"

	"github.com/friendsofgo/errors"
	"github.com/tanmaij/friend-management/internal/model"
	userRepo "github.com/tanmaij/friend-management/internal/repository/user"
)

// BlockInput includes data to block
type BlockInput struct {
	RequestorEmail string
	TargetEmail    string
}

// Block blocks updates from an email address
func (i *impl) Block(ctx context.Context, inp BlockInput) error {
	foundRequestor, err := i.userRepo.GetByEmail(ctx, inp.RequestorEmail)
	if err != nil {
		if errors.Is(err, userRepo.ErrUserNotFound) {
			return ErrUserNotFoundWithGivenEmail
		}

		log.Printf("error retrieving requestor: %v", err)
		return err
	}

	foundTargetUser, err := i.userRepo.GetByEmail(ctx, inp.TargetEmail)
	if err != nil {
		if errors.Is(err, userRepo.ErrUserNotFound) {
			return ErrUserNotFoundWithGivenEmail
		}

		log.Printf("error retrieving target user: %v", err)
		return err
	}

	existingRels, err := i.relationshipRepo.ListByTwoUserIDs(ctx, foundRequestor.ID, foundTargetUser.ID)
	if err != nil {
		log.Printf("error listing relationships: %v", err)
		return err
	}

	if err := checkValidBlocking(foundRequestor, foundTargetUser, existingRels); err != nil {
		return err
	}

	if err := i.relationshipRepo.Create(ctx, model.Relationship{
		RequesterID: foundRequestor.ID,
		TargetID:    foundTargetUser.ID,
		Type:        model.RelationshipTypeBlock,
	}); err != nil {
		log.Printf("error creating relationship: %v", err)
		return err
	}

	return nil
}

func checkValidBlocking(foundRequestor model.User, foundTargetUser model.User, rels []model.Relationship) error {
	for _, rel := range rels {
		switch rel.Type {
		case model.RelationshipTypeBlock:
			if rel.RequesterID == foundRequestor.ID && rel.TargetID == foundTargetUser.ID {
				return ErrAlreadyBlocked
			}

		default:
		}
	}

	return nil
}
