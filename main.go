package main

import (
	"fmt"
	"test_docker/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello world")

	app := fiber.New()
	routes.APIs(app)
	app.Listen(":8080")
}
