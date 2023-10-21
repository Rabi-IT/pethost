package routes

import (
	"pethost/frameworks/http/fiber/controllers/tutor_controller"

	"github.com/gofiber/fiber/v2"
)

func Tutor(app *fiber.App, c *tutor_controller.TutorController) {
	route := app.Group("/tutor")
	route.Delete("/:id", c.Delete)
	route.Get("/", c.Paginate)
	route.Get("/", c.List)
	route.Patch("/:id", c.Patch)
	route.Get("/:id", c.GetByID)
	route.Post("/", c.Create)
}
