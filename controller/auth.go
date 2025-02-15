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
		"user":    user,
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

// GetAllUsers handles retrieving all users from the database
func GetAllUsers(c *fiber.Ctx) error {
	users, err := repository.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users retrieved successfully",
		"users":   users,
	})
}

// UpdateUser handles updating an existing user's information
func UpdateUser(c *fiber.Ctx) error {
	// Get user ID from params
	userID := c.Params("id")

	// Parse request body
	var updatedUser model.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Update the user in the database
	err := repository.UpdateUser(userID, updatedUser)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

// DeleteUser handles deleting an existing user
func DeleteUser(c *fiber.Ctx) error {
	// Get user ID from URL params
	userID := c.Params("id")

	// Delete the user from the database
	err := repository.DeleteUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

// GetUserByID handles retrieving a user by their ID
func GetUserByID(c *fiber.Ctx) error {
	// Get user ID from URL params
	userID := c.Params("id")

	// Retrieve the user from the database
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
		"user":    user,
	})
}

// Logout handles user logout by clearing the JWT cookie
func Logout(c *fiber.Ctx) error {
	// Clear the JWT token by setting an expired cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}