package controller

import (
	"ApiRestaurant/src/entity"
	"ApiRestaurant/src/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

var db *sql.DB

type HamburguesaController struct {
	HamburguesaServices *services.HamburguesaServices
}

func NewHamburguesaController(hamburguesaServices *services.HamburguesaServices) *HamburguesaController {
	return &HamburguesaController{HamburguesaServices: hamburguesaServices}
}

func (hs *HamburguesaController) GetAllHamburguesas(c *gin.Context) {
	hamburguesas, err := hs.HamburguesaServices.GetAllHamburguesas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hamburguesas})
}

func (hs *HamburguesaController) GetHamburguesaById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hamburguesa id invalido"})
		return
	}
	hamburguesa, err := hs.HamburguesaServices.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hamburguesa})
}

func (hs *HamburguesaController) GetHamburguesaByName(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hamburguesa name invalido"})
		return
	}

	hamburguesas, err := hs.HamburguesaServices.GetByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hamburguesas})
}

func (hs *HamburguesaController) GetHamburguesaByPrice(c *gin.Context) {
	price, err := strconv.ParseFloat(c.Param("price"), 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hamburguesa price invalido"})
	}

	if price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hamburguesa price invalido"})
		return
	}

	hamburguesas, err := hs.HamburguesaServices.GetByPrice(price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hamburguesas})
}

func (hs *HamburguesaController) CreateHamburguesa(c *gin.Context) {
	var hamburguesa entity.Hamburguesa
	if err := c.ShouldBindJSON(&hamburguesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON", "details": err.Error()})
		return
	}
	isValid, errorMsg := validarHamburguesa(&hamburguesa)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
		return
	}
	isValid2, errorMsg2 := validarIngredientes(&hamburguesa.Ingrediente)
	if !isValid2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg2})
		return
	}
	if err := hs.HamburguesaServices.CreateHamburguesa(&hamburguesa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": hamburguesa})

}

func (hs *HamburguesaController) DeleteHamburguesaById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hamburguesa id invalido"})
		return
	}
	err2 := hs.HamburguesaServices.DeleteHamburguesaByiD(int(id))
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (hs *HamburguesaController) EditHamburguesaById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hamburguesa id invalido"})
	}
	var hamburguesa entity.Hamburguesa
	if err := c.BindJSON(&hamburguesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	if err := hs.HamburguesaServices.EditHamburguesa(int(id), hamburguesa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hamburguesa actualizada correctamente"})
}

func validarHamburguesa(hamburguesa *entity.Hamburguesa) (bool, string) {
	if strings.TrimSpace(hamburguesa.Nombre) == "" {
		return false, "El campo Nombre no puede estar vacío"
	}
	if hamburguesa.Price <= 0 {
		return false, "El campo Price debe ser mayor a 0"
	}
	if strings.TrimSpace(hamburguesa.Descripcion) == "" {
		return false, "El campo Descripcion no puede estar vacío"
	}
	if strings.TrimSpace(hamburguesa.ImgUrl) == "" {
		return false, "El campo ImgUrl no puede estar vacío"
	}

	return true, ""
}

func validarIngredientes(ingrediente *entity.Ingrediente) (bool, string) {

	if ingrediente.Huevo < 0 || ingrediente.Lechuga < 0 ||
		ingrediente.Tomate < 0 || ingrediente.Cebolla < 0 ||
		ingrediente.Bacon < 0 || ingrediente.Pepino < 0 {
		msg := "Los campos de 'Ingredientes' deben ser no negativos"
		return false, msg
	}
	return true, ""
}
