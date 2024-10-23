package relationship

import (
	"context"

	"github.com/tanmaij/friend-management/internal/model"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ListByTwoUserIDs lists relationships between two user ids
func (i *impl) ListByTwoUserIDs(ctx context.Context, primaryUserID, secondaryUserID int) ([]model.Relationship, error) {
	foundSlice, err := model.
		Relationships(
			qm.Where(model.RelationshipColumns.TargetID+" IN (?,?)", primaryUserID, secondaryUserID),
			qm.Where(model.RelationshipColumns.RequesterID+" IN (?,?)", primaryUserID, secondaryUserID)).
		All(ctx, i.sqlDB)
	if err != nil {
		return []model.Relationship{}, errors.WithStack(err)
	}

	var rs []model.Relationship = make([]model.Relationship, len(foundSlice))
	for i, rel := range foundSlice {
		if rel == nil {
			continue
		}

		rs[i] = *rel
	}

	return rs, nil
}
