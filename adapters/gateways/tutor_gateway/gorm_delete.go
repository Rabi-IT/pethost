package tutor_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g *GormTutorGatewayAdapter) Delete(id string) (bool, error) {
	result := g.DB.Conn.Where(
		"id = ?", id,
	).Delete(&models.Tutor{})

	return result.RowsAffected > 0, result.Error
}
