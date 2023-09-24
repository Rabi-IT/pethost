package pet_controller

import (
	"pethost/usecases/pet_case"
)

type PetController struct {
	usecase *pet_case.PetCase
}

func New(usecase *pet_case.PetCase) *PetController {
	return &PetController{usecase}
}
