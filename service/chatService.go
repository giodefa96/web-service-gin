package service

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/mistral"
)

// Recupera la chiave API di Mistral dall'ambiente
var mistralAPIKey = os.Getenv("MISTRAL_API_KEY")

// ChatService genera una risposta basata su un messaggio ricevuto
func ChatCompletion(message string) (string, error) {
	// Inizializza il modello Mistral
	llm, err := mistral.New(mistral.WithModel("mistral-large-latest"), mistral.WithAPIKey(mistralAPIKey))
	if err != nil {
		log.Println("Errore nella creazione del modello Mistral:", err)
		return "", err
	}

	ctx := context.Background()

	// Genera una risposta basata sul messaggio ricevuto
	response, err := llms.GenerateFromSinglePrompt(ctx, llm, message,
		llms.WithTemperature(0.8),
		llms.WithModel("mistral-large-latest"),
	)
	if err != nil {
		log.Println("Errore nella generazione della risposta:", err)
		return "", err
	}

	// Ritorna la risposta generata
	return response, nil
}

func ChatStream(message string) (chan string, error) {
	// Inizializza il canale di risposta
	responseChannel := make(chan string)

	// Inizializza il modello Mistral
	llm, err := mistral.New(mistral.WithModel("mistral-large-latest"), mistral.WithAPIKey(mistralAPIKey))
	if err != nil {
		log.Println("Errore nella creazione del modello Mistral:", err)
		return nil, err
	}

	ctx := context.Background()

	go func() {
		defer close(responseChannel) // Chiude il canale una volta terminato

		_, err := llms.GenerateFromSinglePrompt(ctx, llm, message,
			llms.WithTemperature(0.8),
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				responseChannel <- string(chunk) // Invia i dati nel canale
				return nil
			}),
		)

		if err != nil {
			log.Println("Errore nella generazione della risposta:", err)
		}
	}()

	return responseChannel, nil
}

// func ChatStream(message string) <-chan string {
// 	responseChannel := make(chan string)

// 	go func() {
// 		defer close(responseChannel) // Chiude il canale al termine dello streaming

// 		ctx := context.Background()
// 		llm, err := mistral.New(mistral.WithModel("open-mistral-7b"), mistral.WithAPIKey(mistralAPIKey))
// 		if err != nil {
// 			log.Println("Errore nell'inizializzazione del modello:", err)
// 			return
// 		}

// 		// Avvia la generazione della risposta con streaming
// 		_, err = llms.GenerateFromSinglePrompt(ctx, llm, message,
// 			llms.WithTemperature(0.8),
// 			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
// 				responseChannel <- string(chunk) // Invia il chunk al canale
// 				return nil
// 			}),
// 		)

// 		if err != nil {
// 			log.Println("Errore durante lo streaming:", err)
// 		}
// 	}()

// 	return responseChannel
// }
