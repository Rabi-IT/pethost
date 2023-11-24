package service_rating_case

import g "pethost/frameworks/database/gateways/service_rating_gateway"

type ServiceRatingCase struct {
	gateway g.ServiceRatingGateway
}

func New(gateway g.ServiceRatingGateway) ServiceRatingCase {
	return ServiceRatingCase{gateway}
}
