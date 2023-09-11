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
		"domain": "Go",
	}

	// Insert the document
	insertResult, err := collection.InsertOne(context.TODO(), document)
	CheckError(err)

	// Print the result
	fmt.Printf("Inserted ID: %v\n", insertResult.InsertedID)
}



// Creating many records into DB at once
func InsertMany(client *mongo.Client){
	// Get in touch with your collection
	collection := InitCollection(client)

	// Initializing context.TODO() to a variable
	ctx := context.TODO()

	documents := []interface{}{
		bson.M{"name": "Manaf", "email": "manaf@email.com", "domain": "Go"},
		bson.M{"name": "Junaid", "email": "junaid@email.com", "domain": "not confirmed"},
		bson.M{"name": "Sreerag", "email": "sreerag@email.com", "domain": "Go"},
		bson.M{"name": "Raheem", "email": "raheem@email.com", "domain": "not confirmed"},
		bson.M{"name": "Raju", "email": "raju@email.com", "domain": "Python"},
		bson.M{"name": "Radha", "email": "radha@email.com", "domain": "Python"},
		bson.M{"name": "Balu", "email": "balu@email.com", "domain": "PHP"},
	}

	// Insert multiple documents
	insertResult, err := collection.InsertMany(ctx, documents)
	CheckError(err)

	// Print the details
	fmt.Printf("\nInserted ID's: %v\n", insertResult.InsertedIDs...)
}