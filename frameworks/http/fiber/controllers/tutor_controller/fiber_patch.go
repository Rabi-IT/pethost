package tutor_controller

import (
	"pethost/frameworks/http/fiber/parser"
	"pethost/usecases/tutor_case"

	"github.com/gofiber/fiber/v2"
)

func (c TutorController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	filter := &tutor_case.PatchFilter{
		ID: &id,
	}

	data := tutor_case.PatchValues{}
	if err := parser.ParseBody(ctx, &data); err != nil {
		return ctx.JSON(err)
	}

	updated, err := c.usecase.Patch(ctx.Context(), *filter, data)

	if err != nil {
		return err
	}

	if updated {
		return ctx.SendStatus(200)
	}

	return ctx.SendStatus(404)
}
