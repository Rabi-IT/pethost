package pet_gateway

import (
	"pethost/usecases/pet_case/pet"
	"time"
)

type PetGateway interface {
	Create(input CreateInput) (string, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	List(input ListInput) ([]ListOutput, error)
	Delete(id string) (bool, error)
	GetByFilter(filter GetByFilterInput) (*GetByFilterOutput, error)
}

type CreateInput struct {
	UserID     string
	Species    string
	Name       string
	Breed      string
	Birthdate  time.Time
	Gender     pet.Gender
	Weight     uint8
	Neutered   bool
	Vaccinated bool
}

type PatchFilter struct {
	ID      *string
	Species *string
	Name    *string
	Breed   *string
	Gender  *pet.Gender
	Weight  *uint8
}

type PatchValues struct {
	Species    string
	Name       string
	Breed      string
	Birthdate  time.Time
	Gender     pet.Gender
	Weight     uint8
	Neutered   *bool
	Vaccinated *bool
}

type ListInput struct {
	TutorID *string
	PetIDs  []string
}

type ListOutput struct {
	ID         string
	Name       string
	Breed      string
	Birthdate  time.Time
	Gender     pet.Gender
	Weight     uint8
	Species    string
	Neutered   bool
	Vaccinated bool
}

type GetByFilterInput struct {
	ID      string
	TutorID *string
}

type GetByFilterOutput struct {
	Name       string
	Breed      string
	Birthdate  time.Time
	Gender     pet.Gender
	Weight     uint8
	Species    string
	Neutered   bool
	Vaccinated bool
}
