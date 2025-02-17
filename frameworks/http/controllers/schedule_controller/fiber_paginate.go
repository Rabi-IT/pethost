package schedule_controller

import (
	"pethost/app_context"
	"pethost/frameworks/database"
	"pethost/usecases/schedule_case"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (c ScheduleController) Paginate(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("Page", "0"))
	if err != nil {
		return err
	}

	pageSize, err := strconv.Atoi(ctx.Query("PageSize", "10"))
	if err != nil {
		return err
	}

	filter := schedule_case.PaginateFilter{}
	if err = ctx.QueryParser(&filter); err != nil {
		return err
	}

	paginate := database.PaginateInput{
		Page:     page,
		PageSize: pageSize,
	}

	result, err := c.usecase.Paginate(app_context.New(ctx.Context()), filter, paginate)

	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
