package dto

type RefreshDTO struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
