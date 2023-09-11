package main

import (
	"context"
	"fmt"
	"log"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up a mongoDB client
	clientOpetions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOpetions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnet db after all operations
	defer client.Disconnect(context.TODO())
	fmt.Println("Connected to MongoDB")

	InsertIntoDb(client)
}
