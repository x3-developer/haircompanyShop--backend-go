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

func (r *repo) GetById(ctx context.Context, id uint) (*domain.Line, error) {
	const q = `
		SELECT id, name, color, created_at, updated_at
		FROM lines
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var model domain.Line
	err := r.DB.Pool.QueryRow(ctx, q, id).Scan(
		&model.ID,
		&model.Name,
		&model.Color,
		&model.CreatedAt,
		&model.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *repo) FindUniqueConflicts(ctx context.Context, name string, id uint) (conflicts []string, err error) {
	const q = `SELECT name FROM lines WHERE name = $1 AND id <> $2;`

	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	rows, err := r.DB.Pool.Query(ctx, q, name, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var conflictName string
		if err := rows.Scan(&conflictName); err != nil {
			return nil, err
		}
		conflicts = append(conflicts, "name")
	}

	return conflicts, nil

}

func (r *repo) Update(ctx context.Context, model *domain.Line) (*domain.Line, error) {
	const q = `
		UPDATE lines
		SET name = $1, color = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.DB.Pool.QueryRow(ctx, q,
		model.ID,
		model.Name,
		model.Color,
	).Scan(
		&model.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return model, nil
}
