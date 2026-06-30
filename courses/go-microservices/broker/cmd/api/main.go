package main

import (
	"common/json"
	"fmt"
	"log"
	"net/http"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	webPort = "80"
)

type App struct {
	json.HTTPClient
	RabbitMQ *amqp.Connection
}

func main() {
	log.Printf("Starting broker service on port %s\n", webPort)

	rabbitMQConn, err := connectToRabbitMQ()
	if err != nil {
		log.Panic("Cannot connect to RabbitMQ")
		os.Exit(1)
	}
	defer rabbitMQConn.Close()
	log.Println("Connected to RabbitMQ")

	app := App{
		RabbitMQ: rabbitMQConn,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
