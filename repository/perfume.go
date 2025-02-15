package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/GilangAndhika/elfume/config"
	"github.com/GilangAndhika/elfume/model"
	"go.mongodb.org/mongo-driver/bson"
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
