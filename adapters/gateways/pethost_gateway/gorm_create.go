package pethost_gateway

import (
	"pethost/frameworks/database/gorm/models"

	"github.com/google/uuid"
)

func (g *GormPetHostGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.PetHost{
		ID:             id,
		Phone:          input.Phone,
		ZIP:            input.ZIP,
		SocialID:       input.SocialID,
		Email:          input.Email,
		EmergencyPhone: input.EmergencyPhone,
		Neighborhood:   input.Neighborhood,
		Street:         input.Street,
		Name:           input.Name,
		TaxID:          input.TaxID,
		City:           input.City,
		State:          input.State,
		Complement:     input.Complement,
	})

	return id, result.Error
}
