package service_rating_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"

	"github.com/google/uuid"
)

func (g *GormServiceRatingGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.ServiceRating{
		ID:      id,
		TutorID: input.TutorID,
		Date:    input.Date,
		Rating:  input.Rating,
		Comment: input.Comment,
	})

	return id, result.Error
}
