package controller

import (
	"github.com/GilangAndhika/elfume/config"
	"github.com/GilangAndhika/elfume/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

// CreateRole handles role creation
func CreateRole(c *fiber.Ctx) error {
	var role model.Role

	// Get database connection
	db := config.MongoClient.Database("elfume")

	// Parse request body
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Get the roles collection
	collection := db.Collection("roles")

	// Create role with the repository
	role.RoleID = primitive.NewObjectID()
	_, err := collection.InsertOne(c.Context(), role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create role",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role created successfully",
		"role":    role,
	})

}
