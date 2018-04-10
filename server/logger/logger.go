package logger

import (
	"time"

	"github.com/mgerb/go-discord-bot/server/db"
)

// Message - discord message
type Message struct {
	ID              string       `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	DeletedAt       *time.Time   `json:"deleted_at"`
	ChannelID       string       `json:"channel_id"`
	Content         string       `json:"content"`
	Timestamp       string       `json:"timestamp"`
	EditedTimestamp string       `json:"edited_timestamp"`
	MentionRoles    string       `json:"mention_roles"`
	Tts             bool         `json:"tts"`
	MentionEveryone bool         `json:"mention_everyone"`
	User            User         `json:"user"`
	UserID          string       `json:"user_id"`
	Attachments     []Attachment `json:"attachments"`
}

// Save -
func (m *Message) Save() error {
	return db.Conn.Save(m).Error
}

// Attachment - discord message attachment
type Attachment struct {
	MessageID string `gorm:"primary_key" json:"id"`
	URL       string `json:"url"`
	ProxyURL  string `json:"proxy_url"`
	Filename  string `json:"filename"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Size      int    `json:"size"`
}

// User -
type User struct {
	ID            string `gorm:"primary_key" json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	Token         string `json:"token"`
	Verified      bool   `json:"verified"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	Bot           bool   `json:"bot"`
}

// Save -
func (u *User) Save() error {
	return db.Conn.Save(u).Error
}
