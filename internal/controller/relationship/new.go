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
}

type impl struct {
	relationshipRepo relationship.Repository
	userRepo         user.Repository
}

// New creates a new instance of the Controller with the provided repositories
func New(relationshipRepo relationship.Repository, userRepo user.Repository) Controller {
	return &impl{relationshipRepo: relationshipRepo, userRepo: userRepo}
}
