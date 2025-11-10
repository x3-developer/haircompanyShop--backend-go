package domain

import (
	"serv_shop_haircompany/internal/modules/line/v1/domain/valueobject"
	"serv_shop_haircompany/internal/shared/domain"
)

type Line struct {
	ID    uint                `json:"id"`
	Name  string              `json:"name"`
	Color valueobject.ColorVO `json:"color"`
	domain.Timestamps
}
