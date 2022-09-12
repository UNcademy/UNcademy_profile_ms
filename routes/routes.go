package routes

import (
	edit2 "UNcademy_profile_ms/controllers/edit"
	view2 "UNcademy_profile_ms/controllers/view"
	handlerEdit "UNcademy_profile_ms/handlers/edit"
	handlerView "UNcademy_profile_ms/handlers/view"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	editRepository := edit2.NewEditRepository(db)
	editService := edit2.NewEditService(editRepository)
	editHandler := handlerEdit.NewEditHandler(editService)

	viewRepository := view2.NewViewRepository(db)
	viewService := view2.NewViewService(viewRepository)
	viewHandler := handlerView.NewViewHandler(viewService)

	groupRoute := route.Group("/api/v1")
	groupRoute.PUT("/edit", editHandler.EditHandler)
	groupRoute.GET("/view", viewHandler.ViewHandler)

}

//Post: Crear - o cuando el request llega en un json
//Get: read
//Put: Update

//Flujo general:
//Request (route)
//Handle
//Service
//Repository
//Variable temporal
