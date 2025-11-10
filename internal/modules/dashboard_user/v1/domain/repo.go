package domain

import "context"

type Repository interface {
	Create(ctx context.Context, model *DashboardUser) (*DashboardUser, error)
	ExistsByUniqueFields(ctx context.Context, email string) (bool, error)
}
