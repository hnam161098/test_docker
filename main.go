package main

import (
	"test_docker/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.APIs(app)
	app.Listen(":8080")
}
