package v1

import (
	"encoding/json"
	"errors"
	"github.com/tanmaij/friend-management/internal/controller/user"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
	stringUtil "github.com/tanmaij/friend-management/pkg/utils/string"
	"net/http"
)

type createUserReq struct {
	Email string `json:"email"`
}

type createUserRes struct {
	Success bool `json:"success"`
}

func (req createUserReq) validate() error {
	email := req.Email

	if !stringUtil.IsEmailValid(email) {
		return errInvalidGivenEmail
	}

	return nil
}

// CreateUser handles request create user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var reqData createUserReq
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

	if err := h.userCtrl.Create(r.Context(), user.CreateInput{Email: reqData.Email}); err != nil {
		converted := convertErrorFromController(err)
		httpUtil.WriteErrorToHttpResponseWriter(w, converted)
		return
	}

	httpUtil.WriteJsonData(w, http.StatusOK, createUserRes{Success: true})
}
