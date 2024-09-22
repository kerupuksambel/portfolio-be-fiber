package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/kerupuksambel/portfolio-be/internal/database"
	// "github.com/kerupuksambel/portfolio-be/internal/database/migration"
	"github.com/kerupuksambel/portfolio-be/routes"
)

func main() {
	// Args read
	var (
		isMigration bool
	)
	flag.BoolVar(&isMigration, "migrate", false, "Include if you want to run migration (not forcing migration)")
	flag.Parse()

	// Initialize DB
	database.InitDB()

	// Initiate migration
	if isMigration {
		database.Migrate()
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
	}))

	routes.Routes(app)

	app.Listen(":6060")
	fmt.Printf("Listening on port 6060")
}
