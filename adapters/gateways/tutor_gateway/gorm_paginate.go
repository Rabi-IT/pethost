package tutor_gateway

import (
	"pethost/adapters/database"
	gorm_adapter "pethost/frameworks/database/gorm"
	"pethost/frameworks/database/gorm/models"
)

func (g GormTutorGatewayAdapter) Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error) {
	data := []PaginateData{}

	query := g.DB.Conn.Model(&models.Tutor{})

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

	if filter.Email != nil {
		query = query.Where("email = ?", filter.Email)
	}

	if filter.ZIP != nil {
		query = query.Where("zip = ?", filter.ZIP)
	}

	if filter.City != nil {
		query = query.Where("city = ?", filter.City)
	}

	if filter.State != nil {
		query = query.Where("state = ?", filter.State)
	}

	var count int64
	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return &PaginateOutput{
			Data:     data,
			MaxPages: 0,
		}, nil
	}

	gorm_adapter.Paginate(query, paginate)

	result = query.Scan(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	output := &PaginateOutput{
		Data:     data,
		MaxPages: database.CalcMaxPages(count, paginate.PageSize),
	}

	return output, nil
}
