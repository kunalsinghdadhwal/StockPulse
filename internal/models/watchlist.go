package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Watchlist struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	Symbol    string    `gorm:"type:varchar(10);not null" json:"symbol" validate:"required,uppercase, max=10"`
	Threshold float64   `gorm:"type:double precision;not null" json:"threshold" validate:"required,gt=0"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (w *Watchlist) Validator() error {
	return validator.New().Struct(w)
}
