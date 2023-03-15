package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	Name        string
	Address     string
	CompanyType string
	Users       []*User `gorm:"many2many:user_services"`
}
