package fiber_adapter

import (
	"pethost/config"
	"pethost/factories"
	"pethost/frameworks/database"
	"pethost/frameworks/http"
	"pethost/frameworks/http/controllers/auth_controller"
	"pethost/frameworks/http/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	jwtware "github.com/gofiber/contrib/jwt"
)

type fiberAdapter struct {
	app *fiber.App
}

func New(d database.Database) http.HTTPServer {
	return newFiber(d)
}

func (f *fiberAdapter) Start(port string) error {
	return f.app.Listen(":" + port)
}

func (f *fiberAdapter) Stop() error {
	return f.app.Shutdown()
}

func newFiber(d database.Database) http.HTTPServer {
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

	routes.Schedule(app, factories.NewSchedule(d))
	routes.Preference(app, factories.NewPreference(d))
	routes.Pet(app, factories.NewPet(d))
	routes.User(app, userController)

	return &fiberAdapter{app}
}
