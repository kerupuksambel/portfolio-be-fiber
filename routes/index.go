package routes

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"

	"github.com/kerupuksambel/portfolio-be/app/controllers"
)

func Routes(app *fiber.App) {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	basePath := filepath.Dir(execPath)

	api := app.Group("/api")

	// Routes here
	api.Get("/works", controllers.WorkExperience)
	api.Get("/projects", controllers.Projects)

	// Static serve
	app.Static("/static", filepath.Join(basePath, "static"))
}
