package routes

import "github.com/gofiber/fiber/v2"

func URL(app *fiber.App) {
	// Define the routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Elfume connected!")
	})

}
