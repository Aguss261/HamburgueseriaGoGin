package entity

// SuccessResponse es la estructura para las respuestas exitosas
// @Description Estructura de respuesta para solicitudes exitosas
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// ErrorResponse es la estructura para los mensajes de error
// @Description Estructura de respuesta para los errores
// @Accept json
// @Produce json
// @Failure 500 {object} ErrorResponse
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
