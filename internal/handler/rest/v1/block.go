package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tanmaij/friend-management/internal/controller/relationship"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
	stringUtil "github.com/tanmaij/friend-management/pkg/utils/string"
)

type blockRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type blockResponse struct {
	Success bool `json:"success"`
}

func (req blockRequest) validate() error {
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
		return errCannotSelfBlock
	}

	return nil
}

// Block handles request block updates from an email address
func (h Handler) Block(w http.ResponseWriter, r *http.Request) {
	var reqData blockRequest
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

	if err := h.relationshipCtrl.Block(r.Context(), relationship.BlockInput{
		RequestorEmail: strings.TrimSpace(reqData.Requestor),
		TargetEmail:    strings.TrimSpace(reqData.Target),
	}); err != nil {
		converted := convertErrorFromController(err)
		httpUtil.WriteErrorToHttpResponseWriter(w, converted)
		return
	}

	httpUtil.WriteJsonData(w, http.StatusOK, blockResponse{Success: true})
}
