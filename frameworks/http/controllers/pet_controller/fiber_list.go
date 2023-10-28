package pet_controller

import (
	"pethost/app_context"

	"github.com/gofiber/fiber/v2"
)

func (c PetController) List(ctx *fiber.Ctx) error {
	result, err := c.usecase.List(app_context.New(ctx.Context()))

	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
