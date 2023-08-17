package models

import (
	"pomo/internal/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name             string    `json:"name" gorm:"not null" validate:"required"`
	Email            string    `json:"email" gorm:"unique_index" validate:"required,email"`
	Password         string    `json:"-" gorm:"not null" validate:"required,min=6,max=100"`
	Image            string    `json:"image" gorm:"default:'static/images/defaults/user.png'"`
	Provider         string    `json:"provider" gorm:"default:'email'"`
	Role             string    `json:"role" gorm:"type:varchar(255);not null;default:'user'"` // user, admin
	VerificationCode string    `json:"-" gorm:"type:varchar(99)"`
	Verified         bool      `json:"verified" gorm:"not null;default:false"`
	CreatedAt        time.Time `json:"-" gorm:"not null"`
	UpdatedAt        time.Time `json:"-" gorm:"not null"`
}

func (entity *UserModel) BeforeCreate(db *gorm.DB) error {
	entity.Password = utils.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *UserModel) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
