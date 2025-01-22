package main

import (
	"log"

	"example/web-service-gin/controllers"
	"example/web-service-gin/database"
	"example/web-service-gin/repository"
	"example/web-service-gin/routes"
	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carica variabili d'ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore nel caricamento del file .env")
	}

	// Connessione al database
	database.ConnectDatabase()

	// Inizializza il router
	r := gin.Default()

	// ✅ Usa database.DB invece di db.DB
	cartRepo := repository.NewCartRepository(database.DB) // Passa il database
	userRepo := repository.NewUserRepository(database.DB) // Passa il database

	// Crea il servizio del carrello
	cartService := service.NewCartService(cartRepo) // Passa il repository
	userService := service.NewUserService(userRepo) // Passa il repository

	// Crea il gestore del carrello
	cartHandler := controllers.NewCartHandler(cartService) // Passa il servizio
	userHanlder := controllers.NewUserHandler(userService) // Passa il servizio

	// Imposta le rotte (Frontend + API)
	routes.SetupRouter(r, cartHandler, userHanlder)

	// Avvia il server
	r.Run(":8080")
}
