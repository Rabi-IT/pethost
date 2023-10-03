package host_controller

import (
	"pethost/frameworks/http/fiber/parser"
	"pethost/usecases/host_case"

	"github.com/gofiber/fiber/v2"
)

func (c PetHostController) Create(ctx *fiber.Ctx) error {
	data := &host_case.CreateInput{}
	if err := parser.ParseBody(ctx, data); err != nil {
		return ctx.JSON(err)
	}

	id, err := c.usecase.Create(ctx.Context(), data)

	if err != nil {
		return err
	}

	return ctx.SendString(id)
}
