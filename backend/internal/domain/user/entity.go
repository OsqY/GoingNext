package user

import "time"

type User struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	RoleID    int64      `json:"role_id"`
	RoleName  string     `json:"role_name"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	CreatedBy *int64     `json:"created_by,omitempty"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
	DeletedBy *int64     `json:"deleted_by,omitempty"`
	ImageUrl  string     `json:"imageUrl"`
}

type Role struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	CreatedBy   *int64     `json:"created_by,omitempty"`
	UpdatedBy   *int64     `json:"updated_by,omitempty"`
	DeletedBy   *int64     `json:"deleted_by,omitempty"`
}
