package main

import (
	"fmt"
	// "github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/kerupuksambel/portfolio-be/internal/database"
	"github.com/kerupuksambel/portfolio-be/routes"
)

func main() {
	// Initialize DB
	database.InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
	}))

	routes.Routes(app)

	app.Listen(":6060")
	fmt.Printf("Listening on port 6060")
}
