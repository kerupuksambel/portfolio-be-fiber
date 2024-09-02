package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kerupuksambel/portfolio-be/app/controllers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")

	// Work experience route
	api.Get("/works", controllers.WorkExperience)
}
