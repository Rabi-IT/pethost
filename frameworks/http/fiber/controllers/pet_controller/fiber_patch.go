package pet_controller

import (
	"pethost/frameworks/http/fiber/parser"
	"pethost/usecases/pet_case"

	"github.com/gofiber/fiber/v2"
)

func (c PetController) Patch(ctx *fiber.Ctx) error {
	filter := &pet_case.PatchFilter{}
	if err := ctx.QueryParser(&filter); err != nil {
		return err
	}

	data := pet_case.PatchValues{}
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
