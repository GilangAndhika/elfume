package config

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	// Load .env file
	mongostr := os.Getenv("MONGOSTRING")
	if mongostr == "" {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongostr))
	if err != nil {
		log.Fatal(err)
	}

	// Set the global MongoClient
	MongoClient = client
}

func GetMongoClient() *mongo.Client {
	return MongoClient
}
