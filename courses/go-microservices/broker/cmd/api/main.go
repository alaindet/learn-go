package main

import (
	"common/json"
	"fmt"
	"log"
	"net/http"

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

	app := App{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
