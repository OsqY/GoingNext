package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OsqY/GoingNext/internal/application/dto"
	"github.com/OsqY/GoingNext/internal/db"
	"github.com/OsqY/GoingNext/lib/hash"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserHandler struct {
	queries *db.Queries
}

func NewUserHandler(queries *db.Queries) *UserHandler {
	return &UserHandler{queries: queries}
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.queries.GetUserByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonStr)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validator := validator.New()

	if err := validator.Struct(req); err != nil {
		http.Error(w, "error validating your request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "there was an error hashing your password", http.StatusInternalServerError)
		return
	}

	user, err := h.queries.CreateUser(r.Context(), db.CreateUserParams{
		Username: req.Username, Email: req.Email, Password: hashedPassword, RoleID: int32(req.RoleID), ImageUrl: pgtype.Text{req.ImageURL, true},
	})
	if err != nil {
		http.Error(w, "there was an error saving your info: "+err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserRequest

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validator := validator.New()

	if err := validator.Struct(req); err != nil {
		http.Error(w, "error validating your request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.queries.UpdateUser(r.Context(), db.UpdateUserParams{
		ID:       int64(id),
		RoleID:   sql.NullInt64{Int64: int64(req.RoleID), Valid: true},
		Username: sql.NullString{String: req.Username, Valid: req.Username != ""},
		Email:    sql.NullString{String: req.Email, Valid: req.Email != ""},
		Password: sql.NullString{String: req.Password, Valid: req.Password != ""},
		ImageUrl: sql.NullString{String: req.ImageURL, Valid: req.ImageURL != ""},
	})
	if err != nil {
		http.Error(w, "there was an error saving your info: "+err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
	}

	claims, ok := r.Context().Value("user").(jwt.MapClaims)
	if !ok {
		http.Error(w, "invalid token claims", http.StatusUnauthorized)
		return
	}

	email := claims["email"].(string)

	user, err := h.queries.GetUserByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "user with that email doesn't exist", http.StatusUnauthorized)
		return
	}

	h.queries.SoftDeleteUser(r.Context(), db.SoftDeleteUserParams{ID: int32(id), DeletedBy: pgtype.Int4{Int32: user.ID}})
}
