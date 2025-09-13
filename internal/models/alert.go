package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Alert struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	Symbol    string    `gorm:"type:varchar(10);not null" json:"symbol" validate:"required"`
	Message   string    `gorm:"type:text;not null" json:"message" validate:"required"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type" validate:"required"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Alert) Validate() error {
	return validator.New().Struct(a)
}
