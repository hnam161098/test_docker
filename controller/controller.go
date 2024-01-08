package controller

import (
	"fmt"
	"test_docker/models"

	"github.com/gofiber/fiber/v2"
)

func SayHello(ctx *fiber.Ctx) error {
	input := ctx.Params("name")
	message := fmt.Sprintf("xin chào %s", input)
	return ctx.JSON(models.Reponse{
		Message: "Success",
		Data:    message,
		Error:   nil,
	})
}
