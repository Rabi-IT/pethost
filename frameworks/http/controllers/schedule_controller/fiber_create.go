package schedule_controller

import (
	"pethost/usecases/schedule_case"

	"github.com/gofiber/fiber/v2"
)

func (c ScheduleController) Create(ctx *fiber.Ctx) error {
	data := &schedule_case.CreateInput{}
	if err := ctx.BodyParser(data); err != nil {
		return ctx.JSON(err)
	}

	id, err := c.usecase.Create(ctx.Context(), data)

	if err != nil {
		return err
	}

	return ctx.Status(201).SendString(id)
}
