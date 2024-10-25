package handler

import (
	"github.com/tanmaij/friend-management/internal/controller/relationship"
	"github.com/tanmaij/friend-management/internal/controller/user"
	restV1 "github.com/tanmaij/friend-management/internal/handler/rest/v1"
)

// Handler serves as the main entry point for managing different API types, API versions
type Handler struct {
	RESTV1Handler restV1.Handler
}

// New initializes a new Handler with the provided controllers.
func New(relCtrl relationship.Controller, userCtrl user.Controller) Handler {
	return Handler{
		RESTV1Handler: restV1.New(relCtrl, userCtrl),
	}
}
