package service_rating_controller

import (
	"pethost/usecases/service_rating_case"
)

type ServiceRatingController struct {
	usecase service_rating_case.ServiceRatingCase
}

func New(usecase service_rating_case.ServiceRatingCase) *ServiceRatingController {
	return &ServiceRatingController{usecase}
}
