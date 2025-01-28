package routes

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPetRoutes(r *gin.Engine) {
	petGroup := r.Group("/pet")
	petGroup.Use(middleware.AuthMiddleware())
	petHandler := controllers.NewPetHandler()
	{
		petGroup.POST("/", petHandler.CreatePet)
		petGroup.GET("/", petHandler.GetPets)
		petGroup.GET("/:id", petHandler.GetPetByID)
		petGroup.PUT("/:id", petHandler.UpdatePet)
		petGroup.DELETE("/:id", petHandler.DeletePet)
	}
}
