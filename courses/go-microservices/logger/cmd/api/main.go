package main

import (
	"common/json"
	"context"
	"fmt"
	"log"
	"logger/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	webPort       = "80"
	rpcPort       = "5001"
	grpcPort      = "50001"
	mongoURL      = "mongodb://mongo:27017"
	mongoUsername = "admin"
	mongoPassword = "password"
)

var client *mongo.Client

type App struct {
	Models data.Models
	json.HTTPClient
}

func main() {
	log.Printf("Starting logger service on port %s\n", webPort)

	// Connect to Mongo
	mongoClient, err := connectToMongoDB()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to MongoDB")

	client = mongoClient
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := App{
		Models: data.New(client),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func connectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: mongoUsername,
		Password: mongoPassword,
	})

	connection, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	return connection, nil
}
