package controllers

import (
	"example/web-service-gin/service"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Struttura per il corpo della richiesta
type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}

// Struttura per la risposta
type ChatResponse struct {
	Response string `json:"response"`
}

// Struttura per il corso della risposta file txt
type ChatResponseFile struct {
	Response string `json:"response"`
}

// ChatCompletion riceve un messaggio, lo elabora con Mistral e restituisce una risposta
func ChatCompletion(c *gin.Context) {
	var req MessageRequest

	// Effettua il binding del JSON ricevuto nella struttura
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or missing 'message' field"})
		return
	}

	// Chiama il service di chat per generare una risposta
	response, err := service.ChatCompletion(req.Message)
	if err != nil {
		// Determinare il codice HTTP basandosi sul messaggio di errore
		statusCode := http.StatusInternalServerError
		errorMessage := "An unexpected error occurred"

		if strings.Contains(err.Error(), "400 Bad Request") {
			statusCode = http.StatusBadRequest
			errorMessage = "Invalid request: check input message"
		} else if strings.Contains(err.Error(), "Failed to initialize Mistral model") {
			statusCode = http.StatusServiceUnavailable
			errorMessage = "Mistral model is currently unavailable"
		} else if strings.Contains(err.Error(), "Failed to generate response") {
			errorMessage = "Error generating response from AI model"
		}

		// Loggare l'errore dettagliato per debugging
		fmt.Println("Errore nel controller:", err)

		// Restituire la risposta con il codice HTTP adeguato
		c.JSON(statusCode, gin.H{"error": errorMessage})
		return
	}

	// Restituisce la risposta generata dal modello
	c.JSON(http.StatusOK, ChatResponse{Response: response})
}

// StreamChatCompletion gestisce lo streaming della risposta in tempo reale
func StreamChatCompletion(c *gin.Context) {
	var req MessageRequest

	// Effettua il binding del JSON ricevuto nella struttura
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or missing 'message' field"})
		return
	}

	// Imposta gli header per lo streaming
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.WriteHeader(http.StatusOK)

	// Ottieni il canale di streaming dal servizio
	responseChannel, err := service.ChatStream(req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante la generazione del messaggio"})
		return
	}

	// Assicura che il buffer venga flushato ad ogni chunk
	c.Writer.Flush()

	// Legge i dati dal canale e li invia al client
	for response := range responseChannel {
		_, err := fmt.Fprintf(c.Writer, "%s", response) // Scrive il chunk al client
		if err != nil {
			log.Println("Errore nell'invio della risposta:", err)
			break
		}
		c.Writer.Flush() // Invia immediatamente il chunk al client
	}

	log.Println("Stream completato")
}

func SummarizeChatCompletion(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Errore nel file upload"})
		return
	}
	// passa il file al servizio
	response, err := service.SummarizeChatCompletion(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore nella generazione del riassunto"})
		return
	}
	// Restituisce la risposta generata dal modello
	c.JSON(http.StatusOK, ChatResponse{Response: response})

}
