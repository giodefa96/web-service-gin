package repository

import (
	"example/web-service-gin/models"

	"gorm.io/gorm"
)

type CartItemRepository struct {
	DB *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) *CartItemRepository {
	return &CartItemRepository{DB: db}
}

func (r *CartItemRepository) CreateItem(cartItem *models.CartItem) error {
	return r.DB.Create(cartItem).Error
}
