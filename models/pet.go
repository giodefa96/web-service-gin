package models

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name         string  `json:"pet_name" gorm:"size:100;not null"`
	Species      string  `json:"species" gorm:"size:100;not null"`
	Breed        string  `json:"breed" gorm:"size:100;not null"`
	DayOfBirth   int     `json:"day_of_birth" gorm:"not null"`
	MonthOfBirth int     `json:"month_of_birth" gorm:"not null"`
	YearOfBirth  int     `json:"year_of_birth" gorm:"not null"`
	Weight       float64 `json:"weight" gorm:"not null"`       // Correzione di "weitgh"
	UserID       uint    `json:"user_id"`                      // Chiave esterna
	User         User    `gorm:"constraint:OnDelete:CASCADE;"` // Relazione con User
}
