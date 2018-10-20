package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/bothandlers"
	"github.com/mgerb/go-discord-bot/server/config"
	log "github.com/sirupsen/logrus"
)

var session *discordgo.Session

// Start the bot
func Start(token string) *discordgo.Session {
	// initialize connection
	session := connect(token)

	// add bot handlers
	addHandler(session, bothandlers.SoundsHandler)
	addHandler(session, bothandlers.LoggerHandler)

	// start listening for commands
	startListener(session)

	return session
}

// GetSession - get current discord session
func GetSession() *discordgo.Session {
	return session
}

// SendEmbeddedNotification - sends notification to default room
func SendEmbeddedNotification(title, description string) {
	if session == nil || config.Config.DefaultRoomID == "" {
		return
	}

	embed := &discordgo.MessageEmbed{
		Color:       0x42adf4,
		Title:       title,
		Description: description,
	}

	session.ChannelMessageSendEmbed(config.Config.DefaultRoomID, embed)
}

func addHandler(session *discordgo.Session, handler interface{}) {
	session.AddHandler(handler)
}

func connect(token string) *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	var err error
	session, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Error(err)
		log.Fatal("Error creating Discord session.", err)
	}

	// Get the account information.
	_, err = session.User("@me")

	if err != nil {
		log.Error("Error obtaining account details. Make sure you have the correct bot token.")
		log.Fatal(err)
	}

	log.Debug("Bot connected")

	return session
}

func startListener(session *discordgo.Session) {
	// start new non blocking go routine
	go func() {
		// Open the websocket and begin listening.
		err := session.Open()
		if err != nil {
			log.Error("error opening connection,", err)
			return
		}

		log.Debug("Bot is now running...")

		// Simple way to keep program running until CTRL-C is pressed.
		<-make(chan struct{})
		return
	}()
}
