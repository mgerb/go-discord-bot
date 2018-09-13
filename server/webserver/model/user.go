package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User -
type User struct {
	ID            string     `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Email         string     `json:"email"`
	Username      string     `json:"username"`
	Avatar        string     `json:"avatar"`
	Discriminator string     `json:"discriminator"`
	Token         string     `gorm:"-" json:"token"`
	Verified      bool       `json:"verified"`
	MFAEnabled    bool       `json:"mfa_enabled"`
	Bot           bool       `json:"bot"`
}

// UserSave -
func UserSave(conn *gorm.DB, u *User) error {
	return conn.Save(u).Error
}
