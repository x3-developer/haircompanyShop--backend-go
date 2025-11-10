package rest

import (
	"github.com/go-chi/chi/v5"
	"serv_shop_haircompany/internal/modules/line/v1/application/usecase"
	"serv_shop_haircompany/internal/modules/line/v1/infrastructure/persistence"
	"serv_shop_haircompany/internal/shared/application/container"
)

func LineV1Routes(r chi.Router, container *container.Container) {
	repo := persistence.NewRepo(container.PostgresDB)
	createUC := usecase.NewCreateUseCase(repo)
	h := NewHandler(createUC)

	const baseRoute = "/line"

	r.Post(baseRoute, h.Create)
}
