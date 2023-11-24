package factories

import (
	"pethost/frameworks/database"
	g "pethost/frameworks/database/gateways/service_rating_gateway"
	"pethost/frameworks/database/gorm_adapter"
	c "pethost/frameworks/http/controllers/service_rating_controller"

	"pethost/usecases/service_rating_case"
)

func NewServiceRating(d database.Database) *c.ServiceRatingController {
	DB, ok := d.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormServiceRatingGatewayAdapter{DB: DB}
	usecase := service_rating_case.New(gateway)
	return c.New(usecase)
}
