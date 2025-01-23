package service

import (
	"errors"
	"example/web-service-gin/database"
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
)

// CartService gestisce la logica di business per il carrello
type CartService struct {
	CartRepo *repository.CartRepository
}

// NewCartService crea automaticamente un'istanza di CartRepository
func NewCartService() *CartService {
	// Usa direttamente database.DB senza ristanziarlo
	cartRepo := repository.NewCartRepository(database.DB)

	// Ritorna il servizio con il repository già inizializzato
	return &CartService{CartRepo: cartRepo}
}

// AddItemToCart aggiunge un prodotto al carrello dell'utente
func (s *CartService) AddItemToCart(userID uint, productID uint, quantity int) error {
	// Verifica che la quantità sia valida
	if quantity <= 0 {
		return errors.New("la quantità deve essere maggiore di zero")
	}

	// Ottiene il carrello dell'utente
	cart, err := s.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	// Crea un nuovo elemento nel carrello
	cartItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
	}

	// Aggiunge l'elemento al database
	return s.CartRepo.AddItem(&cartItem)
}
