package preference_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g *GormPreferenceGatewayAdapter) GetByFilter(filter *GetByFilterInput) (*GetByFilterOutput, error) {
	output := &models.Preference{}
	query := g.DB.Conn.Find(output, "user_id = ?", filter.UserID).Limit(1)

	result := query.Scan(&output)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	adapted := GetByFilterOutput{
		DaysOfMonth:             output.DaysOfMonth,
		OnlyVaccinated:          output.OnlyVaccinated,
		AcceptElderly:           output.AcceptElderly,
		AcceptOnlyNeuteredMales: output.AcceptOnlyNeuteredMales,
		AcceptFemales:           output.AcceptFemales,
		PetWeight:               output.PetWeight,
		AcceptFemaleInHeat:      output.AcceptFemaleInHeat,
		AcceptPuppies:           output.AcceptPuppies,
		AcceptMales:             output.AcceptMales,
	}

	return &adapted, nil
}
