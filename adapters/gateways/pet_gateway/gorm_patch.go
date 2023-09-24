package pet_gateway

import (
	"pethost/frameworks/database/gorm/models"
)

func (g GormPetGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	result := g.DB.Conn.Model(&models.Pet{}).Select("*").Updates(newValues)
	return result.RowsAffected > 0, result.Error
}
