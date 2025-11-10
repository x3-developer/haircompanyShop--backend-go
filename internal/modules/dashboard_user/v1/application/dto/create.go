package dto

type CreateDTO struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=64"`
	Role     string `json:"role" validate:"required"`
}
