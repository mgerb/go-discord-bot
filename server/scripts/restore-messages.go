package main

/**
This script will fetch all messages for the provided channel and store them in the database.
*/

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/bot"
	"github.com/mgerb/go-discord-bot/server/bothandlers"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/db"
)

// var everyoneChannel = "101198129352691712"
var everyoneChannel string

// this is a script to go through chat history and log old message into database
func restoreMessages(channelID string) {
	config.Init()
	db.Init()
	session := bot.Start(config.Config.Token)
	everyoneChannel = channelID
	fetchMessages(session, "")
}

func fetchMessages(s *discordgo.Session, beforeID string) {

	messages, err := s.ChannelMessages(everyoneChannel, 100, beforeID, "", "")

	log.Print("Fetching new messages: ")
	log.Println(messages[0].Timestamp)

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
