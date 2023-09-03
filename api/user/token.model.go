package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id" validate:"uuid"`
	UserID       uuid.UUID `gorm:"type:uuid;not null" json:"user_id" validate:"required,uuid"`
	AccessJwtID  uuid.UUID `gorm:"type:uuid;not null" json:"access_jwt_id" validate:"required,uuid"`
	RefreshJwtID uuid.UUID `gorm:"type:uuid;not null" json:"refresh_jwt_id" validate:"required,uuid"`
}
