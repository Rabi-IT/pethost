package preference_controller

import (
	"pethost/usecases/preference_case"
)

type PreferenceController struct {
	usecase *preference_case.PreferenceCase
}

func New(usecase *preference_case.PreferenceCase) *PreferenceController {
	return &PreferenceController{usecase}
}
