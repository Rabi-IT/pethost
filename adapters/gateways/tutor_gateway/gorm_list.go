package tutor_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g GormTutorGatewayAdapter) List(filter ListInput) ([]ListOutput, error) {
	query := g.DB.Conn.Model(&models.Tutor{})

	if filter.City != nil {
		query = query.Where("city = ?", filter.City)
	}

	if filter.State != nil {
		query = query.Where("state = ?", filter.State)
	}

	if filter.ZIP != nil {
		query = query.Where("zip = ?", filter.ZIP)
	}

	if filter.Phone != nil {
		query = query.Where("phone = ?", filter.Phone)
	}

	if filter.Email != nil {
		query = query.Where("email = ?", filter.Email)
	}

	if filter.Photo != nil {
		query = query.Where("photo = ?", filter.Photo)
	}

	if filter.TaxID != nil {
		query = query.Where("tax_id = ?", filter.TaxID)
	}

	if filter.SocialID != nil {
		query = query.Where("social_id = ?", filter.SocialID)
	}

	if filter.Street != nil {
		query = query.Where("street = ?", filter.Street)
	}

	if filter.Complement != nil {
		query = query.Where("complement = ?", filter.Complement)
	}

	if filter.EmergencyPhone != nil {
		query = query.Where("emergency_phone = ?", filter.EmergencyPhone)
	}

	if filter.Name != nil {
		query = query.Where("name = ?", filter.Name)
	}

	output := []ListOutput{}
	result := query.Scan(&output)

	return output, result.Error
}
