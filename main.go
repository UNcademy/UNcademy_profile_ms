package main

import (
	"UNcademy_profile_ms/configs"
	route "UNcademy_profile_ms/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	router := SetupRouter()

	router.Run(":8000")
}

func SetupRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("Unable to load .env file")
	}

	db := configs.Connection()

	router := gin.Default()

	route.InitAuthRoutes(db, router)

	return router
}
