package rest

import (
	"fmt"
	"net/http"
	"serv_shop_haircompany/internal/modules/line/v1/application/dto"
	"serv_shop_haircompany/internal/modules/line/v1/application/mapper"
	"serv_shop_haircompany/internal/modules/line/v1/application/usecase"
	sharedmapper "serv_shop_haircompany/internal/shared/application/mapper"
	"serv_shop_haircompany/internal/shared/utils/request"
	"serv_shop_haircompany/internal/shared/utils/response"
	"serv_shop_haircompany/internal/shared/utils/validator"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	createUC  usecase.CreateUseCase
	getByIdUC usecase.GetByIDUseCase
	updateUC  usecase.UpdateUseCase
	getListUC usecase.GetListUseCase
}

func NewHandler(createUC usecase.CreateUseCase, updateUC usecase.UpdateUseCase, getByIdUC usecase.GetByIDUseCase, getListUC usecase.GetListUseCase) Handler {
	return &handler{
		createUC:  createUC,
		getByIdUC: getByIdUC,
		updateUC:  updateUC,
		getListUC: getListUC,
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
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
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

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing line id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid line id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	model, err := h.getByIdUC.Execute(ctx, uint(id))
	if model == nil {
		msg := fmt.Sprintf("line with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve line: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	responseDTO := mapper.ToRespDTOFromModel(*model)

	response.SendSuccess(w, http.StatusOK, responseDTO)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		msg := "missing line id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid line id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	updateDTO, err := request.DecodeBody[dto.UpdateDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := validator.ValidateDTO(updateDTO)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	model, err := h.getByIdUC.Execute(ctx, uint(id))
	if err != nil {
		msg := fmt.Sprintf("failed to update line: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}
	if model == nil {
		msg := fmt.Sprintf("line with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}

	mapper.ToModelFromUpdateDTO(updateDTO, model)

	updatedModel, errFields, err := h.updateUC.Execute(ctx, model)
	if err != nil {
		msg := fmt.Sprintf("failed to update line: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	responseDTO := mapper.ToRespDTOFromModel(*updatedModel)

	response.SendSuccess(w, http.StatusOK, responseDTO)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	pageStr := q.Get("page")
	limitStr := q.Get("limit")

	page := 1
	limit := 20

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			msg := fmt.Sprintf("invalid page parameter: %s", pageStr)
			response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
			return
		}
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			msg := fmt.Sprintf("invalid limit parameter: %s", limitStr)
			response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
			return
		}
	}

	ctx := r.Context()
	paginatedResult, err := h.getListUC.Execute(ctx, page, limit)
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve line list: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	responseDTO := sharedmapper.ToPaginationResponseDTO(paginatedResult.Items, paginatedResult.Total)

	response.SendSuccess(w, http.StatusOK, responseDTO)
}
