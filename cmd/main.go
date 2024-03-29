package main

import (
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Listen(":4040")
}