package host_controller

import (
	database "pethost/frameworks/database/gorm"
	"pethost/usecases/host_case"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (c PetHostController) Paginate(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("Page"))
	if err != nil {
		return err
	}

	pageSize, err := strconv.Atoi(ctx.Query("PageSize"))
	if err != nil {
		return err
	}

	ZIP := ctx.Query("ZIP")
	Email := ctx.Query("Email")
	Neighborhood := ctx.Query("Neighborhood")
	Complement := ctx.Query("Complement")
	Name := ctx.Query("Name")
	City := ctx.Query("City")
	State := ctx.Query("State")

	filter := host_case.PaginateFilter{
		ZIP:          &ZIP,
		Email:        &Email,
		Neighborhood: &Neighborhood,
		Complement:   &Complement,
		Name:         &Name,
		City:         &City,
		State:        &State,
	}

	paginate := database.PaginateInput{
		Page:     page,
		PageSize: pageSize,
	}

	result, err := c.usecase.Paginate(ctx.Context(), filter, paginate)

	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
