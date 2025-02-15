package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/GilangAndhika/elfume/config"
	"github.com/GilangAndhika/elfume/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePerfume creates a new perfume product and uploads its image to GitHub
func CreatePerfume(perfume *model.Perfume, imageBase64 string, fileName string) error {
	// Get database connection
	perfumeCollection := config.MongoDB.Collection("perfumes")

	// Upload image to GitHub
	err := UploadtoGithub(fileName, imageBase64)
	if err != nil {
		return fmt.Errorf("failed to upload image to GitHub: %v", err)
	}

	// Construct GitHub image URL
	githubOwner := os.Getenv("GITHUB_OWNER")
	githubRepo := os.Getenv("GITHUB_REPO")
	imageURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", githubOwner, githubRepo, fileName)

	// Assign ID and timestamps
	perfume.Image = imageURL

	// Insert perfume into the database
	_, err = perfumeCollection.InsertOne(context.TODO(), bson.M{
		"_id":         perfume.PerfumeID,
		"name":        perfume.Name,
		"brand":       perfume.Brand,
		"types":       perfume.Types,
		"categories":  perfume.Categories,
		"sizes":       perfume.Sizes,
		"image":       perfume.Image,
		"price":       perfume.Price,
		"description": perfume.Description,
		"stock":       perfume.Stock,
		"created_at":  perfume.CreatedAt,
		"updated_at":  perfume.UpdatedAt,
	})
	if err != nil {
		return fmt.Errorf("failed to insert perfume into database: %v", err)
	}

	return nil
}

// Get all perfumes from the database
func GetAllPerfumes() ([]model.Perfume, error) {
	// Get database connection
	perfumeCollection := config.MongoDB.Collection("perfumes")

	// Find all perfumes
	cursor, err := perfumeCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch perfumes: %v", err)
	}
	defer cursor.Close(context.Background())

	// Decode all perfumes
	var perfumes []model.Perfume
	if err = cursor.All(context.Background(), &perfumes); err != nil {
		return nil, fmt.Errorf("failed to decode perfumes: %v", err)
	}

	return perfumes, nil
}

// Get a perfume by ID from the database
func GetPerfumeByID(id string) (*model.Perfume, error) {
	// Convert string ID to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid perfume ID format: %v", err)
	}

	// Get database connection
	perfumeCollection := config.MongoDB.Collection("perfumes")

	// Find perfume by ID
	var perfume model.Perfume
	err = perfumeCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&perfume)
	if err != nil {
		return nil, fmt.Errorf("failed to find perfume: %v", err)
	}

	return &perfume, nil
}

// GetFilteredPerfumes retrieves perfumes with optional filters (e.g., by size, brand, category)
func GetFilteredPerfumes(filters map[string]string) ([]model.Perfume, error) {
	// Get database connection
	perfumeCollection := config.MongoDB.Collection("perfumes")

	// Build MongoDB query filter
	query := bson.M{}
	for key, value := range filters {
		query[key] = bson.M{"$regex": value, "$options": "i"} // Case-insensitive search
	}

	// Find perfumes using filter
	cursor, err := perfumeCollection.Find(context.TODO(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch perfumes: %v", err)
	}
	defer cursor.Close(context.Background())

	// Decode results
	var perfumes []model.Perfume
	if err = cursor.All(context.Background(), &perfumes); err != nil {
		return nil, fmt.Errorf("failed to decode perfumes: %v", err)
	}

	return perfumes, nil
}

// UpdatePerfume updates a perfume in the database
func UpdatePerfume(id string, updatedPerfume model.Perfume) error {
	// Convert ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid perfume ID format: %v", err)
	}

	// Get database connection
	perfumeCollection := config.MongoDB.Collection("perfumes")

	// Set updated timestamp
	updatedPerfume.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"name":        updatedPerfume.Name,
			"brand":       updatedPerfume.Brand,
			"types":       updatedPerfume.Types,
			"categories":  updatedPerfume.Categories,
			"sizes":       updatedPerfume.Sizes,
			"price":       updatedPerfume.Price,
			"description": updatedPerfume.Description,
			"stock":       updatedPerfume.Stock,
			"updated_at":  updatedPerfume.UpdatedAt,
		},
	}

	// Perform the update
	result, err := perfumeCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return fmt.Errorf("failed to update perfume: %v", err)
	}

	// Check if the perfume was found and modified
	if result.MatchedCount == 0 {
		return fmt.Errorf("perfume not found")
	}

	return nil
}

// DeletePerfume deletes a perfume by its ID
func DeletePerfume(id string) error {
	// Convert string ID to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid perfume ID format: %v", err)
	}

	// Get database connection
	perfumeCollection := config.MongoDB.Collection("perfumes")

	// Delete perfume by ID
	result, err := perfumeCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("failed to delete perfume: %v", err)
	}

	// Check if any document was actually deleted
	if result.DeletedCount == 0 {
		return fmt.Errorf("perfume not found")
	}

	return nil
}