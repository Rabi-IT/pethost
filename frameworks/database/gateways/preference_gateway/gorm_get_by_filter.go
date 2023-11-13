package preference_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g *GormPreferenceGatewayAdapter) GetByFilter(filter *GetByFilterInput) (*GetByFilterOutput, error) {
	output := &GetByFilterOutput{}
	query := g.DB.Conn.Model(
		&models.Preference{},
	).Where(
		"user_id = ?", filter.UserID,
	).Where(
		"pet_weight & ? > 0", filter.PetWeight,
	).Where(
		"days_of_month & ? > 0", filter.DaysOfMonth,
	)

	if filter.AcceptPuppies != nil {
		query = query.Where("accept_puppies = ?", filter.AcceptPuppies)
	}

	if filter.AcceptMales != nil {
		query = query.Where(
			"accept_males = ?", filter.AcceptMales,
		)
	}

	if filter.AcceptFemaleInHeat != nil {
		query = query.Where(
			"accept_female_in_heat = ?", filter.AcceptFemaleInHeat,
		)
	}

	if filter.AcceptElderly != nil {
		query = query.Where(
			"accept_elderly = ?", filter.AcceptElderly,
		)
	}

	if filter.AcceptOnlyNeuteredMales != nil {
		query = query.Where(
			"accept_only_neutered_males = ?", filter.AcceptOnlyNeuteredMales,
		)
	}

	if filter.AcceptFemales != nil {
		query = query.Where(
			"accept_females = ?", filter.AcceptFemales,
		)
	}

	if filter.OnlyVaccinated != nil {
		query = query.Where(
			"only_vaccinated = ?", filter.OnlyVaccinated,
		)
	}

	result := query.Scan(&output)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	adapted := GetByFilterOutput{
		UserID: output.UserID,
	}

	return &adapted, nil
}
