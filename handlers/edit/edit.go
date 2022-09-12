package edit

import (
	update2 "UNcademy_profile_ms/controllers/edit"
	util "UNcademy_profile_ms/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	service update2.Service
}

func NewEditHandler(service update2.Service) *handler {
	return &handler{service: service}
}

func (h *handler) EditHandler(ctx *gin.Context) {
	var input update2.InputEdit
	username := ctx.Query("username")
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		defer logrus.Error(err.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {

		errEdit := h.service.EditService(&input, username)

		switch errEdit {

		case "USER_NOT_FOUND_404":
			util.APIResponse(ctx, "User not found", http.StatusNotFound, http.MethodPut, nil)
		case "USER_NOT_ACTIVE_403":
			util.APIResponse(ctx, "User not active", http.StatusForbidden, http.MethodPut, nil)
		case "CHANGE_FAILED_403":
			util.APIResponse(ctx, "Change failed", http.StatusForbidden, http.MethodPut, nil)
		default:
			util.APIResponse(ctx, "Changed successfully!", http.StatusOK, http.MethodPut, nil)
		}
	}
}
