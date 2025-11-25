package query

import (
	"context"
	"serv_shop_haircompany/internal/modules/line/v1/application/query/list"
	"serv_shop_haircompany/internal/shared/domain"
)

type Repository interface {
	GetPaginatedList(ctx context.Context, page, limit int) (*domain.PaginatedResult[list.LineListReadModel], error)
}
