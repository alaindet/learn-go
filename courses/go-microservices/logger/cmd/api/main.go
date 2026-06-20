package main

import (
	"common/json"
	"context"
	"log"
	"time"
	"logger/data"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	grpcPort = "50001"

	mongoUsername = "admin"
	mongoPassword = "password"
)

var client *mongo.Client

type App struct {
	Models data.Models
	json.HTTPClient
}

func main() {
	// Connect to Mongo
	mongoClient, err := connectoToMongoDB()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}

	app := App{
		Models: data.New(client),
	}

	go app.serve()
}

func (app *App) serve() {
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func connectoToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		UserName: mongoUsername,
		Password: mongoPassword,
	})

	connection, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	return connection, nil
}
