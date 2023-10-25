package user_case

import g "pethost/adapters/gateways/user_gateway"

type UserCase struct {
	gateway g.UserGateway
}

func New(gateway g.UserGateway) *UserCase {
	return &UserCase{gateway}
}
