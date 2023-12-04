package service_rating_controller

import (
	"pethost/app_context"
	"pethost/frameworks/http/fiber_adapter/parser"
	"pethost/usecases/service_rating_case"

	"github.com/gofiber/fiber/v2"
)

func (c ServiceRatingController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	filter := &service_rating_case.PatchFilter{
		ID: &id,
	}

	data := service_rating_case.PatchValues{}
	if err := parser.ParseBody(ctx, &data); err != nil {
		return ctx.JSON(err)
	}

	updated, err := c.usecase.Patch(app_context.New(ctx.Context()), *filter, data)

	if err != nil {
		return err
	}

	if updated {
		return ctx.SendStatus(200)
	}

	return ctx.SendStatus(404)
}
