package service

import (
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
)

type CartService struct {
	CartRepo *repository.CartRepository
}

func NewCartService(cartRepo *repository.CartRepository) *CartService {
	return &CartService{CartRepo: cartRepo}
}

func (s *CartService) AddItemToCart(userID uint, productID uint, quantity int) error {
	cart, err := s.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	cartItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
	}

	return s.CartRepo.AddItem(&cartItem)
}
