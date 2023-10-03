package host_gateway

import (
	database "pethost/frameworks/database/gorm"
)

type PetHostGateway interface {
	Create(input CreateInput) (string, error)
	Delete(id string) (bool, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	GetByID(id string) (*GetByIDOutput, error)
}

type CreateInput struct {
	Name           string
	Phone          string
	ZIP            string
	SocialID       string
	Email          string
	EmergencyPhone string
	Neighborhood   string
	Street         string
	TaxID          string
	City           string
	State          string
	Complement     string
}

type DeleteInput struct {
	ID string
}

type PaginateFilter struct {
	Name         *string
	City         *string
	State        *string
	Neighborhood *string
}

type PaginateData struct {
	Complement   string
	Name         string
	City         string
	State        string
	ZIP          string
	Street       string
	Neighborhood string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}

type PatchFilter struct {
	ID string
}

type PatchValues struct {
	Name           string
	City           string
	State          string
	Complement     string
	Phone          string
	ZIP            string
	Email          string
	EmergencyPhone string
	Neighborhood   string
	Street         string
}

type GetByIDOutput struct {
	Email          string
	EmergencyPhone string
	Neighborhood   string
	Street         string
	Name           string
	City           string
	State          string
	Complement     string
	ZIP            string
}
