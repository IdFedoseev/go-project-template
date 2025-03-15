package dto

type LoginDto struct {
	Email    string `json:"email" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=6"`
}
