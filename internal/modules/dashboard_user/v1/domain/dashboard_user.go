package domain

import (
	"serv_shop_haircompany/internal/modules/dashboard_user/v1/domain/valueobject"
	"serv_shop_haircompany/internal/shared/domain"
	sharedvalueobject "serv_shop_haircompany/internal/shared/domain/valueobject"
)

type DashboardUser struct {
	ID       uint                            `json:"id"`
	Email    valueobject.Email               `json:"email"`
	Password valueobject.PasswordHash        `json:"-"`
	Role     sharedvalueobject.DashboardRole `json:"role"`
	domain.Timestamps
}
