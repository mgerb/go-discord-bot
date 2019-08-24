package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserEventLog - logger for user events
type UserEventLog struct {
	ID        uint       `gorm:"primary_key; auto_increment; not null" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Content   string     `json:"content"`
	User      User       `json:"user"`
	UserID    string     `json:"user_id"`
}

// UserEventLogSave -
func UserEventLogSave(conn *gorm.DB, m *UserEventLog) error {
	return conn.Save(m).Error
}

// UserEventLogGet - returns all messages - must use paging
func UserEventLogGet(conn *gorm.DB, page int) ([]*UserEventLog, error) {
	userEventLog := []*UserEventLog{}
	err := conn.Offset(page*100).Limit(100).Order("created_at desc", true).Preload("User").Find(&userEventLog).Error
	return userEventLog, err
}

// LogSoundPlayedEvent - log event when user plays sound clip
func LogSoundPlayedEvent(conn *gorm.DB, userID, soundName string, fromWebUI bool) error {

	var content string

	// from discord
	if !fromWebUI {
		content = "played sound clip: " + soundName
	} else {
		content = "played sound clip from web UI: " + soundName
	}

	// log play event
	userEventLog := &UserEventLog{
		UserID:  userID,
		Content: content,
	}

	return UserEventLogSave(conn, userEventLog)
}
