package entity

import "time"

type Pedido struct {
	Id           int           `json:"pedido_id,omitempty"`
	UserId       int           `json:"user_id"`
	Direccion    string        `json:"direccion"`
	Price        float32       `json:"price,omitempty" `
	State        string        `json:"state,omitempty"`
	Hamburguesas []Hamburguesa `json:"hamburguesas"`
	Hora         time.Time     `json:"hora,omitempty"`
	Fecha        time.Time     `json:"fecha,omitempty"`
}
