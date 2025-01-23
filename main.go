package main

import (
	"log"

	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carica variabili d'ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore nel caricamento del file .env")
	}

	// // Connessione al database
	// database.ConnectDatabase()

	// Inizializza il router
	r := gin.Default()

	// Imposta le rotte (Frontend + API)
	routes.SetupRouter(r)

	// Avvia il server
	r.Run(":8080")
}
