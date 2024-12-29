package http_internal

import (
	"github.com/OsqY/GoingNext/internal/http_internal/handlers"
	authMiddleware "github.com/OsqY/GoingNext/internal/http_internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	userHandler *handlers.UserHandler
	authHandler *handlers.AuthHandler
}

func NewRouter(userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Post("/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.JWTMiddleware(""))
		r.Route("/api/", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/{id}", userHandler.GetUserById)
				r.Post("/create", userHandler.CreateUser)
				r.Delete("/delete/{id}", userHandler.DeleteUser)
			})
		})
	})
	return r
}
