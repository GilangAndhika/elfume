package repository

import (
	"elfume/config"
	"elfume/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAccount(ctx *fiber.Ctx) error {
	// Get connection
	client := config.GetDB()
	collection := client.Database("elfume").Collection("users")

	// Parse request body into user struct
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// Hash password
	user.Password = HashPassword(user.Password)

	// Insert user data
	_, err := collection.InsertOne(ctx.Context(), bson.M{
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"phone":      user.Phone,
		"role_id":    model.RoleCustomer,
		"created_at": primitive.NewDateTimeFromTime(time.Now()),
		"updated_at": primitive.NewDateTimeFromTime(time.Now()),
	})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to create account",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Account created successfully",
	})
}
