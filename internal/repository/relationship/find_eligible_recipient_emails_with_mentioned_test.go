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

func Test_impl_FindEligibleRecipientEmailsWithMentioned(t *testing.T) {
	tcs := map[string]struct {
		senderEmail      string
		mentionedEmails  []string
		givenSqlFilePath string
		expOutputEmails  []string
		expOutputErr     bool
	}{
		"success_emails_found": {
			senderEmail: "user1@example.com",
			mentionedEmails: []string{
				"user2@example.com",
				"user10@example.com",
				"user11@example.com",
			},
			givenSqlFilePath: "test_data/find_eligible_recipient_emails_with_mentioned.sql",
			expOutputEmails: []string{
				"user10@example.com",
				"user11@example.com",
				"user3@example.com",
				"user4@example.com",
				"user5@example.com",
			},
			expOutputErr: false,
		},
		"no_emails_found": {
			senderEmail: "user6@example.com",
			mentionedEmails: []string{
				"user6@example.com",
			},
			givenSqlFilePath: "test_data/find_eligible_recipient_emails_with_mentioned.sql",
			expOutputEmails:  []string{},
			expOutputErr:     false,
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
				t.Fatal("failed to connect to database")
			}

			testUtil.LoadSqlFile(t, dbTest, tc.givenSqlFilePath)

			defer dbTest.Close()
			defer dbTest.Exec(fmt.Sprintf(`TRUNCATE "%s" CASCADE; TRUNCATE "%s" CASCADE;`, model.TableNames.Relationships, model.TableNames.Users))

			repo := New(dbTest)

			// WHEN
			actEmails, actErr := repo.FindEligibleRecipientEmailsWithMentioned(emptyContext, tc.senderEmail, tc.mentionedEmails)

			// THEN
			if tc.expOutputErr {
				assert.Error(t, actErr)
			} else {
				assert.NoError(t, actErr)
				assert.ElementsMatch(t, tc.expOutputEmails, actEmails)
			}
		})
	}
}
