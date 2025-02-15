package controller

import (
	"time"

	"github.com/GilangAndhika/elfume/model"
	"github.com/GilangAndhika/elfume/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Registration handles user signup
func Registration(ctx *fiber.Ctx) error {
	var user model.User

	// Parse request body
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate email format
	if !repository.IsEmailValid(user.Email) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email format",
		})
	}

	// Validate phone number format & convert to international format
	isPhoneValid, formattedPhone := repository.IsPhoneValid(user.Phone)
	if !isPhoneValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid phone number format",
		})
	}
	user.Phone = formattedPhone

	// Check if email already exists
	emailExists, err := repository.IsEmailExists(user.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database error",
			"error":   err.Error(),
		})
	}
	if emailExists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	// Check if username already exists
	usernameExists, err := repository.IsUsernameExists(user.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database error",
			"error":   err.Error(),
		})
	}
	if usernameExists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	// Hash password securely
	user.Password = repository.HashPassword(user.Password)

	// Assign default role if RoleID is not provided
	if user.RoleID.IsZero() {
		roleID, err := primitive.ObjectIDFromHex(model.RoleCustomer)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid default role ID",
				"error":   err.Error(),
			})
		}
		user.RoleID = roleID
	}

	// Set created_at and updated_at timestamps
	user.UserID = primitive.NewObjectID()
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Create the user account
	err = repository.CreateAccount(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create account",
			"error":   err.Error(),
		})
	}


	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Account created successfully",
	})
}
