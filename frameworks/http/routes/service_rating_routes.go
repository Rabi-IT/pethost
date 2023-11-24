package routes

import (
	"pethost/frameworks/http/controllers/service_rating_controller"

	"github.com/gofiber/fiber/v2"
)

func ServiceRating(app *fiber.App, c *service_rating_controller.ServiceRatingController) {
	route := app.Group("/service_rating")
	route.Delete("/:id", c.Delete)
	route.Patch("/:id", c.Patch)
	route.Get("/", c.Paginate)
	route.Post("/", c.Create)
}
