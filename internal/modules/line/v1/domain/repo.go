package domain

import "context"

type Repository interface {
	Create(ctx context.Context, model *Line) (*Line, error)
	FindUniqueConflicts(ctx context.Context, name string, excludeID uint) (conflicts []string, err error)
	GetById(ctx context.Context, id uint) (*Line, error)
	Update(ctx context.Context, model *Line) (*Line, error)
}
