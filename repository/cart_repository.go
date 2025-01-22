package repository

import (
	"example/web-service-gin/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) GetCartByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.DB.Where("user_id = ?", userID).Preload("CartItems").First(&cart).Error
	return &cart, err
}

func (r *CartRepository) AddItem(cartItem *models.CartItem) error {
	return r.DB.Create(cartItem).Error
}

func (r *CartRepository) ClearCart(cartID uint) error {
	return r.DB.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}
