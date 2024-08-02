package services

import (
	"ApiRestaurant/src/entity"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type HamburguesaServices struct {
	DB *sql.DB
}

func NewHamburguesaService(DB *sql.DB) *HamburguesaServices {
	return &HamburguesaServices{DB}
}

func (hs *HamburguesaServices) GetAllHamburguesas() (*[]entity.Hamburguesa, error) {
	if hs.DB == nil {
		return nil, errors.New("database connection is nil")
	}

	rows, err := hs.DB.Query("SELECT * FROM hamburguesa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hamburguesas []entity.Hamburguesa
	for rows.Next() {
		var hamburguesa entity.Hamburguesa
		var ingredienteJSON []byte
		if err := rows.Scan(&hamburguesa.Id, &hamburguesa.Nombre, &hamburguesa.Price, &hamburguesa.Descripcion, &hamburguesa.ImgUrl, &ingredienteJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(ingredienteJSON, &hamburguesa.Ingrediente); err != nil {
			return nil, err
		}

		hamburguesas = append(hamburguesas, hamburguesa)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &hamburguesas, nil
}

func (hs *HamburguesaServices) GetById(id int) (*entity.Hamburguesa, error) {
	if hs.DB == nil {
		return nil, errors.New("database connection is nil")
	}

	var hamburguesa entity.Hamburguesa
	var ingredienteJSON []byte
	err := hs.DB.QueryRow("SELECT * FROM hamburguesa WHERE id = ?", id).Scan(&hamburguesa.Id, &hamburguesa.Nombre, &hamburguesa.Price, &hamburguesa.Descripcion, &hamburguesa.ImgUrl, &ingredienteJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("hamburguesa not found")
		}
		return nil, err
	}
	if err := json.Unmarshal(ingredienteJSON, &hamburguesa.Ingrediente); err != nil {
		return nil, err
	}
	return &hamburguesa, nil
}

func (s *HamburguesaServices) GetByName(name string) ([]entity.Hamburguesa, error) {
	rows, err := s.DB.Query("SELECT * FROM Hamburguesa WHERE nombre like ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hamburguesas []entity.Hamburguesa
	for rows.Next() {
		var hamburguesa entity.Hamburguesa
		var ingredienteJSON []byte
		if err := rows.Scan(&hamburguesa.Id, &hamburguesa.Nombre, &hamburguesa.Price, &hamburguesa.Descripcion, &hamburguesa.ImgUrl, &ingredienteJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(ingredienteJSON, &hamburguesa.Ingrediente); err != nil {
			return nil, err
		}
		hamburguesas = append(hamburguesas, hamburguesa)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return hamburguesas, nil
}

func (s *HamburguesaServices) GetByPrice(price float64) ([]entity.Hamburguesa, error) {
	rows, err := s.DB.Query("SELECT * FROM Hamburguesa WHERE price > ?", price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hamburguesas []entity.Hamburguesa
	for rows.Next() {
		var hamburguesa entity.Hamburguesa
		var ingredienteJSON []byte
		if err := rows.Scan(&hamburguesa.Id, &hamburguesa.Nombre, &hamburguesa.Price, &hamburguesa.Descripcion, &hamburguesa.ImgUrl, &ingredienteJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(ingredienteJSON, &hamburguesa.Ingrediente); err != nil {
			return nil, err
		}
		hamburguesas = append(hamburguesas, hamburguesa)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return hamburguesas, nil
}

func (hs *HamburguesaServices) CreateHamburguesa(hamburguesa *entity.Hamburguesa) error {
	ingredienteJSON, err := json.Marshal(hamburguesa.Ingrediente)
	if err != nil {
		return err
	}
	_, err = hs.DB.Exec(
		"INSERT INTO hamburguesa (nombre, price, descripcion, imgUrl, ingredientes) VALUES (?, ?, ?, ?, ?)",
		hamburguesa.Nombre,
		hamburguesa.Price,
		hamburguesa.Descripcion,
		hamburguesa.ImgUrl,
		ingredienteJSON,
	)
	if err != nil {
		return err
	}
	return nil
}

func (hs *HamburguesaServices) DeleteHamburguesaByiD(id int) error {
	result, err := hs.DB.Exec("DELETE FROM hamburguesa WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ningúna hamburguesa encontrada con el id %d", id)
	}
	return nil
}

func (hs *HamburguesaServices) EditHamburguesa(id int, hamburguesa entity.Hamburguesa) error {
	existingHamburguesa, _ := hs.GetById(id)
	if existingHamburguesa == nil {
		return errors.New("hamburguesa not found")
	}
	ingredienteJSON, err1 := json.Marshal(hamburguesa.Ingrediente)
	if err1 != nil {
		return err1
	}
	_, err := hs.DB.Exec("UPDATE hamburguesa SET nombre = ?, price = ?, descripcion = ?, imgUrl = ?, ingredientes = ? WHERE id = ?", hamburguesa.Nombre, hamburguesa.Price, hamburguesa.Descripcion, hamburguesa.ImgUrl, ingredienteJSON, id)
	if err != nil {
		return err
	}
	return nil
}
