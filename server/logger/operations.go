package logger

import (
	"regexp"
	"time"

	"github.com/mgerb/go-discord-bot/server/db"
)

const urlRegexp = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`

var linkedPostsCacheTimeout time.Time
var linkedPostsCache map[string]int

// GetMessages - returns all messages - must use paging
func GetMessages(page int) ([]Message, error) {
	messages := []Message{}
	err := db.Conn.Offset(page*100).Limit(100).Order("timestamp desc", true).Preload("User").Find(&messages).Error
	return messages, err
}

// GetLinkedMessages - get count of discord comments that contain URL's - per user
// cached for 10 minutes because there is a lot of data filtering
func GetLinkedMessages() (map[string]int, error) {

	if linkedPostsCacheTimeout.After(time.Now().Add(-10 * time.Minute)) {
		return linkedPostsCache, nil
	}

	result := []map[string]interface{}{}
	rows, err := db.Conn.Table("messages").
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
