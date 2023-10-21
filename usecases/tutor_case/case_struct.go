package tutor_case

import g "pethost/adapters/gateways/tutor_gateway"

type TutorCase struct {
	gateway g.TutorGateway
}

func New(gateway g.TutorGateway) *TutorCase {
	return &TutorCase{gateway}
}
