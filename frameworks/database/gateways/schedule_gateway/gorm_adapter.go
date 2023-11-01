package schedule_gateway

import "pethost/frameworks/database/gorm_adapter"

type GormScheduleGatewayAdapter struct {
	DB *gorm_adapter.GormAdapter
}
