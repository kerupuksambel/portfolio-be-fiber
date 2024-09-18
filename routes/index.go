package routes

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"

	"github.com/kerupuksambel/portfolio-be/app/controllers"
)

func getExecPath() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	return filepath.Dir(execPath)
}

func Routes(app *fiber.App) {
	api := app.Group("/api")

	// Routes here
	api.Get("/works", controllers.WorkExperience)
	api.Get("/projects", controllers.Projects)

	// Static serve
	app.Static("/static", filepath.Join(getExecPath(), "static"))
}
