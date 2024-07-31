package routes

import (
	"ApiRestaurant/src/controller"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
)

func SetupHamburguesaRoutes(router *gin.Engine, hamburguesaServices *services.HamburguesaServices) {
	hamburguesaController := controller.NewHamburguesaController(hamburguesaServices)

	router.GET("/hamburguesas", hamburguesaController.GetAllHamburguesas)
	router.GET("/hamburguesas/id/:id", hamburguesaController.GetHamburguesaById)
	router.GET("/hamburguesas/nombre/:name", hamburguesaController.GetHamburguesaByName)
	router.POST("/hamburguesas", hamburguesaController.CreateHamburguesa)
	router.DELETE("/hamburguesas/:id", hamburguesaController.DeleteHamburguesaById)
	router.PUT("/hamburguesas/:id", hamburguesaController.EditHamburguesaById)
	router.GET("/hamburguesas/price/:price", hamburguesaController.GetHamburguesaByPrice)
}
