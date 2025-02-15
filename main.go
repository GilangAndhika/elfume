package main

import (
	"log"
	"os"

	"github.com/GilangAndhika/elfume/config"
	"github.com/GilangAndhika/elfume/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// // Load .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Initialize db connection
	config.ConnectDB()

	// Create a new Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} ${message} - ${method} ${path}\n",
	}))

	// Routes
	routes.URL(app)

	// Gunakan port dari environment variable Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s...\n", port)
	log.Fatal(app.Listen(":" + port))
}
