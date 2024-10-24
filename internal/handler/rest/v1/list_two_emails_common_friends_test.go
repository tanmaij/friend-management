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

func TestHandler_ListTwoEmailsCommonFriends(t *testing.T) {
	type mockListTwoEmailCommonFriends struct {
		isCalled  bool
		input     relationshipCtrl.ListTwoEmailCommonFriendsInput
		output    relationshipCtrl.ListTwoEmailCommonFriendsOutput
		outputErr error
	}

	type output struct {
		statusCode int
		resBody    listTwoEmailCommonFriendsResponse
		resErr     error
	}

	tcs := map[string]struct {
		inputBody                     string
		mockListTwoEmailCommonFriends mockListTwoEmailCommonFriends
		expOutput                     output
	}{
		"success": {
			inputBody: `{
				"friends": ["user1@example.com", "user2@example.com"]
			}`,
			mockListTwoEmailCommonFriends: mockListTwoEmailCommonFriends{
				isCalled: true,
				input: relationshipCtrl.ListTwoEmailCommonFriendsInput{
					PrimaryEmail:   "user1@example.com",
					SecondaryEmail: "user2@example.com",
				},
				output: relationshipCtrl.ListTwoEmailCommonFriendsOutput{
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
				resBody: listTwoEmailCommonFriendsResponse{
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
				"friends": ["invalid-email", "user2@example.com"]
			}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidGivenEmail,
			},
		},
		"same_email": {
			inputBody: `{
				"friends": ["user1@example.com", "user1@example.com"]
			}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errCannotGetCommonFriendWithSelf,
			},
		},
		"internal_server_error": {
			inputBody: `{
				"friends": ["user1@example.com", "user2@example.com"]
			}`,
			mockListTwoEmailCommonFriends: mockListTwoEmailCommonFriends{
				isCalled: true,
				input: relationshipCtrl.ListTwoEmailCommonFriendsInput{
					PrimaryEmail:   "user1@example.com",
					SecondaryEmail: "user2@example.com",
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
			r := httptest.NewRequest(http.MethodPost, "/api/v1/relationship/friend/list-common", strings.NewReader(tc.inputBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctrlMock := new(relationshipCtrl.MockController)
			if tc.mockListTwoEmailCommonFriends.isCalled {
				ctrlMock.EXPECT().ListTwoEmailCommonFriends(r.Context(), tc.mockListTwoEmailCommonFriends.input).Return(tc.mockListTwoEmailCommonFriends.output, tc.mockListTwoEmailCommonFriends.outputErr)
			}

			h := New(ctrlMock)

			// WHEN
			h.ListTwoEmailsCommonFriends(w, r)

			// THEN
			assert.Equal(t, tc.expOutput.statusCode, w.Code)

			if tc.expOutput.resErr != nil {
				var result httpUtil.Error
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.EqualError(t, result, tc.expOutput.resErr.Error())
			} else {
				var result listTwoEmailCommonFriendsResponse
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expOutput.resBody)
			}

			ctrlMock.AssertExpectations(t)
		})
	}
}
