package user_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"

	"github.com/google/uuid"
)

func (g *GormUserGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.User{
		ID:             id,
		TaxID:          input.TaxID,
		SocialID:       input.SocialID,
		Street:         input.Street,
		Complement:     input.Complement,
		EmergencyPhone: input.EmergencyPhone,
		Name:           input.Name,
		Email:          input.Email,
		Photo:          input.Photo,
		ZIP:            input.ZIP,
		Phone:          input.Phone,
		City:           input.City,
		State:          input.State,
		Neighborhood:   input.Neighborhood,
		IsHost:         input.IsHost,
	})

	return id, result.Error
}
