package relationship

import (
	"context"
	"fmt"
	"strings"

	"github.com/tanmaij/friend-management/internal/model"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

// FindEligibleRecipientEmailsWithMentioned retrieves a list of eligible recipient emails
// based on the specified sender's email and a list of mentioned emails.
func (i *impl) FindEligibleRecipientEmailsWithMentioned(ctx context.Context, senderEmail string, mentionedEmails []string) ([]string, error) {
	transformedEmails := make([]string, len(mentionedEmails))
	for i, email := range mentionedEmails {
		transformedEmails[i] = "'" + email + "'"
	}

	query := fmt.Sprintf(`
        WITH combined_users AS ( 
			SELECT %s, %s  FROM %s 
			UNION ALL
			SELECT NULL, email FROM (SELECT unnest(ARRAY[%s]::varchar[]) AS email) AS emails
        )
        SELECT combined_users.email
        FROM combined_users
        JOIN %s AS sender ON sender.%s = $1
        LEFT JOIN %s ON 
			(%s.%s = '%s' AND %s.%s = combined_users.id AND %s.%s = sender.%s)
			OR
			(%s.%s = '%s' AND sender.%s IN (%s.%s, %s.%s) AND combined_users.%s IN (%s.%s, %s.%s))
			OR
			(%s.%s = '%s' AND %s.%s = combined_users.id AND %s.%s = sender.%s)
        WHERE combined_users.email <> $1
            AND (combined_users.id IS NULL OR (%s.%s IS NOT NULL))
        GROUP BY combined_users.email
		HAVING SUM(CASE WHEN %s.%s = '%s' THEN 1 ELSE 0 END) = 0;`,
		model.UserColumns.ID, model.UserColumns.Email, model.TableNames.Users,
		strings.Join(transformedEmails, ","),
		model.TableNames.Users, model.UserColumns.Email,
		model.TableNames.Relationships,
		model.TableNames.Relationships, model.RelationshipColumns.Type, model.RelationshipTypeSubscribe,
		model.TableNames.Relationships, model.RelationshipColumns.RequesterID,
		model.TableNames.Relationships, model.RelationshipColumns.TargetID, model.UserColumns.ID,
		model.TableNames.Relationships, model.RelationshipColumns.Type, model.RelationshipTypeFriend,
		model.UserColumns.ID, model.TableNames.Relationships, model.RelationshipColumns.TargetID, model.TableNames.Relationships, model.RelationshipColumns.RequesterID,
		model.UserColumns.ID, model.TableNames.Relationships, model.RelationshipColumns.TargetID, model.TableNames.Relationships, model.RelationshipColumns.RequesterID,
		model.TableNames.Relationships, model.RelationshipColumns.Type, model.RelationshipTypeBlock,
		model.TableNames.Relationships, model.RelationshipColumns.RequesterID,
		model.TableNames.Relationships, model.RelationshipColumns.TargetID, model.UserColumns.ID,
		model.TableNames.Relationships, model.RelationshipColumns.ID,
		model.TableNames.Relationships, model.RelationshipColumns.Type, model.RelationshipTypeBlock,
	)

	var bindded []struct {
		Email string `boil:"email"`
	}
	if err := queries.Raw(query, senderEmail).Bind(ctx, i.sqlDB, &bindded); err != nil {
		return nil, errors.WithStack(err)
	}

	var rs = make([]string, len(bindded))
	for i, v := range bindded {
		rs[i] = v.Email
	}

	return rs, nil
}
