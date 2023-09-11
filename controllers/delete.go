package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteFromCollection(client *mongo.Client) {
	// Access your collection in DB
	collection := InitCollection(client)

	// Define a filter for the deletion
	filter := bson.M{"name": "Shebin Sp"}

	// Delete the document
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	CheckError(err)

	// Print the result
	fmt.Printf("Delected %v documents", deleteResult.DeletedCount)

}
