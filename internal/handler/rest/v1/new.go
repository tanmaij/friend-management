package v1

import "github.com/tanmaij/friend-management/internal/controller/relationship"

// Handler provides REST Request handle functions
type Handler struct {
	relationshipCtrl relationship.Controller
}

// New returns a new Handler instance with the given Controllers
func New(relationshipCtrl relationship.Controller) Handler {
	return Handler{relationshipCtrl: relationshipCtrl}
}
