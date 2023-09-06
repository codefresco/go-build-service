package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCredentials struct {
	Email    string `gorm:"type:varchar(128);not null;unique" json:"email" validate:"required,email,min=4,max=128"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type UserDetails struct {
	FirstName string `gorm:"type:varchar(128);not null" json:"first_name" validate:"required,min=4,max=128"`
	LastName  string `gorm:"type:varchar(128);not null" json:"last_name" validate:"required,min=4,max=128"`
	Email     string `gorm:"type:varchar(128);not null;unique" json:"email" validate:"required,email,min=4,max=128"`
}

type UserRegisteration struct {
	UserDetails
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type User struct {
	gorm.Model
	UserDetails
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id" validate:"uuid"`
	PasswordHash string    `gorm:"type:varchar(128);not null" json:"password_hash" validate:"required,min=8,max=128"`
	PasswordSalt string    `gorm:"type:varchar(128);not null" json:"password_salt" validate:"required,min=8,max=128"`
}
