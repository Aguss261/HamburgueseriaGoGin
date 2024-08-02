package routes

import (
	"ApiRestaurant/src/controller"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
)

func SetupPedidoRoutes(router *gin.Engine, pedidosServices *services.PedidoServices) {
	pedidoController := controller.NewPedidoController(pedidosServices)

	router.GET("/pedidos", pedidoController.GetAllPedidos)
	router.GET("/pedidos/id/:id", pedidoController.GetPedidoById)
	router.GET("/pedidos/user/:id", pedidoController.GetPedidoByUserId)
	router.GET("/pedidos/fecha/:fecha", pedidoController.GetPedidoByFecha)
	router.POST("/pedidos", pedidoController.CreatePedido)
	router.DELETE("/pedidos/:id", pedidoController.DeletePedido)
	router.PUT("/pedidos/:id", pedidoController.EditPedido)
}
