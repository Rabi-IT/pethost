package http

import (
	"pethost/adapters/database"
	"pethost/config"
	"pethost/factories"
	"pethost/frameworks/http/fiber/controllers/auth_controller"
	"pethost/frameworks/http/fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	jwtware "github.com/gofiber/contrib/jwt"
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

	userController := factories.NewUser(d)
	app.Use(
		cors.New(),
	).Use(
		requestid.New(),
	).Post(
		"/user", userController.Create,
	).Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(config.AuthSecret)},
		}),
	).Use(
		auth_controller.Session,
	)

	routes.Pet(app, factories.NewPet(d))
	routes.User(app, userController)

	return &fiberAdapter{app}
}
