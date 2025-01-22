package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAlbumRoutes configura le rotte per gli album
func SetupAlbumRoutes(r *gin.Engine) {
	albumGroup := r.Group("/albums")
	{
		albumGroup.GET("/", controllers.GetAlbums)
		albumGroup.GET("/:id", controllers.GetAlbumByID)
		albumGroup.POST("/", controllers.CreateAlbum)
		albumGroup.PUT("/:id", controllers.UpdateAlbum)
		albumGroup.DELETE("/:id", controllers.DeleteAlbum)
	}
}
