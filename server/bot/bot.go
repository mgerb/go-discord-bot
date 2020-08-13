package bot

import (
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/bothandlers"
	"github.com/mgerb/go-discord-bot/server/config"
	log "github.com/sirupsen/logrus"
)

// keep reference to discord session
var _session *discordgo.Session
var _token string
var _sc chan os.Signal

// SendEmbeddedNotification - sends notification to default room
func SendEmbeddedNotification(title, description string) {
	if _session == nil || config.Config.DefaultRoomID == "" {
		return
	}

	embed := &discordgo.MessageEmbed{
		Color:       0x42adf4,
		Title:       title,
		Description: description,
	}

	_session.ChannelMessageSendEmbed(config.Config.DefaultRoomID, embed)
}

// Start bot - this is a blocking function
func Start(token string) {
	_token = token

	// initialize connection
	_session = connect(token)

	// add bot handlers
	_session.AddHandler(bothandlers.SoundsHandler)
	_session.AddHandler(bothandlers.LoggerHandler)
	_session.AddHandler(func(_s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == config.Config.BotPrefix+"restart" {
			restart()
		}
	})

	// start listening for commands
	// Open the websocket and begin listening.
	err := _session.Open()

	if err != nil {
		log.Error("error opening connection,", err)
		return
	}

	log.Debug("Bot is now running...")
}

func Stop() {
	_session.Close()
}

func restart() {
	if _token == "" {
		log.Warn("Unable to restart - token nil")
		return
	}

	for _, vc := range _session.VoiceConnections {
		vc.Disconnect()
	}

	// https://github.com/bwmarrin/discordgo/issues/759
	// Might need to use this to reconnect
	// err := _session.CloseWithCode(1012)
	err := _session.Close()

	if err != nil {
		log.Error(err)
	}

	time.Sleep(time.Second * 5)

	bothandlers.ActiveConnections = map[string]*bothandlers.AudioConnection{}

	Start(_token)
}

func connect(token string) *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + token)

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
