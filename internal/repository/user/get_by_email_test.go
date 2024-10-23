package user

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanmaij/friend-management/internal/model"
	"github.com/tanmaij/friend-management/pkg/db/sql"
	testUtil "github.com/tanmaij/friend-management/pkg/utils/test"
)

func Test_impl_GetByEmail(t *testing.T) {
	tcs := map[string]struct {
		inputEmail       string
		givenSqlFilePath string
		expOutputUser    model.User
		expOutputErr     error
	}{
		"success": {
			inputEmail:       "user1@example.com",
			givenSqlFilePath: "test_data/get_by_email.sql",
			expOutputUser: model.User{
				ID:    1,
				Email: "user1@example.com",
			},
		},
		"user_not_found": {
			inputEmail:       "user3@example.com",
			givenSqlFilePath: "test_data/get_by_email.sql",
			expOutputErr:     ErrUserNotFound,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			emptyContext := context.Background()
			dbUrl := os.Getenv("DB_URL")
			if dbUrl == "" {
				t.Fatal("DB Url is empty")
			}

			dbTest, dbErr := sql.ConnectDB(sql.Postgres, dbUrl, sql.ConnectionOption{})
			if dbErr != nil {
				t.Fatal("failed to connecting database")
			}

			testUtil.LoadSqlFile(t, dbTest, tc.givenSqlFilePath)

			defer dbTest.Close()
			defer dbTest.Exec(fmt.Sprintf(`TRUNCATE "%s" CASCADE;`, model.TableNames.Users))

			repo := New(dbTest)

			// WHEN
			actRs, actErr := repo.GetByEmail(emptyContext, tc.inputEmail)

			// THEN
			if tc.expOutputErr != nil {
				assert.EqualError(t, actErr, tc.expOutputErr.Error())
			} else {
				assert.NoError(t, actErr)
				assert.Equal(t, tc.expOutputUser.ID, actRs.ID)
				assert.Equal(t, tc.expOutputUser.Email, actRs.Email)
			}
		})
	}
}
