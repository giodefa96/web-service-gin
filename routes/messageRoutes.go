package routes

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(r *gin.Engine) {
	messageGroup := r.Group("/message")
	messageGroup.Use(middleware.AuthMiddleware())
	{
		messageGroup.POST("/", controllers.PostMessage)
		messageGroup.GET("/", controllers.StartListener)
	}
}
