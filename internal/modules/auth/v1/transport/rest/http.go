package rest

import (
	"serv_shop_haircompany/internal/modules/auth/v1/application/usecase"
	"serv_shop_haircompany/internal/modules/auth/v1/infrastructure/persistence"
	dashboarduserpersistence "serv_shop_haircompany/internal/modules/dashboard_user/v1/infrastructure/persistence"
	"serv_shop_haircompany/internal/shared/application/container"

	"github.com/go-chi/chi/v5"
)

func AuthV1Routes(r chi.Router, container *container.Container) {
	repo := persistence.NewRepo(container.Redis)
	dashboardUserRepo := dashboarduserpersistence.NewRepo(container.PostgresDB)
	dashboardLoginUC := usecase.NewDashboardLoginUseCase(repo, dashboardUserRepo, container.TokenSvc)
	refreshUC := usecase.NewRefreshUseCase(repo, container.TokenSvc)
	h := NewHandler(dashboardLoginUC, refreshUC)

	const baseRoute = "/auth"

	r.Post(baseRoute+"/dashboard/login", h.DashboardLogin)
	r.Post(baseRoute+"/refresh", h.Refresh)
}
