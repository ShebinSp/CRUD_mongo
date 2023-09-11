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

// Query multiple documents from collection
func ReadManyFromCol(client *mongo.Client) {
	collection := InitCollection(client)
	ctx := context.TODO()

	// To fetch all documents where Go is the domain
	var gophers []bson.M
	filter := bson.M{"domain": "Go"}
	cursoR, err := collection.Find(ctx, filter)
	CheckError(err)
	defer cursoR.Close(ctx)
	err = cursoR.All(ctx, &gophers)
	CheckError(err)
	fmt.Println("\nGophers: ")
	for _, res := range gophers {
		fmt.Println(res)
	}

	// Fetch all the documents. bson.D should be used to included all the documents
	cursor, err := collection.Find(ctx, bson.D{})

	CheckError(err)
	defer cursor.Close(ctx)

	var results bson.A
	err = cursor.All(ctx, &results)
	CheckError(err)

	fmt.Println("\nResult from Collection")
	for _, result := range results {
		fmt.Println(result)
	}
}
