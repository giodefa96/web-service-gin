package controllers

import (
	"example/web-service-gin/dto"
	"example/web-service-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostMessage(c *gin.Context) {
	var body dto.Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.PublishJob(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Ensure the function exits after sending the error response
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job published successfully"})
}

func StartListener(c *gin.Context) {
	go service.StartListener()
	c.JSON(http.StatusOK, gin.H{"message": "Listener started successfully"})
}
