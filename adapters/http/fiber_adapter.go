package http

import (
	"pethost/adapters/database"
	"pethost/factories"
	"pethost/frameworks/http/fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type fiberAdapter struct {
	app *fiber.App
}

func (f *fiberAdapter) Start(port string) error {
	return f.app.Listen(":" + port)
}

func (f *fiberAdapter) Stop() error {
	return f.app.Shutdown()
}

func newFiber(d database.Database) HTTPServer {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Use(
		cors.New(),
	).Use(
		requestid.New(),
	)

	routes.Pet(app, factories.NewPet(d))

	return &fiberAdapter{app}
}
