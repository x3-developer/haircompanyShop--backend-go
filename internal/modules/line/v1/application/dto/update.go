package dto

type UpdateDTO struct {
	Name  string `json:"name" validate:"min=3,max=64"`
	Color string `json:"color"`
}
