package services

import (
	"ApiRestaurant/src/entity"
	"database/sql"
	"encoding/json"
	"errors"
)

type PedidoServices struct {
	DB *sql.DB
}

func NewPedidoServices(DB *sql.DB) *PedidoServices {
	return &PedidoServices{DB}
}

func (ps *PedidoServices) GetAllPedidos() (*[]entity.Pedido, error) {
	if ps.DB == nil {
		return nil, errors.New("la conexion a la base de datos es nil")
	}

	rows, err := ps.DB.Query("select * from pedidos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pedidos []entity.Pedido
	for rows.Next() {
		var pedido entity.Pedido
		var hamburgesaJSON []byte
		var horaStr string
		var fechaStr string

		if err := rows.Scan(&pedido.Id, &pedido.UserId, &pedido.Direccion, &pedido.Price, &pedido.State, &hamburgesaJSON, &horaStr, &fechaStr); err != nil {
			return nil, err
		}

		pedido.Hora = horaStr
		pedido.Fecha = fechaStr

		var hamburguesas []entity.Hamburguesa
		if err := json.Unmarshal(hamburgesaJSON, &hamburguesas); err != nil {
			return nil, err
		}
		pedido.Hamburguesas = hamburguesas
		pedidos = append(pedidos, pedido)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &pedidos, nil
}

func (ps *PedidoServices) GetById(id int) (*entity.Pedido, error) {
	if ps.DB == nil {
		return nil, errors.New("la conexion a la base de datos es nil")
	}

	var pedido entity.Pedido
	var hamburgesaJSON []byte
	var horaStr string
	var fechaStr string

	err := ps.DB.QueryRow("SELECT * FROM pedidos WHERE pedido_id = ?", id).Scan(
		&pedido.Id,
		&pedido.UserId,
		&pedido.Direccion,
		&pedido.Price,
		&pedido.State,
		&hamburgesaJSON,
		&horaStr,
		&fechaStr,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Pedido not found")
		}
		return nil, err
	}

	// Asignar hora y fecha como strings
	pedido.Hora = horaStr
	pedido.Fecha = fechaStr

	// Deserializar hamburguesas JSON en []Hamburguesa
	var hamburguesas []entity.Hamburguesa
	if err := json.Unmarshal(hamburgesaJSON, &hamburguesas); err != nil {
		return nil, err
	}

	pedido.Hamburguesas = hamburguesas

	return &pedido, nil
}
