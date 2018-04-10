package bothandlers

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/logger"
)

// LoggerHandler -
func LoggerHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// upsert user
	user := getUser(m.Author)
	user.Save()

	// create and save message
	timestamp, _ := m.Message.Timestamp.Parse()
	editedTimestamp, _ := m.Message.EditedTimestamp.Parse()
	attachments := getAttachments(m.Message.Attachments)

	message := &logger.Message{
		ID:              m.Message.ID,
		ChannelID:       m.Message.ChannelID,
		Content:         m.Message.Content,
		Timestamp:       strconv.Itoa(int(timestamp.Unix())),
		EditedTimestamp: strconv.Itoa(int(editedTimestamp.Unix())),
		MentionRoles:    strings.Join(m.Message.MentionRoles, ","),
		Tts:             m.Message.Tts,
		MentionEveryone: m.Message.MentionEveryone,
		UserID:          m.Author.ID,
		Attachments:     attachments,
	}

	message.Save()
}

func getAttachments(att []*discordgo.MessageAttachment) []logger.Attachment {
	var attachments []logger.Attachment
	for _, a := range att {
		newAttachment := logger.Attachment{
			MessageID: a.ID,
			Filename:  a.Filename,
			Height:    a.Height,
			ProxyURL:  a.ProxyURL,
			Size:      a.Size,
			URL:       a.URL,
			Width:     a.Width,
		}
		attachments = append(attachments, newAttachment)
	}
	return attachments
}

func getUser(u *discordgo.User) *logger.User {
	return &logger.User{
		ID:            u.ID,
		Email:         u.Email,
		Username:      u.Username,
		Avatar:        u.Avatar,
		Discriminator: u.Discriminator,
		Token:         u.Token,
		Verified:      u.Verified,
		MFAEnabled:    u.MFAEnabled,
		Bot:           u.Bot,
	}
}
