package service

import (
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/mistral"
	"github.com/tmc/langchaingo/textsplitter"
)

// Recupera la chiave API di Mistral dall'ambiente
var mistralAPIKey = os.Getenv("MISTRAL_API_KEY")

// ChatService genera una risposta basata su un messaggio ricevuto
func ChatCompletion(message string) (string, error) {
	// Inizializza il modello Mistral
	llm, err := mistral.New(mistral.WithModel("mistral-large-latest"), mistral.WithAPIKey(mistralAPIKey))
	if err != nil {
		log.Println("Errore nella creazione del modello Mistral:", err)
		return "", errors.New("failed to initialize Mistral model")
	}

	ctx := context.Background()

	// Genera una risposta basata sul messaggio ricevuto
	response, err := llms.GenerateFromSinglePrompt(ctx, llm, message,
		llms.WithTemperature(0.8),
		llms.WithModel("mistral-large-latest"),
	)
	if err != nil {
		log.Println("Errore nella generazione della risposta:", err)
		return "", err // Ritorna l'errore direttamente
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
func SummarizeChatCompletion(file *multipart.FileHeader) (string, error) {
	// Implementazione della funzione
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Leggi il contenuto del file
	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	text := string(content)
	ctx := context.Background()
	// Inizializza il modello Mistral
	llm, err := mistral.New(mistral.WithModel("mistral-large-latest"), mistral.WithAPIKey(mistralAPIKey))
	if err != nil {
		log.Println("Errore nella creazione del modello Mistral:", err)
		return "", err
	}

	llmSummarizationChain := chains.LoadRefineSummarization(llm)
	docs := documentloaders.NewText(strings.NewReader(text))
	docsNew, err := docs.LoadAndSplit(ctx,
		textsplitter.NewRecursiveCharacter(),
	)
	if err != nil {
		return "", err
	}
	outputValues, err := chains.Call(ctx, llmSummarizationChain, map[string]any{"input_documents": docsNew})
	if err != nil {
		return "", err
	}
	out := outputValues["text"].(string)
	return out, nil
}
