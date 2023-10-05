package pethost_case_test

import (
	"net/http"
	"pethost/adapters/gateways/pethost_gateway"
	"pethost/fixtures"
	"pethost/usecases/pethost_case"
	"testing"

	"github.com/stretchr/testify/require"
)

const baseURI = "/pethost/"

func Test_Integration_should_create(t *testing.T) {
	fixtures.CleanDatabase()

	Body := pethost_case.CreateInput{
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
		URI:      baseURI,
		Response: &id,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)
}

func Test_Integration_should_be_able_to_retrive_by_id(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.Pethost.Create(t, nil)

	responseBody := pethost_gateway.GetByIDOutput{}

	fixtures.Get(t, fixtures.GetInput{
		URI:      baseURI + id,
		Response: &responseBody,
	})

	EXPECTED := pethost_gateway.GetByIDOutput{
		Email:          "Email",
		EmergencyPhone: "EmergencyPhone",
		Neighborhood:   "Neighborhood",
		Street:         "Street",
		Name:           "Name",
		City:           "City",
		State:          "State",
		Complement:     "Complement",
		ZIP:            "ZIP",
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_paginate(t *testing.T) {
	fixtures.CleanDatabase()

	for i := 0; i < 5; i++ {
		fixtures.Pethost.Create(t, nil)
	}

	responseBody := pethost_gateway.PaginateOutput{}
	statusCode := fixtures.Get(t, fixtures.GetInput{
		URI:      baseURI,
		Response: &responseBody,
	})

	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := pethost_gateway.PaginateOutput{
		Data:     []pethost_gateway.PaginateData{},
		MaxPages: 1,
	}
	for i := 0; i < 5; i++ {
		EXPECTED.Data = append(EXPECTED.Data, pethost_gateway.PaginateData{
			Complement:   "Complement",
			Name:         "Name",
			City:         "City",
			State:        "State",
			ZIP:          "ZIP",
			Street:       "Street",
			Neighborhood: "Neighborhood",
		})
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_update(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.Pethost.Create(t, nil)

	Body := pethost_case.PatchValues{
		ZIP:            "NewZIP",
		Phone:          "NewPhone",
		Email:          "NewEmail",
		EmergencyPhone: "NewEmergencyPhone",
		Neighborhood:   "NewNeighborhood",
		Street:         "NewStreet",
		SocialID:       "NewSocialID",
		TaxID:          "NewTaxID",
		City:           "NewCity",
		State:          "NewState",
		Complement:     "NewComplement",
		Name:           "NewName",
	}

	ok := fixtures.Patch(t, fixtures.PatchInput{
		Body: Body,
		URI:  baseURI + id,
	})
	require.True(t, ok == "OK")

	found, statusCode := fixtures.Pethost.GetByID(t, id)
	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := pethost_gateway.GetByIDOutput{
		ZIP:            "NewZIP",
		Email:          "NewEmail",
		EmergencyPhone: "NewEmergencyPhone",
		Neighborhood:   "NewNeighborhood",
		Street:         "NewStreet",
		City:           "NewCity",
		State:          "NewState",
		Complement:     "NewComplement",
		Name:           "NewName",
	}

	require.Equal(t, EXPECTED, found)
}

func Test_Integration_should_be_able_to_delete(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.Pethost.Create(t, nil)

	respBody, statusCode := fixtures.Delete(t, fixtures.DeleteInput{
		URI: baseURI + id,
	})

	require.Equal(t, statusCode, http.StatusNoContent)
	require.Empty(t, respBody)

	found, statusCode := fixtures.Pethost.GetByID(t, id)
	require.Equal(t, statusCode, http.StatusNotFound)

	EXPECTED := pethost_gateway.GetByIDOutput{}

	require.Equal(t, EXPECTED, found)
}
