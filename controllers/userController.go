package controllers

import (
	"net/http"
	"os"

	"example/web-service-gin/models"
	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler inizializza UserService e lo assegna al controller
func NewUserHandler() *UserHandler {
	userService := service.NewUserService()
	return &UserHandler{UserService: userService}
}

// CreateUser gestisce la creazione di un nuovo utente
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
