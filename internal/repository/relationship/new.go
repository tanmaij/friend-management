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

	// ListFriendByEmail lists friends for an email, returning a list of users and total count
	ListFriendByEmail(ctx context.Context, email string) ([]model.User, int64, error)

	// ListTwoEmailsCommonFriends retrieves the list of common friends between two email addresses from the database
	ListTwoEmailsCommonFriends(ctx context.Context, primaryEmail string, secondadryEmail string) ([]model.User, int64, error)

	FindEligibleRecipientEmailsWithMentioned(ctx context.Context, sender string, mentionedEmails []string) ([]string, error)
}

type impl struct {
	sqlDB *sql.DB
}

// New creates a new Repository.
func New(sqlDB *sql.DB) Repository {
	return &impl{sqlDB: sqlDB}
}
