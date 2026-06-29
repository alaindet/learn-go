package main

import (
	"listener/event"
	"log"
	"os"
)

func main() {
	rabbitMQConn, err := connectToRabbitMQ()
	if err != nil {
		log.Panic("Cannot connect to RabbitMQ")
		os.Exit(1)
	}
	defer rabbitMQConn.Close()

	log.Println("Connected to RabbitMQ")

	consumer, err := event.NewConsumer(rabbitMQConn)
	if err != nil {
		log.Panic("Cannot create a RabbitMQ consumer")
		os.Exit(1)
	}

	topics := []string{"log.INFO", "log.WARNING", "log.ERROR"}

	if err := consumer.Listen(topics); err != nil {
		log.Println(err)
	}
}
