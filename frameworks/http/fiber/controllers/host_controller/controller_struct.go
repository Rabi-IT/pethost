package host_controller

import (
	"pethost/usecases/host_case"
)

type PetHostController struct {
	usecase host_case.PetHostCase
}

func New(usecase host_case.PetHostCase) *PetHostController {
	return &PetHostController{usecase}
}
