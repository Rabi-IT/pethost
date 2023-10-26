package pet_gateway

import "pethost/frameworks/database/gorm_adapter"

type GormPetGatewayAdapter struct {
	DB *gorm_adapter.GormAdapter
}
