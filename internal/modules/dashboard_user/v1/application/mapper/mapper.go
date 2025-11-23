package mapper

import (
	"serv_shop_haircompany/internal/modules/dashboard_user/v1/application/dto"
	"serv_shop_haircompany/internal/modules/dashboard_user/v1/domain"
	"serv_shop_haircompany/internal/modules/dashboard_user/v1/domain/valueobject"
	sharedvalueobject "serv_shop_haircompany/internal/shared/domain/valueobject"
	"serv_shop_haircompany/internal/shared/utils/response"
)

func ToRespDTOFromModel(model domain.DashboardUser) *dto.ResponseDTO {
	return &dto.ResponseDTO{
		Id:        model.ID,
		Email:     model.Email.String(),
		Role:      string(model.Role),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ToModelFromCreateDTO(createDto dto.CreateDTO) (*domain.DashboardUser, []response.ErrorField) {
	email, err := valueobject.NewEmail(createDto.Email)
	if err != nil {
		return nil, []response.ErrorField{
			response.NewErrorField("email", string(response.BadRequest)),
		}
	}

	pass, err := valueobject.NewPasswordHash(createDto.Password)
	if err != nil {
		return nil, []response.ErrorField{
			response.NewErrorField("password", string(response.BadRequest)),
		}
	}

	role, err := sharedvalueobject.NewDashboardRole(createDto.Role)
	if err != nil {
		return nil, []response.ErrorField{
			response.NewErrorField("role", string(response.BadRequest)),
		}
	}

	return &domain.DashboardUser{
		Email:    email,
		Password: pass,
		Role:     role,
	}, nil
}
