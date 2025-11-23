package rest

import (
	"fmt"
	"net/http"
	"serv_shop_haircompany/internal/modules/auth/v1/application/dto"
	"serv_shop_haircompany/internal/modules/auth/v1/application/mapper"
	"serv_shop_haircompany/internal/modules/auth/v1/application/usecase"
	"serv_shop_haircompany/internal/shared/utils/request"
	"serv_shop_haircompany/internal/shared/utils/response"
	"serv_shop_haircompany/internal/shared/utils/validator"
)

type Handler interface {
	DashboardLogin(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	dashboardLoginUC usecase.DashboardLoginUseCase
	refreshUC        usecase.RefreshUseCase
}

func NewHandler(dashboardLoginUC usecase.DashboardLoginUseCase, refreshUC usecase.RefreshUseCase) Handler {
	return &handler{
		dashboardLoginUC: dashboardLoginUC,
		refreshUC:        refreshUC,
	}
}

func (h *handler) DashboardLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dashboardAuthDTO, err := request.DecodeBody[dto.DashboardAuthDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := validator.ValidateDTO(dashboardAuthDTO)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	model, errFields := mapper.ToModelFromDashboardAuthDTO(dashboardAuthDTO)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	tokenPair, err := h.dashboardLoginUC.Execute(ctx, model)
	if err != nil {
		msg := fmt.Sprintf("login failed: %v", err)
		response.SendError(w, http.StatusUnauthorized, msg, response.Unauthorized)
		return
	}

	responseDTO := mapper.ToRespDTOFromModel(*tokenPair)

	response.SendSuccess(w, http.StatusCreated, responseDTO)
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	refreshDTO, err := request.DecodeBody[dto.RefreshDTO](r.Body)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, "invalid request body", response.BadRequest)
		return
	}

	errFields := validator.ValidateDTO(refreshDTO)
	if errFields != nil {
		response.SendValidationError(w, http.StatusBadRequest, "validation errors occurred", response.BadRequest, errFields)
		return
	}

	tokenPair, err := h.refreshUC.Execute(ctx, refreshDTO.RefreshToken)
	if err != nil {
		response.SendError(w, http.StatusUnauthorized, "refresh failed: "+err.Error(), response.Unauthorized)
		return
	}

	respDTO := mapper.ToRespDTOFromModel(*tokenPair)

	response.SendSuccess(w, http.StatusOK, respDTO)
}
