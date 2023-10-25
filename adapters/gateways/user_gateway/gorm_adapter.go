package user_gateway

import "pethost/adapters/database"

type GormUserGatewayAdapter struct {
	DB *database.GormAdapter
}
