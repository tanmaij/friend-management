package user

import (
	"context"
	"database/sql"

	"github.com/tanmaij/friend-management/internal/model"
)

// Repository accesses user data
type Repository interface {
	// GetByEmail get an user from database with given email
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

type impl struct {
	sqlDB *sql.DB
}

// New creates a new Repository.
func New(sqlDB *sql.DB) Repository {
	return &impl{sqlDB: sqlDB}
}
