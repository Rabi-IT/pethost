package pet_controller

import (
	"pethost/app_context"
	"pethost/usecases/pet_case"

	"github.com/gofiber/fiber/v2"
)

func (c PetController) Create(ctx *fiber.Ctx) error {
	data := &pet_case.CreateInput{}
	if err := ctx.BodyParser(data); err != nil {
		return ctx.JSON(err)
	}

	id, err := c.usecase.Create(app_context.New(ctx.Context()), data)

	if err != nil {
		return err
	}

	return ctx.Status(201).SendString(id)
}
