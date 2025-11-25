package usecase

import (
	"context"
	"errors"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	"serv_shop_haircompany/internal/shared/utils/response"
)

type UpdateUseCase interface {
	Execute(ctx context.Context, model *domain.Line) (*domain.Line, []response.ErrorField, error)
}

type updateUseCase struct {
	repo domain.Repository
}

func NewUpdateUseCase(repo domain.Repository) CreateUseCase {
	return &updateUseCase{
		repo: repo,
	}
}

func (u *updateUseCase) Execute(ctx context.Context, model *domain.Line) (*domain.Line, []response.ErrorField, error) {
	existingModel, err := u.repo.GetById(ctx, model.ID)
	if err != nil {
		return nil, nil, err
	}
	if existingModel == nil {
		return nil, nil, errors.New("line not found")
	}

	conflicts, err := u.repo.FindUniqueConflicts(ctx, model.Name, model.ID)
	if err != nil {
		return nil, nil, err
	}
	if len(conflicts) > 0 {
		var validationErrors []response.ErrorField

		for _, field := range conflicts {
			validationErrors = append(validationErrors, response.NewErrorField(field, string(response.NotUnique)))
		}

		return nil, validationErrors, nil
	}

	updatedModel, err := u.repo.Update(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return updatedModel, nil, nil
}
