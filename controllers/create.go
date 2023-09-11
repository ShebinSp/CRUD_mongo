package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertIntoCollection(client *mongo.Client) {
	// Access collection
	collection := InitCollection(client)

	// Create a document to insert
	document := bson.M{
		"name":  "Shebin Sp",
		"email": "shebin@email.com",
	}

	// Insert the document
	insertResult, err := collection.InsertOne(context.TODO(), document)
	CheckError(err)

	// Print the result
	fmt.Printf("Inserted ID: %v\n", insertResult.InsertedID)
}
