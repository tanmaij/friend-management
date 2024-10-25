package user

import (
	"context"

	"github.com/tanmaij/friend-management/internal/model"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ExistsByEmail checks if user exists with given email
func (i *impl) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := model.Users(qm.Where(model.UserColumns.Email+" = ?", email)).Exists(ctx, i.sqlDB)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return exists, nil
}
