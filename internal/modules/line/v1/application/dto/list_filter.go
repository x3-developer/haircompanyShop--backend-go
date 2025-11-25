package dto

type ListFilter struct {
	Page    int `json:"page" form:"page" validate:"min=1"`
	PerPage int `json:"perPage" form:"perPage" validate:"min=1,max=100"`
}
