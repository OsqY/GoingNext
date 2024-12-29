package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/OsqY/GoingNext/internal/config"
	"github.com/OsqY/GoingNext/internal/db"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	queries *db.Queries
}

func NewAuthHandler(queries *db.Queries) *AuthHandler {
	return &AuthHandler{queries: queries}
}

func createToken(email string) (string, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	secretKey := config.Auth.SecretKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "exp": time.Now().Add(time.Hour * 24).Unix()})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	secretKey := config.Auth.SecretKey

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func GetUserIdFromToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(tkoen *jwt.Token) (interface{}, error) {
		return tokenString, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims["userId"].(int), nil
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user db.User
	json.NewDecoder(r.Body).Decode(&user)

	userInfo, err := a.queries.GetUserByEmail(r.Context(), user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Errorf("user not found")
	}

	if user.Email == userInfo.Email && user.Password == userInfo.Password {
		tokenString, err := createToken(user.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("invalid credentials")

		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid credentials")
	}
}
