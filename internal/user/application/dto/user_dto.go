package dto

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty" validate:"min=3"`
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"min=6"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
