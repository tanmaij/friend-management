package user

import (
	"context"
	"log"

	"github.com/tanmaij/friend-management/internal/model"
)

// CreateInput includes payload to creating user
type CreateInput struct {
	Email string
}

// Create handles the logic for creating user
func (i *impl) Create(ctx context.Context, inp CreateInput) error {
	existing, err := i.userRepo.ExistsByEmail(ctx, inp.Email)
	if err != nil {
		log.Printf("error checking user existing: %v", err)
		return err
	}

	if existing {
		return ErrUserAlreadyExists
	}

	if err := i.userRepo.Create(ctx, model.User{Email: inp.Email}); err != nil {
		log.Printf("error creating user: %v", err)
		return err
	}

	return nil
}
