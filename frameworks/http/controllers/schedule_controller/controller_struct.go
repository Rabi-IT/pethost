package schedule_controller

import (
	"pethost/usecases/schedule_case"
)

type ScheduleController struct {
	usecase schedule_case.ScheduleCase
}

func New(usecase schedule_case.ScheduleCase) *ScheduleController {
	return &ScheduleController{usecase}
}
