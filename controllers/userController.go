package controllers

import (
	"example/web-service-gin/database"
	"example/web-service-gin/models"
	"log"
	"net/http"
	"os"
	"time"

	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Struttura richiesta Login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserProfile gestisce il profilo utente protetto
func GetUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Benvenuto nel profilo!"})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&user)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Registrazione
func Register(c *gin.Context) {
	var user models.User

	// Legge i dati JSON e li associa a user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Controlla se l'email è già registrata
	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "L'email è già registrata"})
		return
	}

	// Hash della password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore nella creazione della password"})
		return
	}

	// Salvataggio dell'utente nel database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante la registrazione"})
		return
	}

	// Successo
	c.JSON(http.StatusOK, gin.H{"message": "Utente registrato con successo"})
}

// Login
func Login(c *gin.Context) {
	var req LoginRequest
	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cerca l'utente nel database
	database.DB.Where("email = ?", req.Email).First(&user)

	// Verifica credenziali
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email non registrata"})
		return
	}

	if !user.CheckPassword(req.Password) {
		log.Println("Errore: Password non corrispondente!")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password errata"})
		return
	}

	// Generazione del token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore nella generazione del token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
