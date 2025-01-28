package routes

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configura le rotte per gli utenti
// SetupUserRoutes configura le rotte per gli utenti
func SetupUserRoutes(r *gin.Engine) {
	userHandler := controllers.NewUserHandler() // Creiamo un'istanza di UserHandler

	// Rotte di autenticazione
	r.POST("/CreateUser", userHandler.CreateUser) // Utilizziamo il metodo dell'istanza
	r.POST("/Login", userHandler.Login)

	// Rotte protette
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user/:id", userHandler.GetUserByID)

	}
}
