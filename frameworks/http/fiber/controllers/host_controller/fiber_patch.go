package host_controller

import (
	"pethost/frameworks/http/fiber/parser"
	"pethost/usecases/pethost_case"

	"github.com/gofiber/fiber/v2"
)

func (c PetHostController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	filter := &pethost_case.PatchFilter{
		ID: id,
	}

	data := pethost_case.PatchValues{}
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
