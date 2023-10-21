package tutor_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g GormTutorGatewayAdapter) GetByID(id string) (*GetByIDOutput, error) {
	output := &models.Tutor{}
	result := g.DB.Conn.Limit(1).Find(output, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	adapted := GetByIDOutput{
		State:          output.State,
		ZIP:            output.ZIP,
		Phone:          output.Phone,
		City:           output.City,
		Photo:          output.Photo,
		TaxID:          output.TaxID,
		SocialID:       output.SocialID,
		Street:         output.Street,
		Complement:     output.Complement,
		EmergencyPhone: output.EmergencyPhone,
		Name:           output.Name,
		Email:          output.Email,
	}

	return &adapted, nil
}
