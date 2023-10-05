package pethost_case

import g "pethost/adapters/gateways/pethost_gateway"

type PetHostCase struct {
	gateway g.PetHostGateway
}

func New(gateway g.PetHostGateway) *PetHostCase {
	return &PetHostCase{gateway}
}
