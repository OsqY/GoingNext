package dto

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
	RoleID   int64  `json:"roleId" validate:"required"`
	ImageURL string `json:"imageURL"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username" validate:"omitempty,min=3"`
	Password string `json:"password" validate:"required,min=8"`
	ID       int    `json:"id" validate:"required"`
	RoleID   int64  `json:"roleId" validate:"required"`
	ImageURL string `json:"imageURL"`
}
