package model

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

// User -
type User struct {
	ID             string     `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	Email          string     `json:"email"`
	Username       string     `json:"username"`
	Avatar         string     `json:"avatar"`
	Discriminator  string     `json:"discriminator"`
	Token          string     `gorm:"-" json:"token"`
	Verified       bool       `json:"verified"`
	MFAEnabled     bool       `json:"mfa_enabled"`
	Bot            bool       `json:"bot"`
	Permissions    *int       `gorm:"default:1;not null" json:"permissions"`
	VoiceJoinSound *string    `json:"voice_join_sound"` // sound clip that plays when user joins channel
}

// UserSave -
func UserSave(conn *gorm.DB, u *User) error {
	var userCopy User
	copier.Copy(&userCopy, u)
	// insert or update user
	// need to make copy of assign object because it must mess
	// with the actual object in FirstOrCreate method
	return conn.Where(&User{ID: u.ID}).Assign(userCopy).FirstOrCreate(u).Error
}

// UserGet - get user by id
func UserGet(conn *gorm.DB, id string) (*User, error) {
	user := &User{ID: id}
	err := conn.First(user).Error
	return user, err
}

func UserGetAll(conn *gorm.DB) (*[]User, error) {
	users := &[]User{}
	err := conn.Find(users).Error
	return users, err
}

func UserUpdate(conn *gorm.DB, user *User) (*User, error) {
	err := conn.Save(user).Error
	return user, err
}
