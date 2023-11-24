package service_rating_gateway

import "pethost/frameworks/database/gorm_adapter"

type GormServiceRatingGatewayAdapter struct {
	DB *gorm_adapter.GormAdapter
}
