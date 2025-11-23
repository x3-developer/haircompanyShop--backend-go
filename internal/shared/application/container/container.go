package container

import (
	"serv_shop_haircompany/internal/config"
	"serv_shop_haircompany/internal/shared/infrastructure/persistence"
	"serv_shop_haircompany/internal/shared/infrastructure/security"

	"go.uber.org/zap"
)

type Container struct {
	Config     *config.Config
	PostgresDB *persistence.Postgres
	Redis      *persistence.Redis
	Logger     *zap.Logger
	TokenSvc   security.TokenService
}

func NewContainer(cfg *config.Config, logger *zap.Logger) *Container {
	pdb := persistence.NewPostgres(cfg, logger)
	redis := persistence.NewRedis(cfg, logger)
	tokenSvc := security.NewTokenService(cfg.AppSecret)

	return &Container{
		Config:     cfg,
		PostgresDB: pdb,
		Redis:      redis,
		Logger:     logger,
		TokenSvc:   tokenSvc,
	}
}
