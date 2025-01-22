package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint       `json:"user_id" gorm:"unique"`
	CartItems []CartItem `json:"cart_items" gorm:"constraint:OnDelete:CASCADE;"`
}

type CartItem struct {
	gorm.Model
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity" gorm:"not null;default:1"`
}
