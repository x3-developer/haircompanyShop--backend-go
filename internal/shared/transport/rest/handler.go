package rest

import (
	"net/http"
	"serv_shop_haircompany/internal/shared/utils/response"
)

func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	response.SendError(w, http.StatusNotFound, "route not found", response.NotFound)
}

func MethodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
	response.SendError(w, http.StatusMethodNotAllowed, "method not allowed", response.MethodNotAllowed)
}
