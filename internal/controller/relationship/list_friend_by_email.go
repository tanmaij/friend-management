package relationship

import (
	"context"
	"log"

	"github.com/tanmaij/friend-management/internal/model"
)

// ListFriendByEmailInput represents the input required to list friends by email address
type ListFriendByEmailInput struct {
	Email string
}

// ListFriendByEmailOutput represents the output containing the list of friends and the total count
type ListFriendByEmailOutput struct {
	Friends []model.User
	Count   int64
}

// ListFriendByEmail handles the logic for listing friends for a given email address.
func (i *impl) ListFriendByEmail(ctx context.Context, inp ListFriendByEmailInput) (ListFriendByEmailOutput, error) {
	foundFriends, count, err := i.relationshipRepo.ListFriendByEmail(ctx, inp.Email)
	if err != nil {
		log.Printf("error listing friend by email: %v", err)
		return ListFriendByEmailOutput{}, err
	}

	return ListFriendByEmailOutput{
		Friends: foundFriends,
		Count:   count,
	}, nil
}
