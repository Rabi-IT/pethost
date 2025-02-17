package routes

import (
	"pethost/frameworks/http/controllers/user_controller"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App, c *user_controller.UserController) {
	route := app.Group("/user")
	route.Patch("/:id", c.Patch)
	route.Delete("/:id", c.Delete)
}
