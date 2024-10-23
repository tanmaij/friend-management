package v1

import (
	"net/http"

	relationshipCtrl "github.com/tanmaij/friend-management/internal/controller/relationship"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
)

var (
	errInternalServer             = httpUtil.Error{Status: http.StatusInternalServerError, Code: "internal_server_error", Message: "Internal server error"}
	errInvalidRequestBody         = httpUtil.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Message: "Invalid request body"}
	errInvalidGivenEmail          = httpUtil.Error{Status: http.StatusBadRequest, Code: "invalid_given_email", Message: "Invalid given email"}
	errUserNotFoundWithGivenEmail = httpUtil.Error{Status: http.StatusBadRequest, Code: "user_not_found_with_given_email", Message: "User not found with given email"}
	errAlreadyFriends             = httpUtil.Error{Status: http.StatusBadRequest, Code: "already_friends", Message: "Already friends"}
	errAlreadyBlocked             = httpUtil.Error{Status: http.StatusBadRequest, Code: "already_blocked", Message: "Already blocked"}
)

func convertErrorFromController(err error) httpUtil.Error {
	switch err {
	case relationshipCtrl.ErrUserNotFoundWithGivenEmail:
		return errUserNotFoundWithGivenEmail
	case relationshipCtrl.ErrAlreadyBlocked:
		return errAlreadyBlocked
	case relationshipCtrl.ErrAlreadyFriends:
		return errAlreadyFriends
	default:
		return errInternalServer
	}
}
