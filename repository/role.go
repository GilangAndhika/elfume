package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/GilangAndhika/elfume/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreateRole inserts a new role into the database
func CreateRole(ctx context.Context, db *mongo.Database, role *model.Role) error {
	// Connect to the database
	collection := db.Collection("roles")

	_, err := collection.InsertOne(ctx, role)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create role")
	}

	return nil
}
