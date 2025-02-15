package controller

import (
	"encoding/base64"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/GilangAndhika/elfume/model"
	"github.com/GilangAndhika/elfume/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePerfume handles the creation of a new perfume product with image upload
func CreatePerfume(c *fiber.Ctx) error {
	// Parse form fields
	perfume := model.Perfume{
		Name:        c.FormValue("name"),
		Brand:       c.FormValue("brand"),
		Types:       c.FormValue("types"),
		Categories:  c.FormValue("categories"),
		Sizes:       c.FormValue("sizes"),
		Price:       c.FormValue("price"),
		Description: c.FormValue("description"),
		Stock:       c.FormValue("stock"),
	}

	// Get uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image file is required",
			"error":   err.Error(),
		})
	}

	// Read file as base64
	base64String, err := readFileAsBase64(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to process image file",
			"error":   err.Error(),
		})
	}

	// Assign ID and timestamps
	perfume.PerfumeID = primitive.NewObjectID()
	perfume.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	perfume.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Generate a unique filename using perfume ID
	fileName := perfume.PerfumeID.Hex() + filepath.Ext(file.Filename)

	// Upload image and save perfume
	err = repository.CreatePerfume(&perfume, base64String, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create perfume",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Perfume created successfully",
		"perfume": perfume,
	})
}

// readFileAsBase64 converts a file to a Base64 encoded string
func readFileAsBase64(file *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Read file content
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	// Encode file content to Base64
	return base64.StdEncoding.EncodeToString(fileBytes), nil
}

// GetAllPerfumes returns all perfumes from the database
func GetAllPerfumes(c *fiber.Ctx) error {
	perfumes, err := repository.GetAllPerfumes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch perfumes",
			"error":   err.Error(),
		})
	}

	return c.JSON(perfumes)
}