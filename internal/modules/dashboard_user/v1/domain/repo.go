package domain

import "context"

type Repository interface {
	Create(ctx context.Context, model *DashboardUser) (*DashboardUser, error)
	FindByEmail(ctx context.Context, email string) (*DashboardUser, error)
	ExistsByUniqueFields(ctx context.Context, email string) (bool, error)
}
