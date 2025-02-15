package routes

import (
	"github.com/GilangAndhika/elfume/controller"

	"github.com/gofiber/fiber/v2"
)

func URL(app *fiber.App) {
	// Define the routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Elfume connected!")
	})

	// Role routes
	RoleRoutes := app.Group("/role")
	RoleRoutes.Post("/create", controller.CreateRole)

	// Auth routes
	AuthRoutes := app.Group("/auth")
	AuthRoutes.Post("/register", controller.Registration)
}
