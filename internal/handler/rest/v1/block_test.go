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

func TestHandler_Block(t *testing.T) {
	type mockBlock struct {
		isCalled  bool
		input     relationshipCtrl.BlockInput
		outputErr error
	}

	type output struct {
		statusCode int
		resBody    blockResponse
		resErr     error
	}

	tcs := map[string]struct {
		inputBody string
		mockBlock mockBlock
		expOutput output
	}{
		"success": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user2@example.com"
				}`,
			mockBlock: mockBlock{
				isCalled: true,
				input: relationshipCtrl.BlockInput{
					RequestorEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: nil,
			},
			expOutput: output{
				statusCode: http.StatusOK,
				resBody: blockResponse{
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
		"cannot_self_block": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user1@example.com"
				}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errCannotSelfBlock,
			},
		},
		"requestor_not_found": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user2@example.com"
				}`,
			mockBlock: mockBlock{
				isCalled: true,
				input: relationshipCtrl.BlockInput{
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
		"internal_server_error": {
			inputBody: `{
					"requestor": "user1@example.com",
					"target": "user2@example.com"
				}`,
			mockBlock: mockBlock{
				isCalled: true,
				input: relationshipCtrl.BlockInput{
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
			r := httptest.NewRequest(http.MethodPost, "/api/v1/relationship/block", strings.NewReader(tc.inputBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctrlMock := new(relationshipCtrl.MockController)
			if tc.mockBlock.isCalled {
				ctrlMock.EXPECT().Block(r.Context(), tc.mockBlock.input).Return(tc.mockBlock.outputErr)
			}

			h := New(ctrlMock, nil)

			// WHEN
			h.Block(w, r)

			// THEN
			assert.Equal(t, tc.expOutput.statusCode, w.Code)

			if tc.expOutput.resErr != nil {
				var result httpUtil.Error
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.EqualError(t, result, tc.expOutput.resErr.Error())
			} else {
				var result blockResponse
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expOutput.resBody)
			}

			ctrlMock.AssertExpectations(t)
		})
	}
}
