package routes

import (
	"github.com/GilangAndhika/elfume/controller"
	"github.com/GilangAndhika/elfume/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func URL(app *fiber.App) {
	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Elfume connected!")
	})

	// Auth routes (Registration & Login)
	AuthRoutes := app.Group("/auth")
	AuthRoutes.Post("/register", controller.Registration)
	AuthRoutes.Post("/login", controller.Login)
	AuthRoutes.Post("/logout", controller.Logout)

	// User Routes
	UserRoutes := app.Group("/user")
	UserRoutes.Get("/all", controller.GetAllUsers)
	UserRoutes.Get("/id/:id", controller.GetUserByID)
	UserRoutes.Put("/update/:id", controller.UpdateUser)
	UserRoutes.Delete("/delete/:id", controller.DeleteUser)

	// Role routes
	RoleRoutes := app.Group("/role")
	RoleRoutes.Post("/create", controller.CreateRole)

	// Perfume routes
	PerfumeRoutes := app.Group("/fume")
	PerfumeRoutes.Post("/create", controller.CreatePerfume)
	PerfumeRoutes.Post("/insert", controller.CreatePerfumeWithoutImage)
	PerfumeRoutes.Get("/all", controller.GetAllPerfumes)
	PerfumeRoutes.Get("/id/:id", controller.GetPerfumeByID)
	PerfumeRoutes.Get("/search", controller.GetFilteredPerfumes)
	PerfumeRoutes.Put("/update/:id", controller.UpdatePerfume)
	PerfumeRoutes.Delete("/delete/:id", controller.DeletePerfume)

	// Protected route (requires authentication)
	app.Get("/protected", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}
		return c.JSON(fiber.Map{"message": "Access granted", "user": user})
	})
}
