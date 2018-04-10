package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/bot"
	"github.com/mgerb/go-discord-bot/server/bothandlers"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/db"
)

const everyoneChannel = "101198129352691712"

// this is a script to go through chat history and log old message into database
func main() {
	config.Init()
	db.Init()
	session := bot.Start(config.Config.Token)
	fetchMessages(session, "")
}

func fetchMessages(s *discordgo.Session, beforeID string) {

	messages, err := s.ChannelMessages(everyoneChannel, 100, beforeID, "", "")

	log.Println("Fetching new messages: " + messages[0].Timestamp)

	if err != nil {
		log.Fatal(err)
	}

	for _, m := range messages {
		messageCreate := &discordgo.MessageCreate{
			Message: m,
		}
		bothandlers.LoggerHandler(s, messageCreate)
	}

	if len(messages) == 100 {
		fetchMessages(s, messages[len(messages)-1].ID)
	}
}
