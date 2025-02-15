package middleware

import (
	"os"
	"time"

	"github.com/GilangAndhika/elfume/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Get JWT secret from .env
var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

// GenerateJWT creates a new JWT token for authentication
func GenerateJWT(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user.UserID.Hex(),
		"username":  user.Username,
		"role_id":   user.RoleID.Hex(),
		"role_name": user.RoleName,
		"exp":       time.Now().Add(time.Hour * 2).Unix(), // Token expires in 2 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWTMiddleware protects routes with authentication
func JWTMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get token from cookie
		tokenStr := c.Cookies("token")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		// Parse token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
		}

		// Pass user data to the next handler
		c.Locals("user", token.Claims.(jwt.MapClaims))
		return c.Next()
	}
}
