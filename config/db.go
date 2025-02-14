package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() {
	// Load .env file
	mongostr := os.Getenv("MONGOSTRING")
	if mongostr == "" {
		log.Fatal("Error loading mongo string")
	}

	// Create MongoDB client options
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongostr))
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to verify database connection", err)
	}

	// Set the global MongoClient
	MongoClient = client
}

// return mongo client instance
func GetDB() *mongo.Client {
	if MongoClient == nil {
		log.Fatal("MongoClient is not initialized. Connect to database first.")
	}
	return MongoClient
}