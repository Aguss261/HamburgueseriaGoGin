package services

import (
	"ApiRestaurant/src/entity"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	DB *sql.DB
}

func NewUserServices(DB *sql.DB) *UserService {
	return &UserService{DB}
}

func (us *UserService) CreateUser(username, email, password, direccion string) (int, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	tx, err := us.DB.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO usuario (username, email, password_hash, direccion) VALUES (?, ?, ?, ?)",
		username, email, passwordHash, direccion)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec("INSERT INTO usuario_roles (usuario_id, rol_id) VALUES (?, ?)", userID, 9)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(userID), nil
}

func (us *UserService) GetUserByUsername(username string) (*entity.User, error) {
	row := us.DB.QueryRow(`
        SELECT u.id, u.username, u.email, u.password_hash, u.direccion, ur.rol_id
        FROM usuario u
        JOIN usuario_roles ur ON u.id = ur.usuario_id
        WHERE u.username = ?`, username)

	var user entity.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Direccion, &user.RolId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) CreateAdmin() error {
	var count int
	err := us.DB.QueryRow("SELECT COUNT(*) FROM usuario WHERE username = 'admin'").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Admin user already exists.")
		return nil
	}

	tx, err := us.DB.Begin()
	if err != nil {
		return err
	}

	// Insertar usuario admin
	_, err = tx.Exec(`
        INSERT INTO usuario (username, email, password_hash, direccion)
        VALUES ('admin', 'admin@example.com', 'admin', 'casa admin')`)
	if err != nil {
		tx.Rollback()
		return err
	}

	var roleID int
	err = us.DB.QueryRow("SELECT id FROM roles WHERE nombre = 'admin'").Scan(&roleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	var userID int
	err = us.DB.QueryRow("SELECT id FROM usuario WHERE username = 'admin'").Scan(&userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO usuario_roles (usuario_id, rol_id) VALUES (?, ?)", userID, roleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetUserByID obtiene un usuario por ID
func (us *UserService) GetUserByID(userID int) (*entity.User, error) {
	row := us.DB.QueryRow(`
        SELECT u.id, u.username, u.email, u.password_hash, u.direccion, ur.rol_id
        FROM usuario u
        JOIN usuario_roles ur ON u.id = ur.usuario_id
        WHERE u.id = ?`, userID)

	var user entity.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Direccion, &user.RolId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUsers obtiene todos los usuarios
func (us *UserService) GetUsers() ([]entity.User, error) {
	rows, err := us.DB.Query(`
        SELECT u.id, u.username, u.email, u.password_hash, u.direccion, ur.rol_id
        FROM usuario u
        JOIN usuario_roles ur ON u.id = ur.usuario_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Direccion, &user.RolId); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// CloseConnection cierra la conexión a la base de datos
func (us *UserService) CloseConnection() {
	if err := us.DB.Close(); err != nil {
		fmt.Println("Error closing the connection:", err)
	} else {
		fmt.Println("MySQL connection closed.")
	}
}
