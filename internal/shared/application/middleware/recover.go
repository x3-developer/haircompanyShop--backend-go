package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
	"serv_shop_haircompany/internal/shared/utils/response"
)

func Recoverer(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logger.Error("panic recovered",
						zap.Any("error", rec),
						zap.ByteString("stack", debug.Stack()),
						zap.String("method", r.Method),
						zap.String("path", r.URL.Path),
					)

					response.SendError(
						w,
						http.StatusInternalServerError,
						"panic occurred while processing the request",
						response.ServerError,
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
