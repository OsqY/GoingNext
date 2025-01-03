package dto

type CreateUserRequest struct {
	Email    string  `json:"email" validate:"required,email"`
	Username string  `json:"username" validate:"required,min=3"`
	Password string  `json:"password" validate:"required,min=8"`
	RoleID   int64   `json:"role_id" validate:"required"`
	ImageURL *string `json:"image_url"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	Username *string `json:"username" validate:"omitempty,min=3"`
	Password string  `json:"password" validate:"required,min=8"`
	RoleID   *int64  `json:"role_id" validate:"omitempty"`
	ImageURL *string `json:"image_url"`
}
