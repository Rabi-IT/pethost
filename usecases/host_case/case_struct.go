package host_case

import g "pethost/adapters/gateways/host_gateway"

type PetHostCase struct {
	gateway g.PetHostGateway
}

func New(gateway g.PetHostGateway) PetHostCase {
	return PetHostCase{gateway}
}
