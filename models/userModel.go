package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles []string

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	Username   string
	Email      string `gorm:"unique"`
	Password   string
	IsActive   bool  `gorm:"default:false"`
	IsVerified bool  `gorm:"default:false"`
	Roles      Roles `gorm:"serializer:json"`
}
