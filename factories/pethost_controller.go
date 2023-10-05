package factories

import (
	"pethost/adapters/database"
	g "pethost/adapters/gateways/pethost_gateway"
	c "pethost/frameworks/http/fiber/controllers/host_controller"
	"pethost/usecases/pethost_case"
)

func NewPetHost(d database.Database) *c.PetHostController {
	DB, ok := d.(*database.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormPetHostGatewayAdapter{DB: DB}
	usecase := pethost_case.New(gateway)
	return c.New(usecase)
}
