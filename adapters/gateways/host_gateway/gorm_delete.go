package host_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g *GormPetHostGatewayAdapter) Delete(id string) (bool, error) {
	result := g.DB.Conn.Where(
		"id = ?", id,
	).Delete(&models.PetHost{})

	return result.RowsAffected > 0, result.Error
}
