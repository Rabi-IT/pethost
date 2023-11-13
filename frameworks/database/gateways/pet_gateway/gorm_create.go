package pet_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"

	"github.com/google/uuid"
)

func (g *GormPetGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.Pet{
		ID:        id,
		Weight:    input.Weight,
		Species:   input.Species,
		Name:      input.Name,
		Breed:     input.Breed,
		Birthdate: input.Birthdate,
		Gender:    input.Gender,
		UserID:    input.UserID,
		Neutered:  input.Neutered,
	})

	return id, result.Error
}
