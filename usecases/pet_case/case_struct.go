package pet_case

import g "pethost/frameworks/database/gateways/pet_gateway"

type PetCase struct {
	gateway g.PetGateway
}

func New(gateway g.PetGateway) *PetCase {
	useCase := PetCase{gateway}
	return &useCase
}
