package controllers

import (
	"net/http"
	"strconv"

	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartService *service.CartService
}

func NewCartHandler(cartService *service.CartService) *CartHandler {
	return &CartHandler{CartService: cartService}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	productID, _ := strconv.Atoi(c.Param("product_id"))
	quantity, _ := strconv.Atoi(c.Param("quantity"))

	err := h.CartService.AddItemToCart(uint(userID), uint(productID), quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}
