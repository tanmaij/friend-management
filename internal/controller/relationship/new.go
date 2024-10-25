package relationship

import (
	"context"

	"github.com/tanmaij/friend-management/internal/repository/relationship"
	"github.com/tanmaij/friend-management/internal/repository/user"
)

// Controller managing relationship business
type Controller interface {
	// CreateFriendConn handles the logic for creating a friend connection
	CreateFriendConn(ctx context.Context, inp CreateFriendConnInp) error

	// ListFriendByEmail handles the logic for a listing friends for an email address
	ListFriendByEmail(ctx context.Context, inp ListFriendByEmailInput) (ListFriendByEmailOutput, error)

	// ListTwoEmailCommonFriends retrieves a list of common friends between two email addresses.
	ListTwoEmailCommonFriends(ctx context.Context, inp ListTwoEmailCommonFriendsInput) (ListTwoEmailCommonFriendsOutput, error)

	// Subscribe subscribes for updates from email
	Subscribe(ctx context.Context, inp SubscribeInput) error

	// Block blocks updates from an email address
	Block(ctx context.Context, inp BlockInput) error

	ListEligibleRecipientEmailsFromUpdate(ctx context.Context, inp ListEligibleRecipientEmailsFromUpdateInput) (ListEligibleRecipientEmailsFromUpdateOutput, error)
}

type impl struct {
	relationshipRepo relationship.Repository
	userRepo         user.Repository
}

// New creates a new instance of the Controller with the provided repositories
func New(relationshipRepo relationship.Repository, userRepo user.Repository) Controller {
	return &impl{relationshipRepo: relationshipRepo, userRepo: userRepo}
}
