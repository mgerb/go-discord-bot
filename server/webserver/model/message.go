package model

import (
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
)

// Message - discord message
type Message struct {
	ID              string       `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	DeletedAt       *time.Time   `json:"deleted_at"`
	ChannelID       string       `json:"channel_id"`
	Content         string       `json:"content"`
	Timestamp       time.Time    `json:"timestamp"`
	EditedTimestamp time.Time    `json:"edited_timestamp"`
	MentionRoles    string       `json:"mention_roles"`
	TTS             bool         `json:"tts"`
	MentionEveryone bool         `json:"mention_everyone"`
	User            User         `json:"user"`
	UserID          string       `json:"user_id"`
	Attachments     []Attachment `json:"attachments"`
}

// MessageSave -
func MessageSave(conn *gorm.DB, m *Message) error {
	return conn.Save(m).Error
}

const urlRegexp = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`

var linkedPostsCacheTimeout time.Time
var linkedPostsCache map[string]int

// MessageGet - returns all messages - must use paging
func MessageGet(conn *gorm.DB, page int) ([]Message, error) {
	messages := []Message{}
	err := conn.Offset(page*100).Limit(100).Order("timestamp desc", true).Preload("User").Find(&messages).Error
	return messages, err
}

// MessageGetLinked - get count of discord comments that contain URL's - per user
// cached for 10 minutes because there is a lot of data filtering
func MessageGetLinked(conn *gorm.DB) (map[string]int, error) {

	if linkedPostsCacheTimeout.After(time.Now().Add(-10 * time.Minute)) {
		return linkedPostsCache, nil
	}

	result := []map[string]interface{}{}
	rows, err := conn.Table("messages").
		Select("users.username, messages.content").
		Joins("join users on messages.user_id = users.id").
		Rows()

	if err != nil {
		return map[string]int{}, err
	}

	for rows.Next() {
		var username, content string
		rows.Scan(&username, &content)
		result = append(result, map[string]interface{}{
			"username": username,
			"content":  content,
		})
	}

	linkedPostsCacheTimeout = time.Now()
	linkedPostsCache = groupPosts(result)

	return linkedPostsCache, nil
}

// group posts by user and count
func groupPosts(posts []map[string]interface{}) map[string]int {

	result := map[string]int{}

	for _, p := range posts {
		match, _ := regexp.MatchString(urlRegexp, p["content"].(string))

		if match {
			if _, ok := result[p["username"].(string)]; ok {
				result[p["username"].(string)]++
			} else {
				result[p["username"].(string)] = 1
			}
		}
	}

	return result
}
