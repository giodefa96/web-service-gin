package controllers

import (
	"net/http"
	"strconv"

	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
)

// CartHandler gestisce le operazioni sul carrello
type CartHandler struct {
	CartService *service.CartService
}

// NewCartHandler inizializza CartService automaticamente
func NewCartHandler() *CartHandler {
	cartService := service.NewCartService()
	return &CartHandler{CartService: cartService}
}

// AddItem gestisce l'aggiunta di un prodotto al carrello
func (h *CartHandler) AddItem(c *gin.Context) {
	// Parsing dei parametri e gestione errori
	userID, err1 := strconv.Atoi(c.Param("user_id"))
	productID, err2 := strconv.Atoi(c.Param("product_id"))
	quantity, err3 := strconv.Atoi(c.Param("quantity"))

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters. Make sure user_id, product_id, and quantity are valid numbers."})
		return
	}

	// Verifica che la quantità sia valida
	if quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than zero."})
		return
	}

	// Aggiunge il prodotto al carrello
	err := h.CartService.AddItemToCart(uint(userID), uint(productID), quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully."})
}
