package user

import (
	"github.com/friendsofgo/errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)
