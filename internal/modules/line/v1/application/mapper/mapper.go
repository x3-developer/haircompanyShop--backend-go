package mapper

import (
	"serv_shop_haircompany/internal/modules/line/v1/application/dto"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	"serv_shop_haircompany/internal/modules/line/v1/domain/valueobject"
	"serv_shop_haircompany/internal/shared/utils/response"
)

func ToRespDTOFromModel(model domain.Line) *dto.ResponseDTO {
	return &dto.ResponseDTO{
		Id:        model.ID,
		Name:      model.Name,
		Color:     model.Color.String(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ToModelFromCreateDTO(createDto dto.CreateDTO) (*domain.Line, []response.ErrorField) {
	colorVO, err := valueobject.NewColorVO(createDto.Color)
	if err != nil {
		return nil, []response.ErrorField{
			response.NewErrorField("color", string(response.BadRequest)),
		}
	}

	return &domain.Line{
		Name:  createDto.Name,
		Color: colorVO,
	}, nil
}
