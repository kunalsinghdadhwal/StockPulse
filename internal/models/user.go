package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string      `gorm:"type:varchar(100);unique;not null"json:"email" validate:"required,email"`
	PasswordHash  string      `gorm:"type:varchar(255);not null"json:"-"`
	OAuthProvider string      `grom:"type:varchar(50)" json:"-"`
	OAuthID       string      `gorm:"type:varchar(255);unique" json:"-"`
	OAuthToken    string      `gorm:"type:text" json:"-"`
	Name          string      `gorm:"type:varchar(100)" json:"name" validate:"omitempty,min=5,max=100"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Watchlists    []Watchlist `gorm:"foreignKey:UserID" json:"watchlists,omitempty"`
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
