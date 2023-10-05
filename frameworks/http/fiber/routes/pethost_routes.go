package routes

import (
	"pethost/frameworks/http/fiber/controllers/host_controller"

	"github.com/gofiber/fiber/v2"
)

func PetHost(app *fiber.App, c *host_controller.PetHostController) {
	route := app.Group("/pethost")
	route.Get("/:id", c.GetByID)
	route.Post("/", c.Create)
	route.Delete("/", c.Delete)
	route.Get("/", c.Paginate)
	route.Patch("/", c.Patch)
}
