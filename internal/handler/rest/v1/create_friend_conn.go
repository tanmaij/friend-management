package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tanmaij/friend-management/internal/controller/relationship"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
	stringUtil "github.com/tanmaij/friend-management/pkg/utils/string"
)

type createFriendConnRequest struct {
	Friends []string `json:"friends"`
}

type createFriendConnResponse struct {
	Success bool `json:"success"`
}

func (req createFriendConnRequest) validate() error {
	if len(req.Friends) < 2 {
		return errInvalidRequestBody
	}

	if !stringUtil.IsEmailValid(strings.TrimSpace(req.Friends[0])) {
		return errInvalidGivenEmail
	}

	if !stringUtil.IsEmailValid(strings.TrimSpace(req.Friends[1])) {
		return errInvalidGivenEmail
	}

	return nil
}

// CreateFriendConn handles request create friend connection between 2 email addresses
func (h Handler) CreateFriendConn(w http.ResponseWriter, r *http.Request) {
	var reqData createFriendConnRequest
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

	if err := h.relationshipCtrl.CreateFriendConn(r.Context(), relationship.CreateFriendConnInp{
		RequesterEmail: strings.TrimSpace(reqData.Friends[0]),
		TargetEmail:    strings.TrimSpace(reqData.Friends[1]),
	}); err != nil {
		converted := convertErrorFromController(err)
		httpUtil.WriteErrorToHttpResponseWriter(w, converted)
		return
	}

	httpUtil.WriteJsonData(w, http.StatusOK, createFriendConnResponse{Success: true})
}
