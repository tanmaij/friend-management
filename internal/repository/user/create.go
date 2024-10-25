package user

import (
	"context"

	"github.com/tanmaij/friend-management/internal/model"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Create inserts a user to database
func (i *impl) Create(ctx context.Context, user model.User) error {
	if err := user.Insert(ctx, i.sqlDB, boil.Whitelist(model.UserColumns.Email)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
