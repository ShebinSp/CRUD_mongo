package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadFromCollection(client *mongo.Client) {
	// Access collection
	collection := InitCollection(client)
	// Define a filter for the query
	filter := bson.M{"name": "Shebin Sp"}

	// Find the document
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	CheckError(err)
	
	// Print the result
	fmt.Println("Your Search Result: ", result)
}
