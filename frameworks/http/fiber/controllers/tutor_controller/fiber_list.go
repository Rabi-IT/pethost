package tutor_controller

import (
	"pethost/usecases/tutor_case"

	"github.com/gofiber/fiber/v2"
)

func (c TutorController) List(ctx *fiber.Ctx) error {
	filter := &tutor_case.ListInput{}
	if err := ctx.QueryParser(filter); err != nil {
		return err
	}

	result, err := c.usecase.List(ctx.Context(), *filter)

	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
