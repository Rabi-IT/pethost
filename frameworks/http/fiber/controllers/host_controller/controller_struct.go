package host_controller

import "pethost/usecases/pethost_case"

type PetHostController struct {
	usecase *pethost_case.PetHostCase
}

func New(usecase *pethost_case.PetHostCase) *PetHostController {
	return &PetHostController{usecase}
}
