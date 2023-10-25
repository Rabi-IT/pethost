package factories

import (
	"pethost/adapters/database"
	g "pethost/adapters/gateways/user_gateway"
	c "pethost/frameworks/http/fiber/controllers/user_controller"

	"pethost/usecases/user_case"
)

func NewUser(d database.Database) *c.UserController {
	DB, ok := d.(*database.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormUserGatewayAdapter{DB: DB}
	usecase := user_case.New(gateway)
	return c.New(usecase)
}
