package pet_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormPetGatewayAdapter) List(filter ListInput) ([]ListOutput, error) {
	query := g.DB.Conn.Model(&models.Pet{})

	if filter.TutorID != nil {
		query = query.Where("tutor_id = ?", filter.TutorID)
	}

	output := []ListOutput{}
	result := query.Scan(&output)

	return output, result.Error
}
