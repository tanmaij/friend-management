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

func Test_impl_Create(t *testing.T) {
	tcs := map[string]struct {
		inputRelationship model.Relationship
		givenSqlFilePath  string
		expOutputErr      bool
	}{
		"success": {
			inputRelationship: model.Relationship{
				RequesterID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeSubscribe,
			},
			givenSqlFilePath: "test_data/create.sql",
			expOutputErr:     false,
		},
		"violate_user_relationship_idx": {
			inputRelationship: model.Relationship{
				RequesterID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeFriend,
			},
			givenSqlFilePath: "test_data/create.sql",
			expOutputErr:     true,
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
			actErr := repo.Create(emptyContext, tc.inputRelationship)

			// THEN
			if tc.expOutputErr {
				assert.Error(t, actErr)
			} else {
				assert.NoError(t, actErr)
			}
		})
	}
}
