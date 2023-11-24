package schedule_gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"pethost/frameworks/database/gateways/schedule_gateway/ports"
	"pethost/frameworks/database/gorm_adapter/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormScheduleGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	query := g.DB.Conn.Model(&models.Schedule{}).Where("id = ?", filter.ID)

	if filter.Status != nil {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.StatusOR != nil {
		query = query.Where("status IN ?", filter.StatusOR)
	}

	if filter.TutorID != nil {
		query = query.Where("tutor_id = ?", filter.TutorID)
	}

	if filter.HostID != nil {
		query = query.Where("host_id = ?", filter.HostID)
	}

	mapValues := map[string]interface{}{
		"status": newValues.Status,
	}

	if newValues.History != nil {
		mapValues["history"] = PushHistory{newValues.History}
	}

	result := query.Updates(mapValues)
	return result.RowsAffected > 0, result.Error
}

type PushHistory struct {
	*ports.ScheduleHistory
}

func (p PushHistory) GormDataType() string {
	return "jsonb"
}

func (p PushHistory) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	value := fmt.Sprintf(
		`{"userId":"%s","newStatus":"%s","date":"%s","notes":"%s"}`,
		p.UserID, p.NewStatus, p.Date.Format(time.RFC3339), p.Notes,
	)

	return clause.Expr{
		SQL:  `history || ?::jsonb`,
		Vars: []interface{}{value},
	}
}

func (p *PushHistory) Scan(v interface{}) error {
	if v, ok := v.([]byte); ok {
		return json.Unmarshal(v, &p)
	}

	return nil
}
