package http_internal

import (
	"github.com/OsqY/GoingNext/backend/internal/http_internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	userHandler *handlers.UserHandler
}

func NewRouter(userHandler *handlers.UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", userHandler.GetUserById)
		})
	})
	return r
}
