package v1

import (
	"net/http"

	relationshipCtrl "github.com/tanmaij/friend-management/internal/controller/relationship"
	userCtrl "github.com/tanmaij/friend-management/internal/controller/user"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
)

var (
	errInternalServer                  = httpUtil.Error{Status: http.StatusInternalServerError, Code: "internal_server_error", Message: "Internal server error"}
	errInvalidRequestBody              = httpUtil.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Message: "Invalid request body"}
	errInvalidGivenEmail               = httpUtil.Error{Status: http.StatusBadRequest, Code: "invalid_given_email", Message: "Invalid given email"}
	errUserNotFoundWithGivenEmail      = httpUtil.Error{Status: http.StatusBadRequest, Code: "user_not_found_with_given_email", Message: "User not found with given email"}
	errAlreadyFriends                  = httpUtil.Error{Status: http.StatusBadRequest, Code: "already_friends", Message: "Already friends"}
	errAlreadyBlocked                  = httpUtil.Error{Status: http.StatusBadRequest, Code: "already_blocked", Message: "Already blocked"}
	errCannotBeFriendWithSelf          = httpUtil.Error{Status: http.StatusBadRequest, Code: "cannot_be_friend_with_self", Message: "Cannot be friend with self"}
	errCannotGetCommonFriendWithSelf   = httpUtil.Error{Status: http.StatusBadRequest, Code: "cannot_get_common_friends_with_self", Message: "Cannot get common friends with self"}
	errRequestorEmailIsRequired        = httpUtil.Error{Status: http.StatusBadRequest, Code: "requestor_email_is_required", Message: "Requestor email is required"}
	errTargetEmailIsRequired           = httpUtil.Error{Status: http.StatusBadRequest, Code: "target_email_is_required", Message: "Target email is required"}
	errInvalidRequestorEmail           = httpUtil.Error{Status: http.StatusBadRequest, Code: "invalid_requestor_email", Message: "Invalid requestor email"}
	errInvalidTargetEmail              = httpUtil.Error{Status: http.StatusBadRequest, Code: "invalid_target_email", Message: "Invalid target email"}
	errRequestorNotFoundWithGivenEmail = httpUtil.Error{Status: http.StatusBadRequest, Code: "requestor_not_found_with_given_email", Message: "Requestor not found with given email"}
	errTargetNotFoundWithGivenEmail    = httpUtil.Error{Status: http.StatusBadRequest, Code: "target_user_not_found_with_given_email", Message: "Target user not found with given email"}
	errCannotSelfSubscribe             = httpUtil.Error{Status: http.StatusBadRequest, Code: "cannot_self_subscribe", Message: "Cannot self subcribe"}
	errAlreadySubscribed               = httpUtil.Error{Status: http.StatusBadRequest, Code: "already_subscribed", Message: "Already subscribed"}
	errCannotSelfBlock                 = httpUtil.Error{Status: http.StatusBadRequest, Code: "cannot_self_block", Message: "Cannot self block"}
	errSenderEmailIsRequired           = httpUtil.Error{Status: http.StatusBadRequest, Code: "cannot_sender_email_is_required", Message: "Sender email is required"}
	errUserAlreadyExists               = httpUtil.Error{Status: http.StatusBadRequest, Code: "user_already_exists", Message: "User already exists"}
)

func convertErrorFromController(err error) httpUtil.Error {
	switch err {
	case relationshipCtrl.ErrUserNotFoundWithGivenEmail:
		return errUserNotFoundWithGivenEmail
	case relationshipCtrl.ErrAlreadyBlocked:
		return errAlreadyBlocked
	case relationshipCtrl.ErrAlreadyFriends:
		return errAlreadyFriends
	case relationshipCtrl.ErrAlreadySubscribed:
		return errAlreadySubscribed
	case relationshipCtrl.ErrRequestorNotFoundWithGivenEmail:
		return errRequestorNotFoundWithGivenEmail
	case relationshipCtrl.ErrTargetUserNotFoundWithGivenEmail:
		return errTargetNotFoundWithGivenEmail

	case userCtrl.ErrUserAlreadyExists:
		return errUserAlreadyExists
	default:
		return errInternalServer
	}
}
