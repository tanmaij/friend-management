package relationship

import (
	"context"
	"log"

	stringUtil "github.com/tanmaij/friend-management/pkg/utils/string"
)

// ListEligibleRecipientEmailsFromUpdateInput represents the input required to list recipients from update
type ListEligibleRecipientEmailsFromUpdateInput struct {
	SenderEmail string
	Text        string
}

// ListEligibleRecipientEmailsFromUpdateOutput represents the output containing the list of recipient emails
type ListEligibleRecipientEmailsFromUpdateOutput struct {
	Recipients []string
}

// ListEligibleRecipientEmailsFromUpdate
func (i *impl) ListEligibleRecipientEmailsFromUpdate(ctx context.Context, inp ListEligibleRecipientEmailsFromUpdateInput) (ListEligibleRecipientEmailsFromUpdateOutput, error) {
	extractedEmails := stringUtil.ExtractEmailsFromText(inp.Text)

	found, err := i.relationshipRepo.FindEligibleRecipientEmailsWithMentioned(ctx, inp.SenderEmail, extractedEmails)
	if err != nil {
		log.Printf("error listing recipients: %v", err)
		return ListEligibleRecipientEmailsFromUpdateOutput{}, err
	}

	return ListEligibleRecipientEmailsFromUpdateOutput{
		Recipients: found,
	}, nil
}
