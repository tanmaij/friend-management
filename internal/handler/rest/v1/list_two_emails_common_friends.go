package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/tanmaij/friend-management/internal/controller/relationship"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
	stringUtil "github.com/tanmaij/friend-management/pkg/utils/string"
)

type listTwoEmailCommonFriendsRequest struct {
	Friends []string `json:"friends"`
}

type listTwoEmailCommonFriendsResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int64    `json:"count"`
}

func (req listTwoEmailCommonFriendsRequest) validate() error {
	if len(req.Friends) < 2 {
		return errInvalidRequestBody
	}

	primaryEmail := strings.TrimSpace(req.Friends[0])
	secondaryEmail := strings.TrimSpace(req.Friends[1])

	if !stringUtil.IsEmailValid(primaryEmail) {
		return errInvalidGivenEmail
	}

	if !stringUtil.IsEmailValid(secondaryEmail) {
		return errInvalidGivenEmail
	}

	if primaryEmail == secondaryEmail {
		return errCannotGetCommonFriendWithSelf
	}

	return nil
}

func convertToListCommonFriendsResFromCtrl(ctrlOutput relationship.ListTwoEmailCommonFriendsOutput) listTwoEmailCommonFriendsResponse {
	friendEmails := make([]string, len(ctrlOutput.Friends))
	for i, friend := range ctrlOutput.Friends {
		friendEmails[i] = friend.Email
	}

	return listTwoEmailCommonFriendsResponse{Success: true, Friends: friendEmails, Count: ctrlOutput.Count}
}

// ListTwoEmailsCommonFriends handles the request to list common friends between two provided email addresses
func (h Handler) ListTwoEmailsCommonFriends(w http.ResponseWriter, r *http.Request) {
	var reqData listTwoEmailCommonFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		httpUtil.WriteErrorToHttpResponseWriter(w, errInvalidRequestBody)
		return
	}

	if err := reqData.validate(); err != nil {
		var expectedErr httpUtil.Error
		if errors.As(err, &expectedErr) {
			httpUtil.WriteErrorToHttpResponseWriter(w, expectedErr)
			return
		}

		httpUtil.WriteErrorToHttpResponseWriter(w, errInvalidRequestBody)
		return
	}

	result, err := h.relationshipCtrl.ListTwoEmailCommonFriends(r.Context(), relationship.ListTwoEmailCommonFriendsInput{
		PrimaryEmail:   strings.TrimSpace(reqData.Friends[0]),
		SecondaryEmail: strings.TrimSpace(reqData.Friends[1]),
	})
	if err != nil {
		converted := convertErrorFromController(err)
		httpUtil.WriteErrorToHttpResponseWriter(w, converted)
		return
	}

	response := convertToListCommonFriendsResFromCtrl(result)
	httpUtil.WriteJsonData(w, http.StatusOK, response)
}
