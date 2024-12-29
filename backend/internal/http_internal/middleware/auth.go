package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/OsqY/GoingNext/internal/http_internal/handlers"
)

func JWTMiddleware(secretKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			publicRoutes := []string{
				"/login", "/register",
			}

			currentPath := r.URL.Path

			for _, route := range publicRoutes {
				if route == currentPath {
					next.ServeHTTP(w, r)
					return
				}
			}

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "No authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			if claims, err := handlers.VerifyToken(tokenString); err != nil {
				http.Error(w, "Invalid JWT", http.StatusUnauthorized)
				return
			} else {
				r = r.WithContext(context.WithValue(r.Context(), "user", claims))
				next.ServeHTTP(w, r)

			}
		})
	}
}
