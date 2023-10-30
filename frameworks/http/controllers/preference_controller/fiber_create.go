package preference_controller

import (
	"pethost/app_context"
	"pethost/usecases/preference_case"

	"github.com/gofiber/fiber/v2"
)

func (c PreferenceController) Create(ctx *fiber.Ctx) error {
	data := &preference_case.CreateInput{}
	if err := ctx.BodyParser(data); err != nil {
		return ctx.JSON(err)
	}

	id, err := c.usecase.Create(app_context.New(ctx.Context()), data)

	if err != nil {
		return err
	}

	return ctx.Status(201).SendString(id)
}
