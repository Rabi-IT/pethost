package preference_case

import g "pethost/frameworks/database/gateways/preference_gateway"

type PreferenceCase struct {
	gateway g.PreferenceGateway
}

func New(gateway g.PreferenceGateway) *PreferenceCase {
	return &PreferenceCase{gateway}
}
