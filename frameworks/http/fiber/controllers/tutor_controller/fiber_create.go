package tutor_controller

import (
	"pethost/usecases/tutor_case"

	"github.com/gofiber/fiber/v2"
)

func (c TutorController) Create(ctx *fiber.Ctx) error {
	data := &tutor_case.CreateInput{}
	if err := ctx.BodyParser(data); err != nil {
		return ctx.JSON(err)
	}

	id, err := c.usecase.Create(ctx.Context(), data)

	if err != nil {
		return err
	}

	return ctx.Status(201).SendString(id)
}
