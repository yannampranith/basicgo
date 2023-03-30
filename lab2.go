package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sample struct {
	ID       string `bson:"_id"`
	movie    string `bson:"movie"`
	diary string `bson:"diary"`
	title int    `bson:"title"`
}

func main() {
	// Set up a client to connect to the MongoDB server
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	
	collection := client.Database("").Collection("Sample")

	// Find a document that matches the given filter
	filter := bson.M{"_id": "1003530"}
	var result Sample
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// Print the result
	fmt.Printf("ID: %s\nName: %s\nSummary: %s\nBedrooms: %d\n", result.ID, result.Name, result.Summary, result.Bedrooms)
}
