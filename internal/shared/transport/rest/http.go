package rest

import (
	"net/http"
	authhttp "serv_shop_haircompany/internal/modules/auth/v1/transport/rest"
	dashboarduserhttp "serv_shop_haircompany/internal/modules/dashboard_user/v1/transport/rest"
	dboarduserhttp "serv_shop_haircompany/internal/modules/dashboard_user/v1/transport/rest"
	linehttp "serv_shop_haircompany/internal/modules/line/v1/transport/rest"
	"serv_shop_haircompany/internal/shared/application/container"
	"serv_shop_haircompany/internal/shared/application/middleware"

	"github.com/go-chi/chi/v5"
)

func NewHTTPRouter(container *container.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger(container.Logger))
	r.Use(middleware.Recoverer(container.Logger))
	r.Use(middleware.CORSMiddleware(container.Config.CORS))
	r.Use(middleware.APIMiddleware(container.Config.AuthAppKey))

	r.NotFound(NotFoundHandler)
	r.MethodNotAllowed(MethodNotAllowedHandler)

	r.Route("/api/v1", func(r chi.Router) {
		dboarduserhttp.DashboardUserV1Routes(r, container)
		dashboarduserhttp.DashboardUserV1Routes(r, container)
		authhttp.AuthV1Routes(r, container)
		linehttp.LineV1Routes(r, container)
	})

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return r
}
