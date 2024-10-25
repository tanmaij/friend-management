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

type subscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type subscribeResponse struct {
	Success bool `json:"success"`
}

func (req subscribeRequest) validate() error {
	requestorEmail := strings.TrimSpace(req.Requestor)
	targetEmail := strings.TrimSpace(req.Target)

	if requestorEmail == "" {
		return errRequestorEmailIsRequired
	}

	if targetEmail == "" {
		return errTargetEmailIsRequired
	}

	if !stringUtil.IsEmailValid(requestorEmail) {
		return errInvalidRequestorEmail
	}

	if !stringUtil.IsEmailValid(targetEmail) {
		return errInvalidTargetEmail
	}

	if requestorEmail == targetEmail {
		return errCannotSelfSubscribe
	}

	return nil
}

// Subscribe handles request subscribe an email address for updates
func (h Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var reqData subscribeRequest
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

	if err := h.relationshipCtrl.Subscribe(r.Context(), relationship.SubscribeInput{
		RequestorEmail: strings.TrimSpace(reqData.Requestor),
		TargetEmail:    strings.TrimSpace(reqData.Target),
	}); err != nil {
		converted := convertErrorFromController(err)
		httpUtil.WriteErrorToHttpResponseWriter(w, converted)
		return
	}

	httpUtil.WriteJsonData(w, http.StatusOK, subscribeResponse{Success: true})
}
