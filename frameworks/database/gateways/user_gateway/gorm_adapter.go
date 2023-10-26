package user_gateway

import "pethost/frameworks/database/gorm_adapter"

type GormUserGatewayAdapter struct {
	DB *gorm_adapter.GormAdapter
}
