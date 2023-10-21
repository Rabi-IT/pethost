package tutor_gateway

import "pethost/adapters/database"

type GormTutorGatewayAdapter struct {
	DB *database.GormAdapter
}
