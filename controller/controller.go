package controller

import (
	"fmt"
	"test_docker/models"

	"github.com/gofiber/fiber/v2"
)

func SayHello(ctx *fiber.Ctx) error {
	input := ctx.Query("name")
	message := fmt.Sprintf("xin chào %s", input)
	return ctx.JSON(models.Reponse{
		Message: "Success",
		Data:    message,
		Error:   nil,
	})
}

func CalculatingNumber(ctx *fiber.Ctx) error {
	body := new(models.NumberModel)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.JSON(models.Reponse{
			Message: "Failed",
			Data:    nil,
			Error:   err,
		})
	}

	if body.NumberOne == 0 {
		body.NumberOne = 1
	}

	if body.NumberTwo == 0 {
		body.NumberTwo = 1
	}

	return ctx.JSON(models.Reponse{
		Message: "Success",
		Data:    body.NumberOne * body.NumberTwo,
		Error:   nil,
	})

}
