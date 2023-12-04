package service_rating_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormServiceRatingGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	query := g.DB.Conn.Model(&models.ServiceRating{})

	if filter.ID != nil {
		query = query.Where("id = ?", filter.ID)
	}

	if filter.TutorID != nil {
		query = query.Where("tutor_id = ?", filter.TutorID)
	}

	result := query.Updates(newValues)
	return result.RowsAffected > 0, result.Error
}
