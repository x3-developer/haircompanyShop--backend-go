package domain

import "context"

type Repository interface {
	Create(ctx context.Context, model *Line) (*Line, error)
	ExistsByUniqueFields(ctx context.Context, name string) (bool, error)
}
