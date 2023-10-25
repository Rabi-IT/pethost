package user_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g GormUserGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	query := g.DB.Conn.Model(&models.User{})

	if filter.ID != nil {
		query = query.Where("id = ?", filter.ID)
	} else {

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

		if filter.Phone != nil {
			query = query.Where("phone = ?", filter.Phone)
		}

		if filter.City != nil {
			query = query.Where("city = ?", filter.City)
		}

		if filter.State != nil {
			query = query.Where("state = ?", filter.State)
		}

		if filter.ZIP != nil {
			query = query.Where("zip = ?", filter.ZIP)
		}

	}

	result := query.Updates(newValues)
	return result.RowsAffected > 0, result.Error
}
