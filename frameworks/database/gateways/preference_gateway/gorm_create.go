package preference_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"

	"github.com/google/uuid"
)

func (g *GormPreferenceGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.Preference{
		ID:                      id,
		PetWeight:               input.PetWeight,
		AcceptFemaleInHeat:      input.AcceptFemaleInHeat,
		AcceptPuppies:           input.AcceptPuppies,
		AcceptMales:             input.AcceptMales,
		AcceptFemales:           input.AcceptFemales,
		DaysOfMonth:             input.DaysOfMonth,
		OnlyVaccinated:          input.OnlyVaccinated,
		AcceptElderly:           input.AcceptElderly,
		AcceptOnlyNeuteredMales: input.AcceptOnlyNeuteredMales,
	})

	return id, result.Error
}
