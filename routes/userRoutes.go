package routes

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configura le rotte per gli utenti
func SetupUserRoutes(r *gin.Engine, userHandler *controllers.UserHandler) {
	// Rotte di autenticazione
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/CreateUser", userHandler.CreateUser)
	// Rotte protette
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", controllers.GetUserProfile) // supponiamo tu abbia una funzione `GetUserProfile`
}
