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
	"github.com/tanmaij/friend-management/internal/model"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
)

func TestHandler_ListFriendByEmail(t *testing.T) {
	type mockListFriendByEmail struct {
		isCalled  bool
		input     relationshipCtrl.ListFriendByEmailInput
		output    relationshipCtrl.ListFriendByEmailOutput
		outputErr error
	}

	type output struct {
		statusCode int
		resBody    listFriendsResponse
		resErr     error
	}

	tcs := map[string]struct {
		inputBody             string
		mockListFriendByEmail mockListFriendByEmail
		expOutput             output
	}{
		"success": {
			inputBody: `{
				"email": "user@example.com"
			}`,
			mockListFriendByEmail: mockListFriendByEmail{
				isCalled: true,
				input: relationshipCtrl.ListFriendByEmailInput{
					Email: "user@example.com",
				},
				output: relationshipCtrl.ListFriendByEmailOutput{
					Friends: []model.User{
						{Email: "friend1@example.com"},
						{Email: "friend2@example.com"},
					},
					Count: 2,
				},
				outputErr: nil,
			},
			expOutput: output{
				statusCode: http.StatusOK,
				resBody: listFriendsResponse{
					Success: true,
					Friends: []string{"friend1@example.com", "friend2@example.com"},
					Count:   2,
				},
			},
		},
		"invalid_request_body_with_invalid_json_body": {
			inputBody: `body}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidRequestBody,
			},
		},
		"invalid_request_body_with_invalid_email": {
			inputBody: `{
				"email": "invalid-email"
			}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidGivenEmail,
			},
		},
		"internal_server_error": {
			inputBody: `{
				"email": "user@example.com"
			}`,
			mockListFriendByEmail: mockListFriendByEmail{
				isCalled: true,
				input: relationshipCtrl.ListFriendByEmailInput{
					Email: "user@example.com",
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
			r := httptest.NewRequest(http.MethodPost, "/api/v1/relationship/friend/list", strings.NewReader(tc.inputBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctrlMock := new(relationshipCtrl.MockController)
			if tc.mockListFriendByEmail.isCalled {
				ctrlMock.EXPECT().ListFriendByEmail(r.Context(), tc.mockListFriendByEmail.input).Return(tc.mockListFriendByEmail.output, tc.mockListFriendByEmail.outputErr)
			}

			h := New(ctrlMock, nil)

			// WHEN
			h.ListFriendByEmail(w, r)

			// THEN
			assert.Equal(t, tc.expOutput.statusCode, w.Code)

			if tc.expOutput.resErr != nil {
				var result httpUtil.Error
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.EqualError(t, result, tc.expOutput.resErr.Error())
			} else {
				var result listFriendsResponse
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expOutput.resBody)
			}

			ctrlMock.AssertExpectations(t)
		})
	}
}
