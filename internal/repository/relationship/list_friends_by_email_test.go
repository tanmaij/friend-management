package relationship

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/tanmaij/friend-management/internal/model"
	"github.com/tanmaij/friend-management/pkg/db/sql"
	testUtil "github.com/tanmaij/friend-management/pkg/utils/test"

	"github.com/stretchr/testify/assert"
)

func Test_impl_ListFriendByEmail(t *testing.T) {
	tcs := map[string]struct {
		inputEmail       string
		givenSqlFilePath string
		expOutputFriends []model.User
		expOutputCount   int64
		expOutputErr     bool
	}{
		"success": {
			inputEmail:       "user1@example.com",
			givenSqlFilePath: "test_data/list_friends_by_email.sql",
			expOutputCount:   2,
			expOutputErr:     false,
			expOutputFriends: []model.User{
				{ID: 2, Email: "user2@example.com"},
				{ID: 3, Email: "user3@example.com"},
			},
		},
		"no_friends_found": {
			inputEmail:       "user4@example.com",
			givenSqlFilePath: "test_data/list_friends_by_email.sql",
			expOutputCount:   0,
			expOutputErr:     false,
			expOutputFriends: []model.User{},
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
			defer dbTest.Exec(fmt.Sprintf(`TRUNCATE "%s" CASCADE; TRUNCATE "%s" CASCADE;`, model.TableNames.Relationships, model.TableNames.Users))

			repo := New(dbTest)

			// WHEN
			actRs, actCount, actErr := repo.ListFriendByEmail(emptyContext, tc.inputEmail)

			// THEN
			if tc.expOutputErr {
				assert.Error(t, actErr)
			} else {
				assert.NoError(t, actErr)
				assert.Equal(t, tc.expOutputCount, actCount)
				assert.Equal(t, len(tc.expOutputFriends), len(actRs))

				for i := range actRs {
					assert.Equal(t, tc.expOutputFriends[i].ID, actRs[i].ID)
					assert.Equal(t, tc.expOutputFriends[i].Email, actRs[i].Email)
				}
			}
		})
	}
}
