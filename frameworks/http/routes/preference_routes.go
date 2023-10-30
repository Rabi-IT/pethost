package routes

import (
	"pethost/frameworks/http/controllers/preference_controller"

	"github.com/gofiber/fiber/v2"
)

func Preference(app *fiber.App, c *preference_controller.PreferenceController) {
	route := app.Group("/preference")
	route.Post("/", c.Create)
}
