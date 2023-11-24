package service_rating_controller

import (
	"pethost/usecases/service_rating_case"

	"github.com/gofiber/fiber/v2"
)

func (c ServiceRatingController) Create(ctx *fiber.Ctx) error {
	data := &service_rating_case.CreateInput{}
	if err := ctx.BodyParser(data); err != nil {
		return ctx.JSON(err)
	}

	id, err := c.usecase.Create(ctx.Context(), data)

	if err != nil {
		return err
	}

	return ctx.Status(201).SendString(id)
}
