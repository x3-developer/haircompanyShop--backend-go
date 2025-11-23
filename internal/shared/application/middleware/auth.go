package middleware

import (
	"context"
	"net/http"
	"serv_shop_haircompany/internal/shared/infrastructure/security"
	"serv_shop_haircompany/internal/shared/utils/response"
	"strings"
)

type dashboardClaimsKey struct{}

func DashboardAuthMiddleware(tokenSvc security.TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				response.SendError(w, http.StatusUnauthorized, "Unauthorized", response.Unauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := tokenSvc.ParseAccessToken(token)
			if err != nil || claims == nil {
				response.SendError(w, http.StatusUnauthorized, "Unauthorized", response.Unauthorized)
				return
			}

			if claims.Role == "" {
				response.SendError(w, http.StatusForbidden, "Forbidden", response.Forbidden)
				return
			}

			ctx := context.WithValue(r.Context(), dashboardClaimsKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
