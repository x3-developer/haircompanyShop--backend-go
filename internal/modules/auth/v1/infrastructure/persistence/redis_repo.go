package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"serv_shop_haircompany/internal/modules/auth/v1/domain"
	"serv_shop_haircompany/internal/shared/infrastructure/persistence"
	"time"
)

type repo struct {
	DB *persistence.Redis
}

func NewRepo(db *persistence.Redis) domain.SessionRepository {
	return &repo{
		DB: db,
	}
}

func (r *repo) Save(ctx context.Context, sess *domain.RefreshSession, ttl time.Duration) error {
	b, _ := json.Marshal(sess)

	return r.DB.Client.Set(ctx, r.getKey(sess.UserID, sess.JTI), b, ttl).Err()
}

func (r *repo) Get(ctx context.Context, userID uint, jti string) (*domain.RefreshSession, error) {
	v, err := r.DB.Client.Get(ctx, r.getKey(userID, jti)).Bytes()
	if err != nil {
		return nil, err
	}

	var sess domain.RefreshSession
	if err := json.Unmarshal(v, &sess); err != nil {
		return nil, err
	}

	return &sess, nil
}

func (r *repo) Delete(ctx context.Context, userID uint, jti string) error {
	return r.DB.Client.Del(ctx, r.getKey(userID, jti)).Err()
}

func (r *repo) getKey(userID uint, jti string) string {
	return fmt.Sprintf("rt:%d:%s", userID, jti)
}
