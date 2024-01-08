package routes

import (
	"test_docker/controller"

	"github.com/gofiber/fiber/v2"
)

func APIs(app *fiber.App) {
	v1 := app.Group("v1.0")
	v1.Get("/say_hello", controller.SayHello)
	v1.Post("/calculating", controller.CalculatingNumber)
}
