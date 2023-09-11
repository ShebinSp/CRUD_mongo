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
	filter := bson.M{"name": "Balu"}

	// Delete the document
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	CheckError(err)

	// Print the result
	fmt.Printf("Delected %v documents", deleteResult.DeletedCount)

}

// Delete multiple entries from collection
func DeleteManyCol(client *mongo.Client){
	collection := InitCollection(client)
	ctx := context.TODO()

	// Define a filter
	filter := bson.M{"domain": "Python"}

	// Perform deletion
	deleteResult, err := collection.DeleteMany(ctx, filter)
	CheckError(err)

	fmt.Printf("\nDeleted %v documents\n", deleteResult.DeletedCount)
}
