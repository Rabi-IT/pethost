package pet_gateway

import "pethost/adapters/database"

type GormPetGatewayAdapter struct {
	DB *database.GormAdapter
}
