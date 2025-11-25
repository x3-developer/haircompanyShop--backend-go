package usecase

import (
	"context"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
)

type GetByIDUseCase interface {
	Execute(ctx context.Context, id uint) (*domain.Line, error)
}

type getByIDUseCase struct {
	repo domain.Repository
}

func NewGetByIDUseCase(repo domain.Repository) GetByIDUseCase {
	return &getByIDUseCase{
		repo: repo,
	}
}

func (u *getByIDUseCase) Execute(ctx context.Context, id uint) (*domain.Line, error) {
	model, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return model, nil
}
