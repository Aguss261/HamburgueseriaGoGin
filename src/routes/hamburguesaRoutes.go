package routes

import (
	"ApiRestaurant/src/controller"
	"ApiRestaurant/src/middleware"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
)

func SetupHamburguesaRoutes(router *gin.Engine, hamburguesaServices *services.HamburguesaServices, userServices *services.UserService) {
	hamburguesaController := controller.NewHamburguesaController(hamburguesaServices)

	protected := router.Group("/")
	protected.Use(middleware.JWTRequired())
	admin := router.Group("/")
	admin.Use(middleware.AdminRequired(userServices))

	router.GET("/hamburguesas", hamburguesaController.GetAllHamburguesas)
	router.GET("/hamburguesas/id/:id", hamburguesaController.GetHamburguesaById)
	router.GET("/hamburguesas/nombre/:name", hamburguesaController.GetHamburguesaByName)
	protected.POST("/hamburguesas", hamburguesaController.CreateHamburguesa)
	admin.DELETE("/hamburguesas/:id", hamburguesaController.DeleteHamburguesaById)
	admin.PUT("/hamburguesas/:id", hamburguesaController.EditHamburguesaById)
	router.GET("/hamburguesas/price/:price", hamburguesaController.GetHamburguesaByPrice)
}
