package relationship

import (
	"context"

	"github.com/tanmaij/friend-management/internal/repository"
)

// Controller defines the interface for managing friend connections, including creating a new friend connection.
type Controller interface {
	// CreateFriendConn handles the logic for creating a friend connection.
	CreateFriendConn(ctx context.Context, inp CreateFriendConnInp) error
}

type impl struct {
	repo repository.Registry
}

// New creates a new instance of the Controller with the provided repository.
func New(repo repository.Registry) Controller {
	return &impl{repo: repo}
}
