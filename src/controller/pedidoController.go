package controller

import (
	"ApiRestaurant/src/entity"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

func (ps *PedidoController) GetPedidoByUserId(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pedidos, err2 := ps.PedidoServices.GetByUserId(int(id))
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pedidos})
}

func (ps *PedidoController) GetPedidoByFecha(c *gin.Context) {
	fecha := c.Param("fecha")
	if fecha == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "la fecha es requerida"})
	}
	pedidos, err := ps.PedidoServices.GetByFecha(fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if pedidos == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no se realizaron pedidos en esa fecha"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pedidos})
}

func (ps *PedidoController) CreatePedido(c *gin.Context) {
	var pedido entity.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pedido.Hora = time.Now()
	pedido.Fecha = time.Now()
	pedido.State = "Pendiente"
	pedido.Price = ps.calcularPrecio(&pedido.Hamburguesas)

	if err := ps.PedidoServices.CreatePedido(&pedido); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": pedido})
}

func (ps *PedidoController) DeletePedido(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hamburguesa id invalido"})
		return
	}
	err2 := ps.PedidoServices.DeletePedido(int(id))
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (ps *PedidoController) EditPedido(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hamburguesa id invalido"})
	}
	var pedido entity.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	pedido.Price = ps.calcularPrecio(&pedido.Hamburguesas)

	if err := ps.PedidoServices.UpdatePedido(int(id), pedido); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (ps *PedidoController) calcularPrecio(hamburguesas *[]entity.Hamburguesa) float32 {
	var devolver float32
	for _, hamburguesa := range *hamburguesas {
		devolver += hamburguesa.Price
	}
	return devolver
}
