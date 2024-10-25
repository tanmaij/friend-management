package user

import (
	"context"
	"database/sql"

	"github.com/tanmaij/friend-management/internal/model"
)

// Repository accesses user data
type Repository interface {
	// GetByEmail returns user from database with given email
	GetByEmail(ctx context.Context, email string) (model.User, error)

	// Create inserts a user to database
	Create(ctx context.Context, user model.User) error

	// ExistsByEmail checks if user exists with given email
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type impl struct {
	sqlDB *sql.DB
}

// New creates a new Repository.
func New(sqlDB *sql.DB) Repository {
	return &impl{sqlDB: sqlDB}
}
