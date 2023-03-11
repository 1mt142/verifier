package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	Otp      string
	Channels string
	SenderId uuid.UUID `gorm:"type:uuid"`
}
