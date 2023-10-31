package schedule_case

import g "pethost/frameworks/database/gateways/schedule_gateway"

type ScheduleCase struct {
	gateway g.ScheduleGateway
}

func New(gateway g.ScheduleGateway) ScheduleCase {
	return ScheduleCase{gateway}
}
