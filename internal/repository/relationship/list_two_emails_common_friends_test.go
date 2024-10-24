package relationship

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

func Test_impl_ListTwoEmailsCommonFriends(t *testing.T) {
	tcs := map[string]struct {
		primaryEmail     string
		secondaryEmail   string
		givenSqlFilePath string
		expOutputFriends []model.User
		expOutputCount   int64
		expOutputErr     bool
	}{
		"success_common_friends_found": {
			primaryEmail:     "user2@example.com",
			secondaryEmail:   "user3@example.com",
			givenSqlFilePath: "test_data/list_two_emails_common_friends.sql",
			expOutputCount:   3,
			expOutputErr:     false,
			expOutputFriends: []model.User{
				{ID: 1, Email: "user1@example.com"},
				{ID: 4, Email: "user4@example.com"},
				{ID: 5, Email: "user5@example.com"},
			},
		},
		"no_common_friends_found": {
			primaryEmail:     "user10@example.com",
			secondaryEmail:   "user5@example.com",
			givenSqlFilePath: "test_data/list_two_emails_common_friends.sql",
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
			actRs, actCount, actErr := repo.ListTwoEmailsCommonFriends(emptyContext, tc.primaryEmail, tc.secondaryEmail)

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
