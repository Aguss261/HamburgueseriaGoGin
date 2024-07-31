package entity

type Ingrediente struct {
	Huevo   int `json:"huevo" binding:"omitempty,gt=-1"`
	Lechuga int `json:"lechuga" binding:"omitempty,gt=-1"`
	Tomate  int `json:"tomate" binding:"omitempty,gt=-1"`
	Cebolla int `json:"cebolla" binding:"omitempty,gt=-1"`
	Bacon   int `json:"bacon" binding:"omitempty,gt=-1"`
	Pepino  int `json:"pepino" binding:"omitempty,gt=-1"`
}
