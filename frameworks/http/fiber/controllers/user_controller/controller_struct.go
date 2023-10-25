package user_controller

import (
	"pethost/usecases/user_case"
)

type UserController struct {
	usecase *user_case.UserCase
}

func New(usecase *user_case.UserCase) *UserController {
	return &UserController{usecase}
}
