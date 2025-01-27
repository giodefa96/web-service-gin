package service

import (
	"example/web-service-gin/amqp"
	"log"
)

func PublishJob() error {
	err := amqp.PublishJob()
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
