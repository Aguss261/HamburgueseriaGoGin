package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Username     string    `json:"Username"`
	Email        string    `json:"Email"`
	PasswordHash string    `json:"-"`
	Direccion    string    `json:"Direccion"`
	RolId        int       `json:"Rol_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword verifies the provided password against the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
