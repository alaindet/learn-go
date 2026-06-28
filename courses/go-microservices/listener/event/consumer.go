package event

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

const logUrl = "http://logger/log"

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	if err := consumer.setup(); err != nil {
		return Consumer{}, nil
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(ch)
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		err := ch.QueueBind(
			queue.Name,   // name
			topic,        // key
			"logs_topic", // exchange
			false,        // no wait
			nil,          // arguments
		)
		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(
		queue.Name, // name
		"",         // consumer
		true,       // auto-acknowledge
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for message := range messages {
			var payload Payload
			_ = json.Unmarshal(message.Body, &payload)
			go handlePayload(payload)
		}
	}()

	log.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", queue.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		if err := logEvent(payload); err != nil {
			log.Println(err)
		}
	case "auth":
		log.Println("TODO: authenticate")
	default:
		if err := logEvent(payload); err != nil {
			log.Println(err)
		}
	}
}

func logEvent(entry Payload) error {
	// Convert auth payload to JSON
	jsonReq, err := json.MarshalIndent(entry, "", "\t")
	if err != nil {
		return err
	}

	// Build a direct HTTP request
	req, err := http.NewRequest("POST", logUrl, bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Any other HTTP error goes here
	if res.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}
