package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id" validate:"uuid"`
	FirstName    string    `gorm:"type:varchar(128);not null" json:"first_name" validate:"required,min=4,max=128"`
	LastName     string    `gorm:"type:varchar(128);not null" json:"last_name" validate:"required,min=4,max=128"`
	Email        string    `gorm:"type:varchar(128);not null;unique" json:"email" validate:"required,email,min=4,max=128"`
	PasswordHash string    `gorm:"type:varchar(128);not null" json:"password_hash" validate:"required,min=8,max=128"`
	PasswordSalt string    `gorm:"type:varchar(128);not null" json:"password_salt" validate:"required,min=8,max=128"`
}
