package entity

type Pedido struct {
	Id           int           `json:"pedido_id"`
	UserId       int           `json:"user_id"`
	Direccion    string        `json:"direccion"`
	Price        float64       `json:"price"`
	State        string        `json:"state"`
	Hamburguesas []Hamburguesa `json:"hamburguesas"`
	Hora         string        `json:"hora"`
	Fecha        string        `json:"fecha"`
}
