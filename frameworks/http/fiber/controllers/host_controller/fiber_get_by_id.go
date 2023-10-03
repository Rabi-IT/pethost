package host_controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c PetHostController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := c.usecase.GetByID(ctx.Context(), id)

	if err != nil {
		return err
	}

	if data == nil {
		return ctx.SendStatus(404)
	}

	return ctx.JSON(data)
}
