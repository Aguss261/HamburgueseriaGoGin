package routes

import (
	"ApiRestaurant/src/controller"
	"ApiRestaurant/src/middleware"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
)

func SetupHamburguesaRoutes(router *gin.Engine, hamburguesaServices *services.HamburguesaServices) {
	hamburguesaController := controller.NewHamburguesaController(hamburguesaServices)

	// Rutas protegidas por JWT
	protected := router.Group("/")
	protected.Use(middleware.JWTRequired())

	protected.GET("/hamburguesas", hamburguesaController.GetAllHamburguesas)
	protected.GET("/hamburguesas/id/:id", hamburguesaController.GetHamburguesaById)
	protected.GET("/hamburguesas/nombre/:name", hamburguesaController.GetHamburguesaByName)
	protected.POST("/hamburguesas", hamburguesaController.CreateHamburguesa)
	protected.DELETE("/hamburguesas/:id", hamburguesaController.DeleteHamburguesaById)
	protected.PUT("/hamburguesas/:id", hamburguesaController.EditHamburguesaById)
	protected.GET("/hamburguesas/price/:price", hamburguesaController.GetHamburguesaByPrice)
}
