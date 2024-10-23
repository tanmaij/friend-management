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

func TestHandler_CreateFriendConn(t *testing.T) {
	type mockCreateFriendConn struct {
		isCalled  bool
		input     relationshipCtrl.CreateFriendConnInp
		outputErr error
	}

	type output struct {
		statusCode int
		resBody    createFriendConnResponse
		resErr     error
	}

	tcs := map[string]struct {
		inputBody            string
		mockCreateFriendConn mockCreateFriendConn
		expOutput            output
	}{
		"success": {
			inputBody: `{
					"friends":[
					    "user1@example.com",
						"user2@example.com"
			  		]
		          }`,
			mockCreateFriendConn: mockCreateFriendConn{
				isCalled: true,
				input: relationshipCtrl.CreateFriendConnInp{
					RequesterEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: nil,
			},
			expOutput: output{
				statusCode: http.StatusOK,
				resBody: createFriendConnResponse{
					Success: true,
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
		"invalid_request_body_with_invalid_friend_array": {
			inputBody: `{
					"friends":[
					    "user1@example.com"
			  		]
		          }`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidRequestBody,
			},
		},
		"invalid_email_payload": {
			inputBody: `{
					"friends":[
					    "user1@example.com",
						"user1example.com"
			  		]
		          }`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidGivenEmail,
			},
		},
		"user_not_found_with_given_email": {
			inputBody: `{
					"friends":[
					    "user1@example.com",
						"user2@example.com"
			  		]
		          }`,
			mockCreateFriendConn: mockCreateFriendConn{
				isCalled: true,
				input: relationshipCtrl.CreateFriendConnInp{
					RequesterEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: relationshipCtrl.ErrUserNotFoundWithGivenEmail,
			},
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errUserNotFoundWithGivenEmail,
			},
		},
		"can_not_be_friends_due_to_already_friends": {
			inputBody: `{
					"friends":[
					    "user1@example.com",
						"user2@example.com"
			  		]
		          }`,
			mockCreateFriendConn: mockCreateFriendConn{
				isCalled: true,
				input: relationshipCtrl.CreateFriendConnInp{
					RequesterEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: relationshipCtrl.ErrAlreadyFriends,
			},
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errAlreadyFriends,
			},
		},
		"can_not_be_friends_due_to_already_blocked": {
			inputBody: `{
					"friends":[
					    "user1@example.com",
						"user2@example.com"
			  		]
		          }`,
			mockCreateFriendConn: mockCreateFriendConn{
				isCalled: true,
				input: relationshipCtrl.CreateFriendConnInp{
					RequesterEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: relationshipCtrl.ErrAlreadyBlocked,
			},
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errAlreadyBlocked,
			},
		},
		"internal_server_error": {
			inputBody: `{
					"friends":[
					    "user1@example.com",
						"user2@example.com"
			  		]
		          }`,
			mockCreateFriendConn: mockCreateFriendConn{
				isCalled: true,
				input: relationshipCtrl.CreateFriendConnInp{
					RequesterEmail: "user1@example.com",
					TargetEmail:    "user2@example.com",
				},
				outputErr: errors.New("an unexpecting error"),
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
			r := httptest.NewRequest(http.MethodPost, "/api/v1/relationship/friend", strings.NewReader(tc.inputBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctrlMock := new(relationshipCtrl.MockController)
			if tc.mockCreateFriendConn.isCalled {
				ctrlMock.EXPECT().CreateFriendConn(r.Context(), tc.mockCreateFriendConn.input).Return(tc.mockCreateFriendConn.outputErr)
			}

			h := New(ctrlMock)

			// WHEN
			h.CreateFriendConn(w, r)

			// THEN
			assert.Equal(t, tc.expOutput.statusCode, w.Code)

			if tc.expOutput.resErr != nil {
				var result httpUtil.Error
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.EqualError(t, result, tc.expOutput.resErr.Error())
			} else {
				var result createFriendConnResponse
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expOutput.resBody)
			}

			ctrlMock.AssertExpectations(t)
		})
	}
}
