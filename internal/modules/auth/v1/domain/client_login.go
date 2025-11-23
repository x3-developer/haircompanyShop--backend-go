package domain

import "serv_shop_haircompany/internal/modules/auth/v1/domain/valueobject"

type ClientLogin struct {
	Phone    valueobject.Phone `json:"phone" validate:"required"`
	Password string            `json:"password" validate:"required"`
}
