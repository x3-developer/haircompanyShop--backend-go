package usecase

import (
	"context"
	"serv_shop_haircompany/internal/modules/line/v1/application/query"
	"serv_shop_haircompany/internal/modules/line/v1/application/query/list"
	shareddomain "serv_shop_haircompany/internal/shared/domain"
)

type GetListUseCase interface {
	Execute(ctx context.Context, page, perPage int) (*shareddomain.PaginatedResult[list.LineListReadModel], error)
}

type getListUseCase struct {
	repo query.Repository
}

func NewGetListUseCase(repo query.Repository) GetListUseCase {
	return &getListUseCase{
		repo: repo,
	}
}

func (u *getListUseCase) Execute(ctx context.Context, page, limit int) (*shareddomain.PaginatedResult[list.LineListReadModel], error) {
	models, err := u.repo.GetPaginatedList(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return models, nil
}
