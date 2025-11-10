package domain

import (
	"serv_shop_haircompany/internal/modules/dashboard_user/v1/domain/valueobject"
	"serv_shop_haircompany/internal/shared/domain"
)

type DashboardUser struct {
	ID       uint                     `json:"id"`
	Email    valueobject.Email        `json:"email"`
	Password valueobject.PasswordHash `json:"-"`
	Role     Role                     `json:"role"`
	domain.Timestamps
}
