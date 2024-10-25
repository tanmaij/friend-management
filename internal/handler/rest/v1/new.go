package v1

import (
	"github.com/tanmaij/friend-management/internal/controller/relationship"
	"github.com/tanmaij/friend-management/internal/controller/user"
)

// Handler provides REST Request handle functions
type Handler struct {
	relationshipCtrl relationship.Controller
	userCtrl         user.Controller
}

// New returns a new Handler instance with the given Controllers
func New(relationshipCtrl relationship.Controller, userCtrl user.Controller) Handler {
	return Handler{relationshipCtrl: relationshipCtrl, userCtrl: userCtrl}
}
