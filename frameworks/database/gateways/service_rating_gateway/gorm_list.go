package service_rating_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormServiceRatingGatewayAdapter) List(filter ListInput) ([]ListOutput, error) {
	query := g.DB.Conn.Model(&models.ServiceRating{})

	if filter.TutorID != nil {
		query = query.Where("tutor_id = ?", filter.TutorID)
	}

	if filter.ScheduleID != nil {
		query = query.Where("schedule_id = ?", filter.ScheduleID)
	}

	if filter.Date != nil {
		query = query.Where("date = ?", filter.Date)
	}

	if filter.Rating != nil {
		query = query.Where("rating = ?", filter.Rating)
	}

	if filter.Comment != nil {
		query = query.Where("comment = ?", filter.Comment)
	}

	output := []ListOutput{}
	result := query.Scan(&output)

	return output, result.Error
}
