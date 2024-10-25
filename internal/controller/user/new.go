package user

import (
	"context"

	"github.com/tanmaij/friend-management/internal/repository/user"
)

// Controller managing user business
type Controller interface {
	// Create handles the logic for creating user
	Create(ctx context.Context, inp CreateInput) error
}

type impl struct {
	userRepo user.Repository
}

// New creates a new instance of the Controller with the provided repositories
func New(userRepo user.Repository) Controller {
	return &impl{userRepo: userRepo}
}
