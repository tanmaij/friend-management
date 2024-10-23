package relationship

import (
	"context"
	"database/sql"

	"github.com/tanmaij/friend-management/internal/model"
)

// Repository accesses relationsip data
type Repository interface {
	// Create inserts a new relationship into the database.
	Create(ctx context.Context, relationship model.Relationship) error

	// ListByTwoUserIDs lists relationships between two user ids
	ListByTwoUserIDs(ctx context.Context, primaryUserID, secondaryUserID int) ([]model.Relationship, error)
}

type impl struct {
	sqlDB *sql.DB
}

// New creates a new Repository.
func New(sqlDB *sql.DB) Repository {
	return &impl{sqlDB: sqlDB}
}
