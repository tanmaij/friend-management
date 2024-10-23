package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tanmaij/friend-management/internal/controller/relationship"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
	stringUtil "github.com/tanmaij/friend-management/pkg/utils/string"
)

type listFriendsRequest struct {
	Email string `json:"email"`
}

type listFriendsResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int64    `json:"count"`
}

func (req listFriendsRequest) validate() error {
	reqEmail := strings.TrimSpace(req.Email)

	if !stringUtil.IsEmailValid(reqEmail) {
		return errInvalidGivenEmail
	}

	return nil
}

func convertToListFriendsResFromCtrl(ctrlOutput relationship.ListFriendByEmailOutput) listFriendsResponse {
	friendEmails := make([]string, len(ctrlOutput.Friends))
	for i, friend := range ctrlOutput.Friends {
		friendEmails[i] = friend.Email
	}

	return listFriendsResponse{Success: true, Friends: friendEmails, Count: ctrlOutput.Count}
}

// ListFriendByEmail handles request retrieving the friends list for an email address
func (h Handler) ListFriendByEmail(w http.ResponseWriter, r *http.Request) {
	var reqData listFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		httpUtil.WriteErrorToHttpResponseWriter(w, errInvalidRequestBody)
		return
	}

	if err := reqData.validate(); err != nil {
		if expectedErr, ok := err.(httpUtil.Error); ok {
			httpUtil.WriteErrorToHttpResponseWriter(w, expectedErr)
			return
		}

		httpUtil.WriteErrorToHttpResponseWriter(w, errInvalidRequestBody)
		return
	}

	result, err := h.relationshipCtrl.ListFriendByEmail(r.Context(), relationship.ListFriendByEmailInput{Email: reqData.Email})
	if err != nil {
		converted := convertErrorFromController(err)
		httpUtil.WriteErrorToHttpResponseWriter(w, converted)
		return
	}

	response := convertToListFriendsResFromCtrl(result)
	httpUtil.WriteJsonData(w, http.StatusOK, response)
}
