package middleware

import (
	"net/http"
	"serv_shop_haircompany/internal/shared/domain/valueobject"
	"serv_shop_haircompany/internal/shared/infrastructure/security"
	"serv_shop_haircompany/internal/shared/utils/response"
)

func DashboardRoleMiddleware(allowedRoles ...valueobject.DashboardRole) func(http.Handler) http.Handler {
	allowed := make(map[string]struct{}, len(allowedRoles))
	for _, r := range allowedRoles {
		allowed[r.String()] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			raw := r.Context().Value(dashboardClaimsKey{})
			if raw == nil {
				response.SendError(w, http.StatusUnauthorized, "Unauthorized", response.Unauthorized)
				return
			}

			claims := raw.(*security.AccessClaims)

			if _, ok := allowed[claims.Role]; !ok {
				response.SendError(w, http.StatusForbidden, "Forbidden", response.Forbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
