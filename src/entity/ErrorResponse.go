package entity

// ErrorResponse es la estructura para los mensajes de error
// @Description Estructura de respuesta para los errores
// @Accept json
// @Produce json
// @Failure 500 {object} ErrorResponse
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
