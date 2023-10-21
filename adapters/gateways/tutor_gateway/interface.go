package tutor_gateway

import (
	database "pethost/frameworks/database/gorm"
)

type TutorGateway interface {
	Create(input CreateInput) (string, error)
	GetByID(id string) (*GetByIDOutput, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	List(input ListInput) ([]ListOutput, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
	Delete(id string) (bool, error)
}

type CreateInput struct {
	State          string
	ZIP            string
	Phone          string
	City           string
	Photo          string
	TaxID          string
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
	Name           string
	Email          string
	Neighborhood   string
}

type GetByIDOutput struct {
	Phone          string
	City           string
	State          string
	ZIP            string
	Name           string
	Email          string
	Photo          string
	TaxID          string
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
}

type PatchFilter struct {
	ID             *string
	ZIP            *string
	Phone          *string
	City           *string
	State          *string
	TaxID          *string
	SocialID       *string
	Street         *string
	Complement     *string
	EmergencyPhone *string
	Name           *string
	Email          *string
	Photo          *string
}

type PatchValues struct {
	ZIP            string
	Phone          string
	City           string
	State          string
	TaxID          string
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
	Name           string
	Email          string
	Photo          string
}

type ListInput struct {
	EmergencyPhone *string
	Name           *string
	Email          *string
	Photo          *string
	TaxID          *string
	SocialID       *string
	Street         *string
	Complement     *string
	Phone          *string
	City           *string
	State          *string
	ZIP            *string
}

type ListOutput struct {
	EmergencyPhone string
	Name           string
	Email          string
	Photo          string
	TaxID          string
	SocialID       string
	Street         string
	Complement     string
	Phone          string
	City           string
	State          string
	ZIP            string
}

type PaginateFilter struct {
	State          *string
	ZIP            *string
	Phone          *string
	City           *string
	Photo          *string
	TaxID          *string
	SocialID       *string
	Street         *string
	Complement     *string
	EmergencyPhone *string
	Name           *string
	Email          *string
}

type PaginateData struct {
	Photo string
	Name  string
	State string
	City  string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}
