package pethost_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g GormPetHostGatewayAdapter) GetByID(id string) (*GetByIDOutput, error) {
	output := &models.PetHost{}
	result := g.DB.Conn.Limit(1).Find(output, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	adapted := GetByIDOutput{
		Neighborhood:   output.Neighborhood,
		Street:         output.Street,
		Email:          output.Email,
		EmergencyPhone: output.EmergencyPhone,
		State:          output.State,
		Name:           output.Name,
		City:           output.City,
		ZIP:            output.ZIP,
	}

	return &adapted, nil
}
