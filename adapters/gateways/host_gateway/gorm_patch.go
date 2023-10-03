package host_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g GormPetHostGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	result := g.DB.Conn.Model(&models.PetHost{}).Select("*").Updates(newValues)
	return result.RowsAffected > 0, result.Error
}
