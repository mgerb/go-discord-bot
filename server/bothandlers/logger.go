package bothandlers

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
)

// LoggerHandler -
func LoggerHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// upsert user
	user := getUser(m.Author)
	model.UserSave(db.GetConn(), user)

	// create and save message
	timestamp := m.Message.Timestamp
	var editedTimestamp time.Time
	if m.Message.EditedTimestamp != nil {
		editedTimestamp = *m.Message.EditedTimestamp
	}
	attachments := getAttachments(m.Message.Attachments)

	message := &model.Message{
		ID:              m.Message.ID,
		ChannelID:       m.Message.ChannelID,
		Content:         m.Message.Content,
		Timestamp:       timestamp,
		EditedTimestamp: editedTimestamp,
		MentionRoles:    strings.Join(m.Message.MentionRoles, ","),
		TTS:             m.Message.TTS,
		MentionEveryone: m.Message.MentionEveryone,
		UserID:          m.Author.ID,
		Attachments:     attachments,
	}

	model.MessageSave(db.GetConn(), message)
}

func getAttachments(att []*discordgo.MessageAttachment) []model.Attachment {
	var attachments []model.Attachment
	for _, a := range att {
		newAttachment := model.Attachment{
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

func getUser(u *discordgo.User) *model.User {
	return &model.User{
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
