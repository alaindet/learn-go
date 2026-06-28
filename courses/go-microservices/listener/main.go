package main

import (
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

	// Start listening for messages from RabbitMQ
	// Create a consumer
	// Watch the queue and consume events
}
