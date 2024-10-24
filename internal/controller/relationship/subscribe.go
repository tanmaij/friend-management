package relationship

import (
	"context"
	"log"

	"github.com/friendsofgo/errors"
	"github.com/tanmaij/friend-management/internal/model"
	userRepo "github.com/tanmaij/friend-management/internal/repository/user"
)

// SubscribeInput
type SubscribeInput struct {
	RequestorEmail string
	TargetEmail    string
}

// Subscribe
func (i *impl) Subscribe(ctx context.Context, inp SubscribeInput) error {
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

	if err := checkValidSubscription(foundRequestor, foundTargetUser, existingRels); err != nil {
		return err
	}

	return i.relationshipRepo.Create(ctx, model.Relationship{
		RequesterID: foundRequestor.ID,
		TargetID:    foundTargetUser.ID,
		Type:        model.RelationshipTypeSubscribe,
	})
}

func checkValidSubscription(foundRequestor model.User, foundTargetUser model.User, rels []model.Relationship) error {
	for _, rel := range rels {
		switch rel.Type {
		case model.RelationshipTypeSubscribe:
			if rel.RequesterID == foundRequestor.ID && rel.TargetID == foundTargetUser.ID {
				return ErrAlreadySubscribed
			}

		default:
		}
	}

	return nil
}
