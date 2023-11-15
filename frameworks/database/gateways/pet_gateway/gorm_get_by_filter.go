package pet_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormPetGatewayAdapter) GetByFilter(filter GetByFilterInput) (*GetByFilterOutput, error) {
	query := g.DB.Conn.Model(&models.Pet{}).Where(
		"id = ?", filter.ID,
	).Limit(1)

	if filter.TutorID != nil {
		query = query.Where("tutor_id = ?", filter.TutorID)
	}

	output := &GetByFilterOutput{}
	result := query.Scan(output)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return output, nil
}
