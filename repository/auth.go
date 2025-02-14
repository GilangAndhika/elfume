package repository

import (
	"context"
	"elfume/config"
	"elfume/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateAccount handles user registration
func CreateAccount(ctx *fiber.Ctx) error {
	// Get database connection
	client := config.GetDB()
	collection := client.Database("elfume").Collection("users")

	// Parse request body into user struct
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// Validate email format
	if !IsEmailValid(user.Email) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email format",
		})
	}

	// Validate phone number format & convert to international format
	isPhoneValid, formattedPhone := IsPhoneValid(user.Phone)
	if !isPhoneValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid phone number format",
		})
	}
	user.Phone = formattedPhone

	// Check if email already exists
	emailExists, err := IsEmailExists(user.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database error",
		})
	}
	if emailExists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	// Check if username already exists
	usernameExists, err := IsUsernameExists(user.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database error",
		})
	}
	if usernameExists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	// Hash password securely
	hashedPassword := HashPassword(user.Password)
	user.Password = hashedPassword

	// Insert user into the database
	_, err = collection.InsertOne(ctx.Context(), bson.M{
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"phone":      user.Phone,
		"role_id":    model.RoleCustomer,
		"created_at": primitive.NewDateTimeFromTime(time.Now()),
		"updated_at": primitive.NewDateTimeFromTime(time.Now()),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create account",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Account created successfully",
	})
}

// GetUserbyEmail finds a user by email
func GetUserbyEmail(email string) (model.User, error) {
	client := config.GetDB()
	collection := client.Database("elfume").Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}

// GetUserbyUsername finds a user by username
func GetUserbyUsername(username string) (model.User, error) {
	client := config.GetDB()
	collection := client.Database("elfume").Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	return user, err
}

// GetUserbyID finds a user by ID
func GetUserbyID(id string) (model.User, error) {
	client := config.GetDB()
	collection := client.Database("elfume").Collection("users")

	// Convert string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, err // Return an empty struct and an error
	}

	// Find user by ID
	var user model.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	return user, err
}
