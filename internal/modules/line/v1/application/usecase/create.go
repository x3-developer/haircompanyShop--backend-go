package usecase

import (
	"context"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	"serv_shop_haircompany/internal/shared/utils/response"
)

type CreateUseCase interface {
	Execute(ctx context.Context, model *domain.Line) (*domain.Line, []response.ErrorField, error)
}

type createUseCase struct {
	repo domain.Repository
}

func NewCreateUseCase(repo domain.Repository) CreateUseCase {
	return &createUseCase{
		repo: repo,
	}
}

func (u *createUseCase) Execute(ctx context.Context, model *domain.Line) (*domain.Line, []response.ErrorField, error) {
	existingModel, err := u.repo.ExistsByUniqueFields(ctx, model.Name)
	if err != nil {
		return nil, nil, err
	}
	if existingModel {
		var validationErrors []response.ErrorField
		validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		return nil, validationErrors, nil
	}

	createdModel, err := u.repo.Create(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return createdModel, nil, nil
}
