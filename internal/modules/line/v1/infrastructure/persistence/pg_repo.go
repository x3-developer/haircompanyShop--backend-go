package persistence

import (
	"context"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	"serv_shop_haircompany/internal/shared/infrastructure/persistence"
	"time"
)

type repo struct {
	DB *persistence.Postgres
}

func NewRepo(db *persistence.Postgres) domain.Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) Create(ctx context.Context, model *domain.Line) (*domain.Line, error) {
	const q = `
		INSERT INTO lines (name, color, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.DB.Pool.QueryRow(ctx, q,
		model.Name,
		model.Color,
	).Scan(
		&model.ID,
		&model.CreatedAt,
		&model.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *repo) ExistsByUniqueFields(ctx context.Context, name string) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM lines WHERE name = $1);`

	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	var exists bool
	err := r.DB.Pool.QueryRow(ctx, q, name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
