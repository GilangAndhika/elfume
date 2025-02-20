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

// GetPerfumeByID returns a single perfume by ID
func GetPerfumeByID(c *fiber.Ctx) error {
	perfumeID := c.Params("id")
	perfume, err := repository.GetPerfumeByID(perfumeID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Perfume not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(perfume)
}

// GetFilteredPerfumes returns perfumes with optional filters
func GetFilteredPerfumes(c *fiber.Ctx) error {
	// Get query parameters (e.g., ?name=Dior&size=100ml)
	filters := make(map[string]string)
	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}
	if size := c.Query("size"); size != "" {
		filters["sizes"] = size
	}
	if brand := c.Query("brand"); brand != "" {
		filters["brand"] = brand
	}
	if category := c.Query("categories"); category != "" {
		filters["categories"] = category
	}
	if types := c.Query("types"); types != "" {
		filters["types"] = types
	}
	if price := c.Query("price"); price != "" {
		filters["price"] = price
	}

	// Fetch perfumes with filters
	perfumes, err := repository.GetFilteredPerfumes(filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch perfumes",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Perfumes retrieved successfully",
		"perfumes": perfumes,
	})
}

// UpdatePerfume handles updating an existing perfume
func UpdatePerfume(c *fiber.Ctx) error {
	// Get perfume ID from params
	perfumeID := c.Params("id")

	// Parse request body
	var updatedPerfume model.Perfume
	if err := c.BodyParser(&updatedPerfume); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Update the perfume in the database
	err := repository.UpdatePerfume(perfumeID, updatedPerfume)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to update perfume",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Perfume updated successfully",
	})
}

// DeletePerfume handles deleting an existing perfume
func DeletePerfume(c *fiber.Ctx) error {
	// Get perfume ID from URL params
	perfumeID := c.Params("id")

	// Delete the perfume from the database
	err := repository.DeletePerfume(perfumeID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to delete perfume",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Perfume deleted successfully",
	})
}

// create perfume without image upload instead of image url
func CreatePerfumeWithoutImage(c *fiber.Ctx) error {
	perfume := new(model.Perfume)

	if err := c.BodyParser(perfume); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	perfume.PerfumeID = primitive.NewObjectID()
	perfume.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	perfume.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	err := repository.CreatePerfumeWithImageURL(perfume)
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