package rest

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	linehttp "serv_shop_haircompany/internal/modules/line/v1/transport/rest"
	"serv_shop_haircompany/internal/shared/application/container"
	"serv_shop_haircompany/internal/shared/application/middleware"
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
		linehttp.LineV1Routes(r, container)
	})

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return r
}
