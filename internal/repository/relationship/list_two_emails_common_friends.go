package relationship

import (
	"context"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/tanmaij/friend-management/internal/model"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ListTwoEmailsCommonFriends retrieves the list of common friends between two email addresses from the database
func (i *impl) ListTwoEmailsCommonFriends(ctx context.Context, primaryEmail string, secondadryEmail string) ([]model.User, int64, error) {
	var queries = []qm.QueryMod{
		qm.Select(model.TableNames.Users + ".*"),
		qm.InnerJoin(
			fmt.Sprintf("%s AS primary_user ON primary_user.%s = ?", model.TableNames.Users, model.UserColumns.Email),
			primaryEmail,
		),
		qm.InnerJoin(
			fmt.Sprintf("%s AS seconday_user ON seconday_user.%s = ?", model.TableNames.Users, model.UserColumns.Email),
			secondadryEmail,
		),
		qm.InnerJoin(
			fmt.Sprintf("%s ON %s = ? AND (primary_user.%s IN (%s, %s) OR seconday_user.%s IN (%s, %s))",
				model.TableNames.Relationships,
				model.RelationshipColumns.Type,
				model.UserColumns.ID,
				model.RelationshipColumns.RequesterID,
				model.RelationshipColumns.TargetID,
				model.UserColumns.ID,
				model.RelationshipColumns.RequesterID,
				model.RelationshipColumns.TargetID),
			model.RelationshipTypeFriend,
		),
		qm.Where(
			fmt.Sprintf("%s.%s <> ? AND %s.%s <> ?",
				model.TableNames.Users, model.UserColumns.Email,
				model.TableNames.Users, model.UserColumns.Email),
			primaryEmail, secondadryEmail,
		),
		qm.GroupBy(fmt.Sprintf("%s.%s", model.TableNames.Users, model.UserColumns.ID)),
		qm.OrderBy(fmt.Sprintf("%s.%s ASC", model.TableNames.Users, model.UserColumns.ID)),
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
