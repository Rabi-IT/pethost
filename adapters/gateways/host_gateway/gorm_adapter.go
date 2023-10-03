package host_gateway

import "pethost/adapters/database"

type GormPetHostGatewayAdapter struct {
	DB *database.GormAdapter
}
