package routes

import (
	"pethost/frameworks/http/fiber/controllers/pet_controller"

	"github.com/gofiber/fiber/v2"
)

func Pet(app *fiber.App, c *pet_controller.PetController) {
	route := app.Group("/pet")
	route.Post("/", c.Create)
	route.Patch("/:id", c.Patch)
	route.Get("/:id", c.GetByID)
	route.Get("/", c.List)
	route.Delete("/:id", c.Delete)
}
