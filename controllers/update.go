package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	fmt.Printf("\nMatched %v documents and modified  %v documentns\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}


func UpdateManyInCol(client *mongo.Client){
	collection := InitCollection(client)
	ctx := context.TODO()
	
	// Define a filter for the update
	filter := bson.M{"domain": "not confirmed"}

	// Define an update
	update := bson.M{"$set": bson.M{"domain": "Go"}}

	// Specify options for the operation
	updateOperations := options.Update()
	// Perform the update for all matchting documents
	updateResult, err := collection.UpdateMany(ctx, filter, update, updateOperations)
	CheckError(err)
	fmt.Printf("\nMatched %v documents and modified %v documents\n",updateResult.MatchedCount,updateResult.ModifiedCount)
}
