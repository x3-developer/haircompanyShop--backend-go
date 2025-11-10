package container

import (
	"go.uber.org/zap"
	"serv_shop_haircompany/internal/config"
	"serv_shop_haircompany/internal/shared/infrastructure/persistence"
)

type Container struct {
	Config     *config.Config
	PostgresDB *persistence.Postgres
	Redis      *persistence.Redis
	Logger     *zap.Logger
}

func NewContainer(cfg *config.Config, logger *zap.Logger) *Container {
	pdb := persistence.NewPostgres(cfg, logger)
	redis := persistence.NewRedis(cfg, logger)

	return &Container{
		Config:     cfg,
		PostgresDB: pdb,
		Redis:      redis,
		Logger:     logger,
	}
}
