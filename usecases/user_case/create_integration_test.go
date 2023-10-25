package user_case_test

import (
	"net/http"
	"pethost/adapters/gateways/user_gateway"
	"pethost/fixtures"
	"pethost/usecases/user_case"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Integration_should_create(t *testing.T) {
	fixtures.CleanDatabase()

	Body := user_case.CreateInput{
		Name:           "Name",
		TaxID:          "TaxID",
		City:           "City",
		State:          "State",
		Phone:          "Phone",
		ZIP:            "ZIP",
		SocialID:       "SocialID",
		Email:          "Email",
		EmergencyPhone: "EmergencyPhone",
		Neighborhood:   "Neighborhood",
		Street:         "Street",
		Complement:     "Complement",
	}

	id := ""
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      fixtures.User.URI,
		Response: &id,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)
}

func Test_Integration_should_be_able_to_retrive_by_id(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.User.Create(t, nil)

	responseBody := user_gateway.GetByIDOutput{}

	fixtures.Get(t, fixtures.GetInput{
		URI:      fixtures.User.URI + id,
		Response: &responseBody,
	})

	EXPECTED := user_gateway.GetByIDOutput{
		City:           "City",
		State:          "State",
		ZIP:            "ZIP",
		Name:           "Name",
		Email:          "Email",
		Phone:          "Phone",
		Photo:          "Photo",
		TaxID:          "TaxID",
		SocialID:       "SocialID",
		Street:         "Street",
		Complement:     "Complement",
		EmergencyPhone: "EmergencyPhone",
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_paginate(t *testing.T) {
	fixtures.CleanDatabase()

	for i := 0; i < 5; i++ {
		fixtures.User.Create(t, nil)
	}

	responseBody := user_gateway.PaginateOutput{}
	statusCode := fixtures.Get(t, fixtures.GetInput{
		URI:      fixtures.User.URI,
		Response: &responseBody,
	})

	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := user_gateway.PaginateOutput{
		Data:     []user_gateway.PaginateData{},
		MaxPages: 1,
	}
	for i := 0; i < 5; i++ {
		EXPECTED.Data = append(EXPECTED.Data, user_gateway.PaginateData{
			State: "State",
			City:  "City",
			Photo: "Photo",
			Name:  "Name",
		})
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_update(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.User.Create(t, nil)

	Body := user_case.PatchValues{
		ZIP:            "NewZIP",
		Phone:          "NewPhone",
		Email:          "NewEmail",
		EmergencyPhone: "NewEmergencyPhone",
		Street:         "NewStreet",
		SocialID:       "NewSocialID",
		TaxID:          "NewTaxID",
		City:           "NewCity",
		State:          "NewState",
		Complement:     "NewComplement",
		Name:           "NewName",
		Photo:          "NewPhoto",
	}

	ok := fixtures.Patch(t, fixtures.PatchInput{
		Body: Body,
		URI:  fixtures.User.URI + id,
	})
	require.True(t, ok == "OK")

	found, statusCode := fixtures.User.GetByID(t, id)
	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := user_gateway.GetByIDOutput{
		ZIP:            "NewZIP",
		Phone:          "NewPhone",
		Email:          "NewEmail",
		EmergencyPhone: "NewEmergencyPhone",
		Street:         "NewStreet",
		SocialID:       "NewSocialID",
		TaxID:          "NewTaxID",
		City:           "NewCity",
		State:          "NewState",
		Complement:     "NewComplement",
		Name:           "NewName",
		Photo:          "NewPhoto",
	}

	require.Equal(t, EXPECTED, found)
}

func Test_Integration_should_be_able_to_delete(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.User.Create(t, nil)

	respBody, statusCode := fixtures.Delete(t, fixtures.DeleteInput{
		URI: fixtures.User.URI + id,
	})

	require.Equal(t, statusCode, http.StatusNoContent)
	require.Empty(t, respBody)

	found, statusCode := fixtures.User.GetByID(t, id)
	require.Equal(t, statusCode, http.StatusNotFound)

	EXPECTED := user_gateway.GetByIDOutput{}

	require.Equal(t, EXPECTED, found)
}
