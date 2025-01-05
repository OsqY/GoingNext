package http_internal

import (
	"github.com/OsqY/GoingNext/internal/http_internal/handlers"
	authMiddleware "github.com/OsqY/GoingNext/internal/http_internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	userHandler *handlers.UserHandler
	authHandler *handlers.AuthHandler
	roleHandler *handlers.RoleHandler
	fileHandler *handlers.FileHandler
}

func NewRouter(userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler, roleHandler *handlers.RoleHandler, fileHandler *handlers.FileHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	r.Use(cors.Handler)

	r.Group(func(r chi.Router) {
		r.Post("/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.JWTMiddleware(""))
		r.Route("/api/", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/{id}", userHandler.GetUserById)
				r.Get("/current", authHandler.GetCurrentUser)
				r.Post("/create", userHandler.CreateUser)
				r.Put("/update/{id}", userHandler.UpdateUser)
				r.Delete("/delete/{id}", userHandler.DeleteUser)
			})
			r.Route("/roles", func(r chi.Router) {
				r.Get("/all", roleHandler.GetRoles)
			})
			r.Route("/files", func(r chi.Router) {
				r.Post("/sends3", fileHandler.SendFileToS3)
			})
		})
	})
	return r
}
