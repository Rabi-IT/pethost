package pet_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormPetGatewayAdapter) GetByID(id string) (*GetByIDOutput, error) {
	output := &models.Pet{}
	result := g.DB.Conn.Limit(1).Find(output, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	adapted := GetByIDOutput{
		Size:      output.Size,
		Birthdate: output.Birthdate,
		Gender:    output.Gender,
		Weight:    output.Weight,
		Species:   output.Species,
		Name:      output.Name,
		Breed:     output.Breed,
	}

	return &adapted, nil
}
