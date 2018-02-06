package bot

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	BotID   string
	Session *discordgo.Session
)

func Connect(token string) {
	// Create a new Discord session using the provided bot token.
	var err error
	Session, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Error(err)
		log.Fatal("Error creating Discord session.", err)
	}

	// Get the account information.
	u, err := Session.User("@me")

	if err != nil {
		log.Error("Error obtaining account details. Make sure you have the correct bot token.")
		log.Fatal(err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	log.Debug("Bot connected")
}

// Start - blocking function that starts a websocket listenting for discord callbacks
func Start() {

	// start new non blocking go routine
	go func() {
		// Open the websocket and begin listening.
		err := Session.Open()
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

func AddHandler(handler interface{}) {
	Session.AddHandler(handler)
}
