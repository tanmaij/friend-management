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

func Test_impl_ListByTwoUserIDs(t *testing.T) {
	tcs := map[string]struct {
		inputPrimaryUserID   int
		inputSecondaryUserID int
		givenSqlFilePath     string
		expOutputRels        []model.Relationship
		expOutputErr         bool
	}{
		"success": {
			inputPrimaryUserID:   1,
			inputSecondaryUserID: 2,
			givenSqlFilePath:     "test_data/list_by_two_user_ids.sql",
			expOutputErr:         false,
			expOutputRels: []model.Relationship{
				{
					ID:          1,
					RequesterID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeFriend,
				},
				{
					ID:          2,
					RequesterID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeBlock,
				},
			},
		},
		"empty_result": {
			inputPrimaryUserID:   1,
			inputSecondaryUserID: 3,
			givenSqlFilePath:     "test_data/list_by_two_user_ids.sql",
			expOutputErr:         false,
			expOutputRels:        []model.Relationship{},
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
			actRs, actErr := repo.ListByTwoUserIDs(emptyContext, tc.inputPrimaryUserID, tc.inputSecondaryUserID)

			// THEN
			if tc.expOutputErr {
				assert.Error(t, actErr)
			} else {
				assert.NoError(t, actErr)
				assert.Equal(t, len(tc.expOutputRels), len(actRs))

				for i := range actRs {
					assert.Equal(t, tc.expOutputRels[i].ID, actRs[i].ID)
					assert.Equal(t, tc.expOutputRels[i].RequesterID, actRs[i].RequesterID)
					assert.Equal(t, tc.expOutputRels[i].TargetID, actRs[i].TargetID)
					assert.Equal(t, tc.expOutputRels[i].Type, actRs[i].Type)
				}
			}
		})
	}
}
