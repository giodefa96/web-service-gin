package controllers

import (
	"example/web-service-gin/database"
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	result := database.DB.Find(&albums)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	result := database.DB.First(&album, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func CreateAlbum(c *gin.Context) {
	var album models.Album
	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&album)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, album)
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	result := database.DB.First(&album, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = database.DB.Save(&album)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	result := database.DB.First(&album, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	result = database.DB.Delete(&album)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
