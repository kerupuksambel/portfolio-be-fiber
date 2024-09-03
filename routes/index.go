package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kerupuksambel/portfolio-be/app/controllers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")

	// Routes here
	api.Get("/works", controllers.WorkExperience)
	api.Get("/projects", controllers.Projects)

	// Static serve
	app.Static("/public", "./public")
}
