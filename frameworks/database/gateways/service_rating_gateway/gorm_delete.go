package service_rating_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g *GormServiceRatingGatewayAdapter) Delete(id string) (bool, error) {
	result := g.DB.Conn.Where(
		"id = ?", id,
	).Delete(&models.ServiceRating{})

	return result.RowsAffected > 0, result.Error
}
