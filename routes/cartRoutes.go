package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

// SetupCartRoutes configura le rotte per il carrello
func SetupCartRoutes(r *gin.Engine) {
	cartHandler := controllers.NewCartHandler() // Creiamo un'istanza di CartHandler

	cartGroup := r.Group("/cart")
	{
		cartGroup.POST("/:user_id/:product_id/:quantity", cartHandler.AddItem) // Usiamo il metodo dell'istanza
	}
}
