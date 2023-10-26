package pet_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormPetGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	query := g.DB.Conn.Model(&models.Pet{})

	if filter.ID != nil {
		query = query.Where("id = ?", filter.ID)
	} else {
		if filter.Species != nil {
			query = query.Where("Species = ?", filter.Species)
		}

		if filter.Name != nil {
			query = query.Where("Name = ?", filter.Name)
		}

		if filter.Breed != nil {
			query = query.Where("Breed = ?", filter.Breed)
		}

		if filter.Size != nil {
			query = query.Where("Size = ?", filter.Size)
		}

		if filter.Birthdate != nil {
			query = query.Where("Birthdate = ?", filter.Birthdate)
		}

		if filter.Gender != nil {
			query = query.Where("Gender = ?", filter.Gender)
		}

		if filter.Weight != nil {
			query = query.Where("Weight = ?", filter.Weight)
		}
	}

	result := query.Updates(newValues)
	return result.RowsAffected > 0, result.Error
}
