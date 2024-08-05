package services

import (
	"ApiRestaurant/src/entity"
	"ApiRestaurant/src/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
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

		err := ps.parseTime(&pedido, horaStr, fechaStr)
		if err != nil {
			return nil, err
		}

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

func (ps *PedidoServices) GetById(id int, userId int) (*entity.Pedido, error) {
	if ps.DB == nil {
		return nil, errors.New("la conexion a la base de datos es nil")
	}
	var pedido entity.Pedido
	var hamburgesaJSON []byte
	var horaStr string
	var fechaStr string
	rolId := utils.GetUserRoles(userId, ps.DB)
	var err error
	if rolId == 8 {
		err = ps.DB.QueryRow("SELECT * FROM pedidos WHERE pedido_id = ?", id).Scan(&pedido.Id, &pedido.UserId, &pedido.Direccion, &pedido.Price, &pedido.State, &hamburgesaJSON, &horaStr, &fechaStr)
	} else {
		err = ps.DB.QueryRow("SELECT * FROM pedidos WHERE pedido_id = ? AND user_id = ?", id, userId).Scan(&pedido.Id, &pedido.UserId, &pedido.Direccion, &pedido.Price, &pedido.State, &hamburgesaJSON, &horaStr, &fechaStr)
		if pedido.UserId != userId {
			return nil, errors.New("El pedido no pertenece a este usuario")
		}
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Pedido not found")
		}
		return nil, err
	}

	err = ps.parseTime(&pedido, horaStr, fechaStr)
	if err != nil {
		return nil, err
	}

	var hamburguesas []entity.Hamburguesa
	if err := json.Unmarshal(hamburgesaJSON, &hamburguesas); err != nil {
		return nil, err
	}

	pedido.Hamburguesas = hamburguesas

	return &pedido, nil
}

func (ps *PedidoServices) GetByUserId(user_id int) (*[]entity.Pedido, error) {
	if ps.DB == nil {
		return nil, errors.New("la conexion a la base de datos es nil")
	}

	rows, err := ps.DB.Query("select * from pedidos where user_id = ?", user_id)
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

		err := ps.parseTime(&pedido, horaStr, fechaStr)
		if err != nil {
			return nil, err
		}

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

func (ps *PedidoServices) GetByFecha(fecha string) (*[]entity.Pedido, error) {
	if ps.DB == nil {
		return nil, errors.New("la conexion a la base de datos es nil")
	}

	rows, err := ps.DB.Query("select * from pedidos where fecha = ?", fecha)
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

		err := ps.parseTime(&pedido, horaStr, fechaStr)
		if err != nil {
			return nil, err
		}

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

func (ps *PedidoServices) CreatePedido(pedido *entity.Pedido) error {
	if ps.DB == nil {
		return errors.New("la conexion a la base de datos es nil")
	}
	hamburguesaJson, err := json.Marshal(pedido.Hamburguesas)
	if err != nil {
		return err
	}
	_, err = ps.DB.Exec(
		"INSERT INTO pedidos (user_id, direccion, price, state, hamburguesas, hora, fecha) VALUES (?, ?, ?, ?, ?, ? ,?)",
		pedido.UserId,
		pedido.Direccion,
		pedido.Price,
		pedido.State,
		hamburguesaJson,
		pedido.Hora,
		pedido.Fecha)

	if err != nil {
		return err
	}
	return nil
}

func (ps *PedidoServices) DeletePedido(id int, userId int) error {
	if ps.DB == nil {
		return errors.New("la conexion a la base de datos es nil")
	}

	var err error
	var result sql.Result
	roleId := utils.GetUserRoles(userId, ps.DB)
	if roleId == 8 {
		result, err = ps.DB.Exec("DELETE FROM pedidos WHERE pedido_id = ?", id)
	} else {
		result, err = ps.DB.Exec("DELETE FROM pedidos WHERE pedido_id = ? and user_id = ?", id, userId)
	}
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("El pedido no pudo ser borrado")
	}

	return nil
}

func (ps *PedidoServices) UpdatePedido(id int, pedido entity.Pedido, userId int) error {
	if ps.DB == nil {
		return errors.New("la conexion a la base de datos es nil")
	}
	existingPedido, _ := ps.GetById(id, userId)
	if existingPedido == nil {
		err := fmt.Errorf("Pedido no encontrado con el id %d", id)
		return err
	}
	hamburguesasJson, err := json.Marshal(pedido.Hamburguesas)
	if err != nil {
		return err
	}
	_, err1 := ps.DB.Exec("UPDATE pedidos SET direccion = ?, price = ?, state = ?, hamburguesas = ? ", pedido.Direccion, pedido.Price, pedido.State, hamburguesasJson)
	if err1 != nil {
		return err1
	}
	return nil
}

func (ps *PedidoServices) parseTime(pedido *entity.Pedido, horaStr string, fechaStr string) error {

	hora, err := time.Parse("15:04:05", horaStr)
	if err != nil {
		return err
	}
	pedido.Hora = hora

	fecha, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		return err
	}
	pedido.Fecha = fecha
	return nil
}
