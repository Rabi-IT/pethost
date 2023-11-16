package pet_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormPetGatewayAdapter) List(filter ListInput) ([]ListOutput, error) {
	query := g.DB.Conn.Model(&models.Pet{})

	if filter.TutorID != nil {
		query = query.Where("tutor_id = ?", filter.TutorID)
	}

	if filter.PetIDs != nil {
		query = query.Where("id IN ?", filter.PetIDs)
	}

	output := []ListOutput{}
	result := query.Scan(&output)

	return output, result.Error
}
