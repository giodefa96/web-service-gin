package database

import (
	"example/web-service-gin/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Errore di connessione al database:", err)
	}

	// Migrazioni automatiche
	err = DB.AutoMigrate(&models.User{}, &models.Album{})
	if err != nil {
		log.Fatal("Errore durante le migrazioni:", err)
	}

	fmt.Println("✅ Connessione al database con GORM avvenuta con successo")
}
