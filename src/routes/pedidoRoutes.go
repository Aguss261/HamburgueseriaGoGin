package routes

import (
	"ApiRestaurant/src/controller"
	"ApiRestaurant/src/middleware"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
)

func SetupPedidoRoutes(router *gin.Engine, pedidosServices *services.PedidoServices, userServices *services.UserService) {
	pedidoController := controller.NewPedidoController(pedidosServices)

	protected := router.Group("/")
	protected.Use(middleware.JWTRequired())
	admin := router.Group("/")
	admin.Use(middleware.AdminRequired(userServices))

	admin.GET("/pedidos", pedidoController.GetAllPedidos)
	protected.GET("/pedidos/id/:id", pedidoController.GetPedidoById)
	protected.GET("/pedidos/user", pedidoController.GetPedidoByUserId)
	admin.GET("/pedidos/fecha/:fecha", pedidoController.GetPedidoByFecha)
	protected.POST("/pedidos", pedidoController.CreatePedido)
	protected.DELETE("/pedidos/:id", pedidoController.DeletePedido)
	protected.PUT("/pedidos/:id", pedidoController.EditPedido)
}
