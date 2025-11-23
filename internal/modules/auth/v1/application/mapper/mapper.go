package mapper

import (
	"serv_shop_haircompany/internal/modules/auth/v1/application/dto"
	"serv_shop_haircompany/internal/modules/auth/v1/domain"
	"serv_shop_haircompany/internal/modules/auth/v1/domain/valueobject"
	"serv_shop_haircompany/internal/shared/infrastructure/security"
	"serv_shop_haircompany/internal/shared/utils/response"
)

func ToRespDTOFromModel(tokenPair security.TokenPair) *dto.ResponseDTO {
	return &dto.ResponseDTO{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}
}

func ToModelFromDashboardAuthDTO(dashboardAuth dto.DashboardAuthDTO) (*domain.DashboardLogin, []response.ErrorField) {
	email, err := valueobject.NewEmail(dashboardAuth.Email)
	if err != nil {
		return nil, []response.ErrorField{
			response.NewErrorField("email", string(response.BadRequest)),
		}
	}

	return &domain.DashboardLogin{
		Email:    email,
		Password: dashboardAuth.Password,
	}, nil
}
