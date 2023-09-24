package pet_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g *GormPetGatewayAdapter) Delete(id string) (bool, error) {
	result := g.DB.Conn.Where(
		"id = ?", id,
	).Delete(&models.Pet{})

	return result.RowsAffected > 0, result.Error
}
