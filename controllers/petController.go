package controllers

import (
	"net/http"

	"example/web-service-gin/models"
	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	PetService *service.PetService
}

// NewPetHandler inizializza PetService e lo assegna al controller
func NewPetHandler() *PetHandler {
	petService := service.NewPetService()
	return &PetHandler{PetService: petService}
}

// GetPets gestisce la richiesta di tutti gli animali
func (h *PetHandler) GetPets(c *gin.Context) {
	pets, err := h.PetService.GetPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pets)
}

// GetPetByID gestisce la richiesta di un animale per ID
func (h *PetHandler) GetPetByID(c *gin.Context) {
	id := c.Param("id")
	pet, err := h.PetService.GetPetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}

// GetPetsByUserID gestisce la richiesta di tutti gli animali di un utente
func (h *PetHandler) GetPetsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	pets, err := h.PetService.GetPetsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pets)
}

// CreatePet gestisce la creazione di un nuovo animale
func (h *PetHandler) CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.BindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.PetService.CreatePet(&pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pet)
}

// UpdatePet gestisce l'aggiornamento di un animale esistente
func (h *PetHandler) UpdatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.BindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.PetService.UpdatePet(&pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}

// DeletePet gestisce l'eliminazione di un animale esistente
func (h *PetHandler) DeletePet(c *gin.Context) {
	var pet models.Pet
	if err := c.BindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.PetService.DeletePet(&pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully"})
}

// DeletePetByID gestisce l'eliminazione di un animale per ID
func (h *PetHandler) DeletePetByID(c *gin.Context) {
	id := c.Param("id")
	err := h.PetService.DeletePetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully"})
}

// DeletePetsByUserID gestisce l'eliminazione di tutti gli animali di un utente
func (h *PetHandler) DeletePetsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	err := h.PetService.DeletePetsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pets deleted successfully"})
}

// DeleteAllPets gestisce l'eliminazione di tutti gli animali
func (h *PetHandler) DeleteAllPets(c *gin.Context) {
	err := h.PetService.DeleteAllPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All pets deleted successfully"})
}
