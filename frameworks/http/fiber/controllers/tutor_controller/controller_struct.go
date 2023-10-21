package tutor_controller

import (
	"pethost/usecases/tutor_case"
)

type TutorController struct {
	usecase *tutor_case.TutorCase
}

func New(usecase *tutor_case.TutorCase) *TutorController {
	return &TutorController{usecase}
}
