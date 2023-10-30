package preference_gateway

import "pethost/frameworks/database/gorm_adapter"

type GormPreferenceGatewayAdapter struct {
	DB *gorm_adapter.GormAdapter
}
