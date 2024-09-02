package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/kerupuksambel/portfolio-be/routes"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Static("/public", "./public")

	app.Listen(":6060")
	fmt.Printf("Listening on port 6060")
}
