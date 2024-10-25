package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanmaij/friend-management/internal/controller/user"
	httpUtil "github.com/tanmaij/friend-management/pkg/utils/http"
)

func TestHandler_CreateUser(t *testing.T) {
	type mockCreateUser struct {
		isCalled  bool
		input     user.CreateInput
		outputErr error
	}

	type output struct {
		statusCode int
		resBody    createUserRes
		resErr     error
	}

	tcs := map[string]struct {
		inputBody      string
		mockCreateUser mockCreateUser
		expOutput      output
	}{
		"success": {
			inputBody: `{"email":"user@example.com"}`,
			mockCreateUser: mockCreateUser{
				isCalled: true,
				input: user.CreateInput{
					Email: "user@example.com",
				},
				outputErr: nil,
			},
			expOutput: output{
				statusCode: http.StatusOK,
				resBody: createUserRes{
					Success: true,
				},
			},
		},
		"invalid_request_body": {
			inputBody: `invalid json`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidRequestBody,
			},
		},
		"invalid_email": {
			inputBody: `{"email":"invalid-email"}`,
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errInvalidGivenEmail,
			},
		},
		"error_user_already_exists": {
			inputBody: `{"email":"user@example.com"}`,
			mockCreateUser: mockCreateUser{
				isCalled: true,
				input: user.CreateInput{
					Email: "user@example.com",
				},
				outputErr: user.ErrUserAlreadyExists,
			},
			expOutput: output{
				statusCode: http.StatusBadRequest,
				resErr:     errUserAlreadyExists,
			},
		},
		"error_creating_user": {
			inputBody: `{"email":"user@example.com"}`,
			mockCreateUser: mockCreateUser{
				isCalled: true,
				input: user.CreateInput{
					Email: "user@example.com",
				},
				outputErr: errors.New("some error"),
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
			r := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(tc.inputBody))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctrlMock := new(user.MockController)
			if tc.mockCreateUser.isCalled {
				ctrlMock.EXPECT().Create(r.Context(), tc.mockCreateUser.input).Return(tc.mockCreateUser.outputErr)
			}

			h := New(nil, ctrlMock)

			// WHEN
			h.CreateUser(w, r)

			// THEN
			assert.Equal(t, tc.expOutput.statusCode, w.Code)

			if tc.expOutput.resErr != nil {
				var result httpUtil.Error
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.EqualError(t, result, tc.expOutput.resErr.Error())
			} else {
				var result createUserRes
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expOutput.resBody)
			}
		})
	}
}
