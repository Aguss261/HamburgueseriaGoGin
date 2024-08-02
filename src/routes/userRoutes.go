package routes

import (
	"ApiRestaurant/src/controller"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userServices *services.UserService) {
	userController := controller.NewUserController(userServices)
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.GET("/usuarios", userController.GetUsers)
}
