package factories

import (
	"pethost/frameworks/database"
	g "pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/frameworks/database/gorm_adapter"
	c "pethost/frameworks/http/controllers/schedule_controller"

	"pethost/usecases/schedule_case"
)

func NewSchedule(d database.Database) *c.ScheduleController {
	DB, ok := d.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormScheduleGatewayAdapter{DB: DB}
	usecase := schedule_case.New(gateway)
	return c.New(usecase)
}
