package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (server *RPCServer) LogInfo(payload RPCPayload, feedback *string) error {

	logEntry := data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	}

	collection := client.Database("logs").Collection("logs")

	if _, err := collection.InsertOne(context.TODO(), logEntry); err != nil {
		log.Println("error writing to MongoDB")
		return err
	}

	*feedback = fmt.Sprintf("Processed payload via RPC: %s", payload.Name)
	return nil
}
