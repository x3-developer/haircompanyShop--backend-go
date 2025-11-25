package persistence

import (
	"context"
	"serv_shop_haircompany/internal/modules/line/v1/application/query"
	"serv_shop_haircompany/internal/modules/line/v1/application/query/list"
	"serv_shop_haircompany/internal/shared/domain"
	"serv_shop_haircompany/internal/shared/infrastructure/persistence"
	"time"
)

type queryRepo struct {
	DB *persistence.Postgres
}

func NewQueryRepo(db *persistence.Postgres) query.Repository {
	return &queryRepo{
		DB: db,
	}
}

func (r *queryRepo) GetPaginatedList(ctx context.Context, page, limit int) (*domain.PaginatedResult[list.LineListReadModel], error) {
	const q = `
		SELECT id, name, color, COUNT(*) OVER() as total
		FROM lines
		ORDER BY id
		LIMIT $1 OFFSET $2;
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	offset := (page - 1) * limit

	rows, err := r.DB.Pool.Query(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := &domain.PaginatedResult[list.LineListReadModel]{
		Items: make([]list.LineListReadModel, 0),
	}
	var total int64

	for rows.Next() {
		var item list.LineListReadModel

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Color,
			&total,
		); err != nil {
			return nil, err
		}

		result.Items = append(result.Items, item)
	}

	result.Total = total

	return result, nil
}
