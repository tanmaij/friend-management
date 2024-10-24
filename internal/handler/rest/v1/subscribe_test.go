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

func TestHandler_Subscribe(t *testing.T) {
	type mockSubscribe struct {
		isCalled  bool
		input     relationshipCtrl.SubscribeInput
		outputErr error
	}

	type output struct {
		statusCode int
		resBody    subscribeResponse
		resErr     error
	}

	tcs := map[string]struct {
		inputBody     string
		mockSubscribe mockSubscribe
		expOutput     output
	}{
		"success": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user2@example.com"
				}`,
			mockSubscribe: mockSubscribe{
				isCalled: true,
				input: relationshipCtrl.SubscribeInput{
					RequestorEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: nil,
			},
			expOutput: output{
				statusCode: http.StatusOK,
				resBody: subscribeResponse{
					Success: true,
				},
			},
		},
		"invalid_request_body": {
			inputBody: `body}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidRequestBody,
			},
		},
		"missing_requestor": {
			inputBody: `{
					"requestor": "",
					"target": "user2@example.com"
				}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errRequestorEmailIsRequired,
			},
		},
		"missing_target": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": ""
				}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errTargetEmailIsRequired,
			},
		},
		"invalid_requestor_email": {
			inputBody: `{
					"requestor": "invalid-email",
					"target": "user2@example.com"
				}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidRequestorEmail,
			},
		},
		"invalid_target_email": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "invalid-email"
				}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidTargetEmail,
			},
		},
		"requestor_not_found": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user2@example.com"
				}`,
			mockSubscribe: mockSubscribe{
				isCalled: true,
				input: relationshipCtrl.SubscribeInput{
					RequestorEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: relationshipCtrl.ErrUserNotFoundWithGivenEmail,
			},
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errUserNotFoundWithGivenEmail,
			},
		},
		"already_subscribed": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user2@example.com"
				}`,
			mockSubscribe: mockSubscribe{
				isCalled: true,
				input: relationshipCtrl.SubscribeInput{
					RequestorEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: relationshipCtrl.ErrAlreadySubscribed,
			},
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errAlreadySubscribed,
			},
		},
		"internal_server_error": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user2@example.com"
				}`,
			mockSubscribe: mockSubscribe{
				isCalled: true,
				input: relationshipCtrl.SubscribeInput{
					RequestorEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: errors.New("unexpected error"),
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
			r := httptest.NewRequest(http.MethodPost, "/api/v1/relationship/subscribe", strings.NewReader(tc.inputBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctrlMock := new(relationshipCtrl.MockController)
			if tc.mockSubscribe.isCalled {
				ctrlMock.EXPECT().Subscribe(r.Context(), tc.mockSubscribe.input).Return(tc.mockSubscribe.outputErr)
			}

			h := New(ctrlMock)

			// WHEN
			h.Subscribe(w, r)

			// THEN
			assert.Equal(t, tc.expOutput.statusCode, w.Code)

			if tc.expOutput.resErr != nil {
				var result httpUtil.Error
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.EqualError(t, result, tc.expOutput.resErr.Error())
			} else {
				var result subscribeResponse
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expOutput.resBody)
			}

			ctrlMock.AssertExpectations(t)
		})
	}
}
