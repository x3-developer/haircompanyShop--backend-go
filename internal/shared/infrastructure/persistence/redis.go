package persistence

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"serv_shop_haircompany/internal/config"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(cfg *config.Config, logger *zap.Logger) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
		Protocol: 2,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal("failed to connect to redis", zap.Error(err))
	}

	return &Redis{
		Client: client,
	}
}
