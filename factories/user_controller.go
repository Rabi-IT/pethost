package factories

import (
	"pethost/frameworks/database"
	"pethost/frameworks/database/gateways/user_gateway"
	"pethost/frameworks/database/gorm_adapter"
	"pethost/frameworks/http/controllers/user_controller"

	"pethost/usecases/user_case"
)

func NewUser(d database.Database) *user_controller.UserController {
	DB, ok := d.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &user_gateway.GormUserGatewayAdapter{DB: DB}
	usecase := user_case.New(gateway)
	return user_controller.New(usecase)
}
