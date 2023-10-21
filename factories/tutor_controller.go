package factories

import (
	"pethost/adapters/database"
	g "pethost/adapters/gateways/tutor_gateway"
	c "pethost/frameworks/http/fiber/controllers/tutor_controller"

	"pethost/usecases/tutor_case"
)

func NewTutor(d database.Database) *c.TutorController {
	DB, ok := d.(*database.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormTutorGatewayAdapter{DB: DB}
	usecase := tutor_case.New(gateway)
	return c.New(usecase)
}
