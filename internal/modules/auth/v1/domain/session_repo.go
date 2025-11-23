package domain

import (
	"context"
	"time"
)

type SessionRepository interface {
	Save(ctx context.Context, sess *RefreshSession, ttl time.Duration) error
	Get(ctx context.Context, userID uint, jti string) (*RefreshSession, error)
	Delete(ctx context.Context, userID uint, jti string) error
}
