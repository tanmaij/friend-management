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

type listEligibleRecipientEmailsFromUpdateReq struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type listEligibleRecipientEmailsFromUpdateRes struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

func (req listEligibleRecipientEmailsFromUpdateReq) validate() error {
	var senderEmail = strings.TrimSpace(req.Sender)

	if senderEmail == "" {
		return errSenderEmailIsRequired
	}

	if !stringUtil.IsEmailValid(senderEmail) {
		return errInvalidGivenEmail
	}

	return nil
}

func convertToListRecipientsFromCtrl(ctrlOutput relationship.ListEligibleRecipientEmailsFromUpdateOutput) listEligibleRecipientEmailsFromUpdateRes {
	return listEligibleRecipientEmailsFromUpdateRes{Success: true, Recipients: ctrlOutput.Recipients}
}

// ListEligibleRecipientEmailsFromUpdate handles requests to list eligible recipient emails from an update.
func (h Handler) ListEligibleRecipientEmailsFromUpdate(w http.ResponseWriter, r *http.Request) {
	var reqData listEligibleRecipientEmailsFromUpdateReq
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

	result, err := h.relationshipCtrl.ListEligibleRecipientEmailsFromUpdate(
		r.Context(),
		relationship.ListEligibleRecipientEmailsFromUpdateInput{
			SenderEmail: strings.TrimSpace(reqData.Sender),
			Text:        strings.TrimSpace(reqData.Text),
		})
	if err != nil {
		converted := convertErrorFromController(err)
		httpUtil.WriteErrorToHttpResponseWriter(w, converted)
		return
	}

	response := convertToListRecipientsFromCtrl(result)
	httpUtil.WriteJsonData(w, http.StatusOK, response)
}
