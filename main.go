package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up a client to connect to the MongoDB server
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Select the database and collection
	collection := client.Database("airbnb").Collection("listings")

	// Define the query selectors for each collector
	collectors := []struct {
		name       string
		selector   string
		projection bson.M
	}{
		{
			name:     "Name",
			selector: "Enter a listing name: ",
			projection: bson.M{
				"name": 1,
			},
		},
		{
			name:     "Neighborhood",
			selector: "Enter a neighborhood name: ",
			projection: bson.M{
				"name":                 1,
				"neighborhood_overview": 1,
			},
		},
		{
			name:     "Property Type",
			selector: "Enter a property type: ",
			projection: bson.M{
				"name":          1,
				"property_type": 1,
			},
		},
		{
			name:     "Minimum Nights",
			selector: "Enter a minimum number of nights: ",
			projection: bson.M{
				"name":           1,
				"minimum_nights": 1,
			},
		},
		{
			name:     "Maximum Nights",
			selector: "Enter a maximum number of nights: ",
			projection: bson.M{
				"name":           1,
				"maximum_nights": 1,
			},
		},
	}

	// Ask the user for input and execute the corresponding query
	for _, c := range collectors {
		var input string
		fmt.Print(c.selector)
		fmt.Scanln(&input)

		filter := bson.M{}
		switch c.name {
		case "Name":
			filter["name"] = bson.M{"$regex": primitive.Regex{Pattern: input, Options: "i"}}
		case "Neighborhood":
			filter["neighborhood"] = bson.M{"$regex": primitive.Regex{Pattern: input, Options: "i"}}
		case "Property Type":
			filter["property_type"] = bson.M{"$regex": primitive.Regex{Pattern: input, Options: "i"}}
		case "Minimum Nights":
			filter["minimum_nights"] = bson.M{"$gte": input}
		case "Maximum Nights":
			filter["maximum_nights"] = bson.M{"$lte": input}
		default:
			continue
		}

		cursor, err := collection.Find(context.Background(), filter, options.Find().SetProjection(c.projection))
		if err != nil {
			log.Fatal(err)
		}

		// Print the results
		var results []bson.M
		if err = cursor.All(context.Background(), &results); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Results for %s: %d\n", c.name, len(results))
		for _, result := range results {
			fmt.Printf("Name: %s\nSummary: %s\nNeighborhood Overview: %s\nTransit: %s\nProperty Type: %s\n",
			result["name"], result["summary"], result["neighborhood_overview"], result["transit"], result["property_type"])
	}
	fmt.Println()
}

// Close the cursor
err = cursor.Close(context.Background())
if err != nil {
	log.Fatal(err)
}
}

func getQuerySelector(name string) bson.M {
switch name {
case "Name":
return bson.M{"name": bson.M{"$regex": ".*"}}
case "Private Room":
return bson.M{"room_type": "Private room"}
case "Entire Home":
return bson.M{"room_type": "Entire home/apt"}
case "Price Range":
return bson.M{"price": bson.M{"$lt": 100}}
case "Superhost":
return bson.M{"host.is_superhost": true}
default:
return bson.M{}
}
}