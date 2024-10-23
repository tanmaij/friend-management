package relationship

import (
	"context"
	"fmt"

	"github.com/tanmaij/friend-management/internal/model"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ListFriendByEmail lists friends for an email, returning a list of users and total count
func (i *impl) ListFriendByEmail(ctx context.Context, email string) ([]model.User, int64, error) {
	var queries = []qm.QueryMod{
		qm.Select(model.TableNames.Users + ".*"),
		qm.InnerJoin(
			fmt.Sprintf("%s AS found_user ON found_user.%s = ?", model.TableNames.Users, model.UserColumns.Email),
			email,
		),
		qm.InnerJoin(
			fmt.Sprintf("%s ON %s = ? AND found_user.%s IN (%s, %s)",
				model.TableNames.Relationships,
				model.RelationshipColumns.Type,
				model.UserColumns.ID,
				model.RelationshipColumns.RequesterID,
				model.RelationshipColumns.TargetID),
			model.RelationshipTypeFriend,
		),
		qm.Where(
			fmt.Sprintf("%s.%s <> ?", model.TableNames.Users, model.UserColumns.Email),
			email,
		),
		qm.GroupBy(fmt.Sprintf("%s.%s", model.TableNames.Users, model.UserColumns.ID)),
	}

	foundSlice, err := model.Users(
		queries...,
	).All(ctx, i.sqlDB)

	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	rs := make([]model.User, len(foundSlice))
	for i, user := range foundSlice {
		if user == nil {
			continue
		}

		rs[i] = *user
	}

	return rs, int64(len(rs)), nil
}
