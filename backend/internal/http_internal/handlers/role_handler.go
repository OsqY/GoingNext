package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/OsqY/GoingNext/internal/db"
)

type RoleHandler struct {
	queries *db.Queries
}

func NewRoleHandler(queries *db.Queries) *RoleHandler {
	return &RoleHandler{queries: queries}
}

func (ro *RoleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := ro.queries.ListRoles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}
