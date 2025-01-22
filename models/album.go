package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model        // Include campi predefiniti come ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `json:"title" gorm:"size:100;not null"`
	Artist     string `json:"artist" gorm:"size:100;unique;not null"`
	Price      string `json:"Price" gorm:"size:100;unique;not null"`
}
