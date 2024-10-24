package relationship

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/tanmaij/friend-management/internal/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Create inserts a new relationship into the database.
func (i *impl) Create(ctx context.Context, relationship model.Relationship) error {
	if err := relationship.Insert(ctx, i.sqlDB, boil.Whitelist(
		model.RelationshipColumns.RequesterID,
		model.RelationshipColumns.TargetID,
		model.RelationshipColumns.Type),
	); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
