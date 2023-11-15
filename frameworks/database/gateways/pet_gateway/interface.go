package pet_gateway

import "pethost/usecases/pet_case/pet"

type PetGateway interface {
	Create(input CreateInput) (string, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	List(input ListInput) ([]ListOutput, error)
	Delete(id string) (bool, error)
	GetByID(id string) (*GetByIDOutput, error)
}

type CreateInput struct {
	UserID    string
	Species   string
	Name      string
	Breed     string
	Birthdate string
	Gender    pet.Gender
	Weight    uint8
	Neutered  bool
}

type PatchFilter struct {
	ID        *string
	Species   *string
	Name      *string
	Breed     *string
	Birthdate *string
	Gender    *pet.Gender
	Weight    *uint8
}

type PatchValues struct {
	Species    string
	Name       string
	Breed      string
	Birthdate  string
	Gender     pet.Gender
	Weight     uint8
	Neutered   *bool
	Vaccinated *bool
}

type ListInput struct {
	TutorID *string
}

type ListOutput struct {
	Name      string
	Breed     string
	Birthdate string
	Gender    pet.Gender
	Weight    uint8
	Species   string
}

type GetByIDOutput struct {
	Name       string
	Breed      string
	Birthdate  string
	Gender     pet.Gender
	Weight     uint8
	Species    string
	Neutered   bool
	Vaccinated bool
}
