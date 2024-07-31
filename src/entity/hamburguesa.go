package entity

type Hamburguesa struct {
	Id          int         `json:"id,omitempty"`
	Nombre      string      `json:"nombre" binding:"required"`
	Price       float32     `json:"price" binding:"required,gt=0"`
	Descripcion string      `json:"descripcion" binding:"required"`
	ImgUrl      string      `json:"imgUrl" binding:"required"`
	Ingrediente Ingrediente `json:"ingredientes" binding:"required"`
}
