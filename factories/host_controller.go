package factories

import (
	"pethost/adapters/database"
	g "pethost/adapters/gateways/host_gateway"
	c "pethost/frameworks/http/fiber/controllers/host_controller"
	"pethost/usecases/host_case"
)

func NewPetHost(d database.Database) *c.PetHostController {
	DB, ok := d.(*database.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormPetHostGatewayAdapter{DB: DB}
	usecase := host_case.New(gateway)
	return c.New(usecase)
}
