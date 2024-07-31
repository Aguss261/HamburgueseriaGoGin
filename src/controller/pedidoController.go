package controller

import (
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PedidoController struct {
	PedidoServices *services.PedidoServices
}

func NewPedidoController(pedidoServices *services.PedidoServices) *PedidoController {
	return &PedidoController{PedidoServices: pedidoServices}
}

func (ps *PedidoController) GetAllPedidos(c *gin.Context) {
	pedidos, err := ps.PedidoServices.GetAllPedidos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pedidos})
}

func (ps *PedidoController) GetPedidoById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pedidos, err := ps.PedidoServices.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pedidos})
}
