package controllers

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// Access collections from Database
func InitCollection(client *mongo.Client) *mongo.Collection {
	// Assuming you have a collection named "users"
	collection := client.Database("mydb").Collection("users")
	return collection
}

// Fix the errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}