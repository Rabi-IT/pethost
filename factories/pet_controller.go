package factories

import (
	"pethost/adapters/database"
	g "pethost/adapters/gateways/pet_gateway"
	c "pethost/frameworks/http/fiber/controllers/pet_controller"

	"pethost/usecases/pet_case"
)

func NewPet(d database.Database) *c.PetController {
	DB, ok := d.(*database.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormPetGatewayAdapter{DB: DB}
	usecase := pet_case.New(gateway)
	return c.New(usecase)
}
