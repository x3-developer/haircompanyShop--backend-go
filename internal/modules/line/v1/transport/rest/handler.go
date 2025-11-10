package rest

import (
	"fmt"
	"net/http"
	"serv_shop_haircompany/internal/modules/line/v1/application/dto"
	"serv_shop_haircompany/internal/modules/line/v1/application/mapper"
	"serv_shop_haircompany/internal/modules/line/v1/application/usecase"
	"serv_shop_haircompany/internal/shared/utils/request"
	"serv_shop_haircompany/internal/shared/utils/response"
	"serv_shop_haircompany/internal/shared/utils/validator"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	createUC usecase.CreateUseCase
}

func NewHandler(createUC usecase.CreateUseCase) Handler {
	return &handler{
		createUC: createUC,
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	createDTO, err := request.DecodeBody[dto.CreateDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := validator.ValidateDTO(createDTO)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	model, errFields := mapper.ToModelFromCreateDTO(createDTO)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	createdModel, errFields, err := h.createUC.Execute(ctx, model)
	if err != nil {
		msg := fmt.Sprintf("failed to create line: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	responseDTO := mapper.ToRespDTOFromModel(*createdModel)

	response.SendSuccess(w, http.StatusCreated, responseDTO)
}
