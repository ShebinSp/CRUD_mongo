package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateInCollection(client *mongo.Client) {
	// Access your collection in database
	collection := InitCollection(client)
	filter := bson.M{"name": "Shebin Sp"}

	// Define an update
	update := bson.M{"$set": bson.M{"email": "shebinsp@newemail.com"}}

	// Perform an update
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	CheckError(err)

	// Print the result
	fmt.Printf("Matched %v documents and modified  %v documentns\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
