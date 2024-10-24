package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	relationshipCtrl "github.com/tanmaij/friend-management/internal/controller/relationship"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
)

func TestHandler_ListEligibleRecipientEmailsFromUpdate(t *testing.T) {
	type mockListEligibleRecipientEmailsFromUpdate struct {
		isCalled  bool
		input     relationshipCtrl.ListEligibleRecipientEmailsFromUpdateInput
		output    relationshipCtrl.ListEligibleRecipientEmailsFromUpdateOutput
		outputErr error
	}

	type output struct {
		statusCode int
		resBody    listEligibleRecipientEmailsFromUpdateRes
		resErr     error
	}

	tcs := map[string]struct {
		inputBody                                 string
		mockListEligibleRecipientEmailsFromUpdate mockListEligibleRecipientEmailsFromUpdate
		expOutput                                 output
	}{
		"success": {
			inputBody: `{
				"sender": "user@example.com",
				"text": "Hello everyone!"
			}`,
			mockListEligibleRecipientEmailsFromUpdate: mockListEligibleRecipientEmailsFromUpdate{
				isCalled: true,
				input: relationshipCtrl.ListEligibleRecipientEmailsFromUpdateInput{
					SenderEmail: "user@example.com",
					Text:        "Hello everyone!",
				},
				output: relationshipCtrl.ListEligibleRecipientEmailsFromUpdateOutput{
					Recipients: []string{"recipient1@example.com", "recipient2@example.com"},
				},
				outputErr: nil,
			},
			expOutput: output{
				statusCode: http.StatusOK,
				resBody: listEligibleRecipientEmailsFromUpdateRes{
					Success:    true,
					Recipients: []string{"recipient1@example.com", "recipient2@example.com"},
				},
			},
		},
		"invalid_request_body_with_invalid_json": {
			inputBody: `{"sender": "user@example.com", "text": "Hello"`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidRequestBody,
			},
		},
		"invalid_request_body_with_empty_sender": {
			inputBody: `{
				"sender": "",
				"text": "Hello!"
			}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errSenderEmailIsRequired,
			},
		},
		"invalid_request_body_with_invalid_sender_email": {
			inputBody: `{
				"sender": "invalid-email",
				"text": "Hello!"
			}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidGivenEmail,
			},
		},
		"internal_server_error": {
			inputBody: `{
				"sender": "user@example.com",
				"text": "Hello everyone!"
			}`,
			mockListEligibleRecipientEmailsFromUpdate: mockListEligibleRecipientEmailsFromUpdate{
				isCalled: true,
				input: relationshipCtrl.ListEligibleRecipientEmailsFromUpdateInput{
					SenderEmail: "user@example.com",
					Text:        "Hello everyone!",
				},
				outputErr: errors.New("an unexpected error"),
			},
			expOutput: output{
				statusCode: http.StatusInternalServerError,
				resErr:     errInternalServer,
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			r := httptest.NewRequest(http.MethodPost, "/api/v1/relationship/eligible/recipients/update", strings.NewReader(tc.inputBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctrlMock := new(relationshipCtrl.MockController)
			if tc.mockListEligibleRecipientEmailsFromUpdate.isCalled {
				ctrlMock.EXPECT().ListEligibleRecipientEmailsFromUpdate(r.Context(), tc.mockListEligibleRecipientEmailsFromUpdate.input).Return(tc.mockListEligibleRecipientEmailsFromUpdate.output, tc.mockListEligibleRecipientEmailsFromUpdate.outputErr)
			}

			h := New(ctrlMock)

			// WHEN
			h.ListEligibleRecipientEmailsFromUpdate(w, r)

			// THEN
			assert.Equal(t, tc.expOutput.statusCode, w.Code)

			if tc.expOutput.resErr != nil {
				var result httpUtil.Error
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.EqualError(t, result, tc.expOutput.resErr.Error())
			} else {
				var result listEligibleRecipientEmailsFromUpdateRes
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expOutput.resBody)
			}

			ctrlMock.AssertExpectations(t)
		})
	}
}
