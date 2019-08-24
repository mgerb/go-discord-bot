package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Sound struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `gorm:"unique" json:"name"`
	Extension string     `json:"extension"`
	UserID    string     `json:"user_id"`
	User      User       `json:"user"`
}

func SoundSave(conn *gorm.DB, sound *Sound) error {
	return conn.Create(sound).Error
}

func SoundGet(conn *gorm.DB) ([]Sound, error) {
	sound := []Sound{}
	err := conn.Set("gorm:auto_preload", true).Find(&sound).Error
	return sound, err
}
