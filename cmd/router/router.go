package router

import (
	"net/http"

	"github.com/tanmaij/friend-management/internal/handler"
	v1 "github.com/tanmaij/friend-management/internal/handler/rest/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// InitRouter initializes routes
func InitRouter(handler handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	v1Route(r, handler.RESTV1Handler)
	return r
}

func v1Route(r chi.Router, h v1.Handler) {
	r.Route("/api/v1", func(apiV1Router chi.Router) {
		apiV1Router.Route("/relationship", func(relRouter chi.Router) {
			relRouter.Post("/friend", h.CreateFriendConn)
			relRouter.Post("/friend/list", h.ListFriendByEmail)
			relRouter.Post("/friend/list-common", h.ListTwoEmailsCommonFriends)
			relRouter.Post("/subscribe", h.Subscribe)
			relRouter.Post("/block", h.Block)
		})

		apiV1Router.Route("/update", func(relRouter chi.Router) {
			relRouter.Post("/recipient", h.ListEligibleRecipientEmailsFromUpdate)
		})
	})
}
