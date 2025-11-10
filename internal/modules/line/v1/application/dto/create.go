package dto

type CreateDTO struct {
	Name  string `json:"name" validate:"required,min=3,max=64"`
	Color string `json:"color" validate:"required"`
}
