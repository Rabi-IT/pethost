package models

import "pethost/usecases/auth_case/role"

type User struct {
	ID             string `gorm:"type:uuid"`
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
	Name           string `gorm:"not null"`
	Email          string `gorm:"not null"`
	Photo          string
	TaxID          string `gorm:"not null"`
	Phone          string `gorm:"not null"`
	City           string
	State          string
	ZIP            string
	Neighborhood   string
	Role           role.Role
	IsHost         bool
}

func (m User) TableName() string {
	return "users"
}
