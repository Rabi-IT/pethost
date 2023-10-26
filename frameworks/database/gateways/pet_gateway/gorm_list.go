package pet_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormPetGatewayAdapter) List(filter ListInput) ([]ListOutput, error) {
	query := g.DB.Conn.Model(&models.Pet{})

	if filter.Name != nil {
		query = query.Where("name = ?", filter.Name)
	}

	if filter.Breed != nil {
		query = query.Where("breed = ?", filter.Breed)
	}

	if filter.Size != nil {
		query = query.Where("size = ?", filter.Size)
	}

	if filter.Birthdate != nil {
		query = query.Where("birthdate = ?", filter.Birthdate)
	}

	if filter.Gender != nil {
		query = query.Where("gender = ?", filter.Gender)
	}

	if filter.Weight != nil {
		query = query.Where("weight = ?", filter.Weight)
	}

	if filter.Species != nil {
		query = query.Where("species = ?", filter.Species)
	}

	output := []ListOutput{}
	result := query.Scan(&output)

	return output, result.Error
}
