package relationship

import (
	"context"
	"log"

	"github.com/tanmaij/friend-management/internal/model"
)

// ListTwoEmailCommonFriendsInput represents the input required to list friends by email address
type ListTwoEmailCommonFriendsInput struct {
	PrimaryEmail   string
	SecondaryEmail string
}

// ListFriendByEmailOutput represents the output containing the list of friends and the total count
type ListTwoEmailCommonFriendsOutput struct {
	Friends []model.User
	Count   int64
}

// ListTwoEmailCommonFriends retrieves a list of common friends between two email addresses.
func (i *impl) ListTwoEmailCommonFriends(ctx context.Context, inp ListTwoEmailCommonFriendsInput) (ListTwoEmailCommonFriendsOutput, error) {
	foundFriends, count, err := i.relationshipRepo.ListTwoEmailsCommonFriends(ctx, inp.PrimaryEmail, inp.SecondaryEmail)
	if err != nil {
		log.Printf("error listing friend by email: %v", err)
		return ListTwoEmailCommonFriendsOutput{}, err
	}

	return ListTwoEmailCommonFriendsOutput{
		Friends: foundFriends,
		Count:   count,
	}, nil
}
