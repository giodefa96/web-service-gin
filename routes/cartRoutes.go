package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAlbumRoutes configura le rotte per gli album
func SetupCartRoutes(r *gin.Engine, cartHandler *controllers.CartHandler) {
	cartGroup := r.Group("/cart")
	{
		cartGroup.POST("/:user_id/:product_id/:quantity", cartHandler.AddItem) // ✅ Percorso corretto
	}
}
