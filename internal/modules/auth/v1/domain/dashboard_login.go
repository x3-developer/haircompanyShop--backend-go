package domain

import "serv_shop_haircompany/internal/modules/auth/v1/domain/valueobject"

type DashboardLogin struct {
	Email    valueobject.Email `json:"email" validate:"required"`
	Password string            `json:"password" validate:"required"`
}
