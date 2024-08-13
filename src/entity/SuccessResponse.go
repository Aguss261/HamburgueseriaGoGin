package entity

// SuccessResponse es la estructura para las respuestas exitosas
// @Description Estructura de respuesta para solicitudes exitosas
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
type SuccessResponse struct {
	Data interface{} `json:"data"`
}
