package schedule_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormScheduleGatewayAdapter) GetByID(id string) (*GetByIDOutput, error) {
	output := &models.Schedule{}
	result := g.DB.Conn.Limit(1).Find(output, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	adapted := GetByIDOutput{
		Status: output.Status,
	}

	return &adapted, nil
}
