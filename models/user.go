package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:100;not null"`
	Email    string `json:"email" gorm:"size:100;unique;not null"`
	Password string `json:"-" gorm:"type:varchar(120);not null"`
}
