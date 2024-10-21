package relationship

import "context"

// CreateFriendConnInp holds the input data for creating a friend connection, including the requester's email and the target's email.
type CreateFriendConnInp struct {
	RequesterEmail string // The email of the person requesting the friendship
	TargetEmail    string // The email of the person to be added as a friend
}

// CreateFriendConn handles the logic for creating a friend connection.
func (i *impl) CreateFriendConn(ctx context.Context, inp CreateFriendConnInp) error {
	// TODO: implement this later
	return nil
}
