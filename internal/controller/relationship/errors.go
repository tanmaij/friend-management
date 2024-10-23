package relationship

import "github.com/friendsofgo/errors"

var (
	ErrUserNotFoundWithGivenEmail = errors.New("user not found with given email")
	ErrAlreadyFriends             = errors.New("already friends")
	ErrAlreadyBlocked             = errors.New("already blocked")
)
