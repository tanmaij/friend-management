package user

import (
	"context"
	"database/sql"

	"github.com/tanmaij/friend-management/internal/model"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetByEmail
func (i *impl) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := model.Users(qm.Where(model.UserColumns.Email+" = ?", email)).One(ctx, i.sqlDB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, ErrUserNotFound
		}

		return model.User{}, errors.WithStack(err)
	}

	return *user, nil
}
