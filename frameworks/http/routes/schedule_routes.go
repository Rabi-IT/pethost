package routes

import (
	"pethost/frameworks/http/controllers/schedule_controller"

	"github.com/gofiber/fiber/v2"
)

func Schedule(app *fiber.App, c *schedule_controller.ScheduleController) {
	route := app.Group("/schedule")
	route.Get("/", c.Paginate)
	route.Patch("/:id", c.Patch)
	route.Post("/", c.Create)
}
