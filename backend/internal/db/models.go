// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Role struct {
	ID          int32
	Name        string
	Description pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	DeletedAt   pgtype.Timestamptz
	CreatedBy   pgtype.Int4
	UpdatedBy   pgtype.Int4
	DeletedBy   pgtype.Int4
}

type User struct {
	ID        int32
	Email     string
	Username  string
	Password  string
	RoleID    int32
	IsActive  pgtype.Bool
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt pgtype.Timestamptz
	CreatedBy pgtype.Int4
	UpdatedBy pgtype.Int4
	DeletedBy pgtype.Int4
}
