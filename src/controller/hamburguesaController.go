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

// HamburguesaController es el controlador para manejar las peticiones relacionadas con hamburguesas
type HamburguesaController struct {
	HamburguesaServices *services.HamburguesaServices
}

// NewHamburguesaController crea una nueva instancia de HamburguesaController
func NewHamburguesaController(hamburguesaServices *services.HamburguesaServices) *HamburguesaController {
	return &HamburguesaController{HamburguesaServices: hamburguesaServices}
}

// GetAllHamburguesas maneja la solicitud para obtener todas las hamburguesass
// @Summary Obtener todas las hamburguesas
// @Description Obtiene todas las hamburguesas
// @Tags hamburguesas
// @Produce json
// @Success 200 {object} entity.SuccessResponse{data=[]entity.Hamburguesa}
// @Failure 500 {object} entity.ErrorResponse "Internal server error"
// @Router /hamburguesas [get]
func (hs *HamburguesaController) GetAllHamburguesas(c *gin.Context) {
	hamburguesas, err := hs.HamburguesaServices.GetAllHamburguesas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hamburguesas})
}

// GetHamburguesaById maneja la solicitud para obtener una hamburguesa por ID
// @Summary Obtener hamburguesa por ID
// @Description Obtiene una hamburguesa basada en el ID proporcionado
// @Tags hamburguesas
// @Produce json
// @Param id path int true "ID de la hamburguesa"
// @Success 200 {object} entity.SuccessResponse{data=entity.Hamburguesa}
// @Failure 400 {object} entity.ErrorResponse "ID invalido"
// @Failure 404 {object} entity.ErrorResponse "Hamburguesa no encontrada"
// @Failure 500 {object} entity.ErrorResponse "Error interno del servidor"
// @Router /hamburguesas/id/{id} [get]
func (hs *HamburguesaController) GetHamburguesaById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID invalido"})
		return
	}
	hamburguesa, err := hs.HamburguesaServices.GetById(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hamburguesa})
}

// GetHamburguesaByName maneja la solicitud para obtener una hamburguesa por su nombre
// @Summary Obtener hamburguesa por nombre
// @Description Obtiene una hamburguesa basada en el nombre proporcionado
// @Tags hamburguesas
// @Produce json
// @Param name path string true "Nombre de la hamburguesa"
// @Success 200 {object} entity.SuccessResponse{data=entity.Hamburguesa}
// @Failure 400 {object} entity.ErrorResponse "Nombre inválido"
// @Failure 404 {object} entity.ErrorResponse "Hamburguesa no encontrada"
// @Failure 500 {object} entity.ErrorResponse "Error interno del servidor"
// @Router /hamburguesas/nombre/{name} [get]
func (hs *HamburguesaController) GetHamburguesaByName(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre invalido"})
		return
	}

	hamburguesas, err := hs.HamburguesaServices.GetByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(hamburguesas) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hamburguesa no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hamburguesas})
}

// GetHamburguesaByPrice maneja la solicitud para obtener hamburguesas con un precio mayor al proporcionado
// @Summary Obtener hamburguesas por precio
// @Description Obtiene una lista de hamburguesas cuyo precio es mayor al proporcionado
// @Tags hamburguesas
// @Produce json
// @Param price path float64 true "Precio mínimo de la hamburguesa"
// @Success 200 {object} entity.SuccessResponse{data=[]entity.Hamburguesa}
// @Failure 400 {object} entity.ErrorResponse "Precio inválido"
// @Failure 500 {object} entity.ErrorResponse "Error interno del servidor"
// @Router /hamburguesas/price/{price} [get]
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

// CreateHamburguesa maneja la solicitud para crear una nueva hamburguesa
// @Summary Crear nueva hamburguesa
// @Description Crea una nueva hamburguesa con los datos proporcionados
// @Tags hamburguesas
// @Accept json
// @Produce json
// @Param hamburguesa body entity.Hamburguesa true "Datos de la hamburguesa"
// @Success 201 {object} entity.SuccessResponse{data=entity.Hamburguesa}
// @Failure 400 {object} entity.ErrorResponse "Datos inválidos"
// @Failure 401 {object} entity.ErrorResponse "Unauthorized: Invalid user ID"
// @Failure 403 {object} entity.ErrorResponse "Forbidden: Access denied"
// @Failure 500 {object} entity.ErrorResponse "Error interno del servidor"
// @Router /hamburguesas [post]
// @Security Bearer
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

// DeleteHamburguesaById maneja la solicitud para eliminar una hamburguesa por su ID
// @Summary Eliminar hamburguesa por ID
// @Description Elimina una hamburguesa basada en el ID proporcionado
// @Tags hamburguesas
// @Produce json
// @Param id path int true "ID de la hamburguesa"
// @Success 200 {object} entity.SuccessResponse{data=boolean}
// @Failure 400 {object} entity.ErrorResponse "ID inválido"
// @Failure 404 {object} entity.ErrorResponse "Hamburguesa no encontrada"
// @Failure 401 {object} entity.ErrorResponse "Unauthorized: Invalid user ID"
// @Failure 403 {object} entity.ErrorResponse "Forbidden: Access denied"
// @Failure 500 {object} entity.ErrorResponse "Error interno del servidor"
// @Router /hamburguesas/{id} [delete]
// @Security Bearer
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

// EditHamburguesaById maneja la solicitud para actualizar una hamburguesa por su ID
// @Summary Actualizar hamburguesa por ID
// @Description Actualiza una hamburguesa existente basada en el ID proporcionado y los datos JSON enviados
// @Tags hamburguesas
// @Produce json
// @Param id path int true "ID de la hamburguesa"
// @Param body body entity.Hamburguesa true "Datos de la hamburguesa a actualizar"
// @Success 200 {object} entity.SuccessResponse{message=string}
// @Failure 400 {object} entity.ErrorResponse "JSON inválido"
// @Failure 404 {object} entity.ErrorResponse "Hamburguesa no encontrada"
// @Failure 401 {object} entity.ErrorResponse "Unauthorized: Invalid user ID"
// @Failure 403 {object} entity.ErrorResponse "Forbidden: Access denied"
// @Failure 500 {object} entity.ErrorResponse "Error interno del servidor"
// @Router /hamburguesas/{id} [put]
// @Security Bearer
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
