package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:100;not null"`
	Email    string `json:"email" gorm:"size:100;unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"` // Assicuriamo spazio sufficiente
}

// Funzione per hashare la password
func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// Funzione per verificare la password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
