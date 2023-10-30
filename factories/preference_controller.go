package factories

import (
	"pethost/frameworks/database"
	g "pethost/frameworks/database/gateways/preference_gateway"
	"pethost/frameworks/database/gorm_adapter"
	c "pethost/frameworks/http/controllers/preference_controller"

	"pethost/usecases/preference_case"
)

func NewPreference(d database.Database) *c.PreferenceController {
	DB, ok := d.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(ErrDatabaseAdapter)
	}

	gateway := &g.GormPreferenceGatewayAdapter{DB: DB}
	usecase := preference_case.New(gateway)
	return c.New(usecase)
}
