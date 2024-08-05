package utils

import (
	"ApiRestaurant/src/entity"
	"database/sql"
)

func GetUserRoles(id int, db *sql.DB) int {

	row := db.QueryRow(`
        SELECT u.id, u.username, u.email, u.password_hash, u.direccion, ur.rol_id
        FROM usuario u
        JOIN usuario_roles ur ON u.id = ur.usuario_id
        WHERE u.id = ?`, id)

	var user entity.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Direccion, &user.RolId)
	if err != nil {
		return -0
	}
	return user.RolId

}
