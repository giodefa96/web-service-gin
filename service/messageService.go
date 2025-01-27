package service

import (
	"example/web-service-gin/amqp"
	"example/web-service-gin/dto"
	"log"
)

func PublishJob(body *dto.Body) error {
	err := amqp.PublishJob(body)
	if err != nil {
		return err
	}
	return nil
}

func StartListener() {
	if err := amqp.StartReceiver(); err != nil {
		log.Printf("Error starting receiver: %v", err)
	}
}
