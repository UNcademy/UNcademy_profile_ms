package view

import (
	see2 "UNcademy_profile_ms/controllers/view"
	util "UNcademy_profile_ms/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service see2.Service
}

func NewViewHandler(service see2.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ViewHandler(ctx *gin.Context) {

	username := ctx.Query("username")
	user, errView := h.service.ViewService(username)

	switch errView {

	case "LOGIN_NOT_FOUND_404":
		util.APIResponse(ctx, "User not found", http.StatusNotFound, http.MethodPut, nil)
	case "LOGIN_NOT_ACTIVE_403":
		util.APIResponse(ctx, "User not active", http.StatusForbidden, http.MethodPut, nil)
	default:
		util.APIResponse(ctx, "Changed successfully!", http.StatusOK, http.MethodPut, user)
	}

}
