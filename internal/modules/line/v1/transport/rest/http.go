package rest

import (
	"serv_shop_haircompany/internal/modules/line/v1/application/usecase"
	"serv_shop_haircompany/internal/modules/line/v1/infrastructure/persistence"
	"serv_shop_haircompany/internal/shared/application/container"
	"serv_shop_haircompany/internal/shared/application/middleware"
	"serv_shop_haircompany/internal/shared/domain/valueobject"

	"github.com/go-chi/chi/v5"
)

func LineV1Routes(r chi.Router, container *container.Container) {
	repo := persistence.NewRepo(container.PostgresDB)
	createUC := usecase.NewCreateUseCase(repo)
	h := NewHandler(createUC)

	const baseRoute = "/line"

	r.Group(func(r chi.Router) {
		r.Use(middleware.DashboardAuthMiddleware(container.TokenSvc))

		r.With(middleware.DashboardRoleMiddleware(valueobject.RoleAdmin)).Post(baseRoute, h.Create)
	})
}
