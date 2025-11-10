package persistence

import (
	"context"
	"serv_shop_haircompany/internal/modules/dashboard_user/v1/domain"
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

func (r *repo) Create(ctx context.Context, model *domain.DashboardUser) (*domain.DashboardUser, error) {
	const q = `
		INSERT INTO dashboard_users (email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.DB.Pool.QueryRow(
		ctx,
		q,
		model.Email.String(),
		string(model.Password),
		string(model.Role),
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

func (r *repo) ExistsByUniqueFields(ctx context.Context, email string) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM dashboard_users WHERE email = $1);`

	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	var exists bool
	err := r.DB.Pool.QueryRow(ctx, q, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
