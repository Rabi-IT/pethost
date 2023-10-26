package factories

import (
	"pethost/frameworks/database"
	"pethost/frameworks/database/gateways/pet_gateway"
	"pethost/frameworks/database/gorm_adapter"
	"pethost/frameworks/http/controllers/pet_controller"

	"pethost/usecases/pet_case"
)

func NewPet(d database.Database) *pet_controller.PetController {
	DB, ok := d.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &pet_gateway.GormPetGatewayAdapter{DB: DB}
	usecase := pet_case.New(gateway)
	return pet_controller.New(usecase)
}
