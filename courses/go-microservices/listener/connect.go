package main

import (
	"fmt"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	username = "guest"
	password = "guest"
	host     = "rabbitmq"
)

func connectToRabbitMQ() (*amqp.Connection, error) {
	var attempts int64
	backOff := 1 * time.Second
	var connection *amqp.Connection
	url := fmt.Sprintf("amqp://%s:%s@%s", username, password, host)

	for {
		c, err := amqp.Dial(url)
		if err != nil {
			log.Println("RabbitMQ not yet ready..")
			attempts++
		} else {
			connection = c
			break
		}

		if attempts > 5 {
			log.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(attempts), 2))
		log.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
