package controller

import (
	"time"

	"github.com/GilangAndhika/elfume/middleware"
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
		"user":   user,
	})
}

// Login handles user login
func Login(ctx *fiber.Ctx) error {
	var user model.User

	// Parse request body
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Find user by email or username
	foundUser, err := repository.GetUserByEmailOrUsername(user.Email, user.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find user",
			"error":   err.Error(),
		})
	}
	if foundUser == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Compare passwords
	match, err := repository.ComparePassword(foundUser.Password, user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to compare passwords",
			"error":   err.Error(),
		})
	}
	if !match {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Generate JWT Token
	token, err := middleware.GenerateJWT(foundUser)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
	}

	// Set JWT token in HTTP-Only Cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(2 * time.Hour), // Cookie expires in 2 hours
		HTTPOnly: true,                          // Secure HTTP-Only
		Secure:   true,                          // Set to true if using HTTPS
	})

	// Return token in response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
