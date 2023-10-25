package user_case_test

import (
	"context"
	"pethost/adapters/gateways/user_gateway"
	"pethost/fixtures/mocks"
	"pethost/usecases/user_case"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func makeSut(g user_gateway.UserGateway) *user_case.UserCase {
	return user_case.New(g)
}

func Test_Unit_Create__should_fail_if_required_fields_are_empty(t *testing.T) {
	DUMMY_GATEWAY := mocks.NewUserGateway(t)
	sut := makeSut(DUMMY_GATEWAY)

	_, err := sut.Create(context.Background(), &user_case.CreateInput{
		Name:           "",
		Photo:          "Photo",
		TaxID:          "",
		City:           "",
		State:          "",
		Phone:          "",
		ZIP:            "",
		SocialID:       "",
		Email:          "",
		EmergencyPhone: "EmergencyPhone",
		Neighborhood:   "",
		Street:         "",
		Complement:     "Complement",
	})

	expectedMsg := []string{
		"Key: 'CreateInput.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		"Key: 'CreateInput.TaxID' Error:Field validation for 'TaxID' failed on the 'required' tag",
		"Key: 'CreateInput.City' Error:Field validation for 'City' failed on the 'required' tag",
		"Key: 'CreateInput.State' Error:Field validation for 'State' failed on the 'required' tag",
		"Key: 'CreateInput.Phone' Error:Field validation for 'Phone' failed on the 'required' tag",
		"Key: 'CreateInput.ZIP' Error:Field validation for 'ZIP' failed on the 'required' tag",
		"Key: 'CreateInput.SocialID' Error:Field validation for 'SocialID' failed on the 'required' tag",
		"Key: 'CreateInput.Email' Error:Field validation for 'Email' failed on the 'required' tag",
		"Key: 'CreateInput.Neighborhood' Error:Field validation for 'Neighborhood' failed on the 'required' tag",
		"Key: 'CreateInput.Street' Error:Field validation for 'Street' failed on the 'required' tag",
	}

	require.Equal(t, strings.Join(expectedMsg, "\n"), err.Error())
}

func Test_Unit_Create__should_not_fail_if_all_optional_fields_are_not_filled_in(t *testing.T) {
	gateway := mocks.NewUserGateway(t)
	expectedID := "ANY_ID"
	gateway.On("Create", mock.Anything).Return(expectedID, nil)
	sut := user_case.New(gateway)

	id, err := sut.Create(context.Background(), &user_case.CreateInput{
		Name:           "Name",
		TaxID:          "TaxID",
		City:           "City",
		State:          "State",
		Phone:          "Phone",
		ZIP:            "ZIP",
		SocialID:       "SocialID",
		Email:          "Email",
		EmergencyPhone: "",
		Neighborhood:   "Neighborhood",
		Street:         "Street",
		Complement:     "",
	})

	require.Nil(t, err)
	require.Equal(t, expectedID, id)
}
