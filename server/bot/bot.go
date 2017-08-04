package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
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
		log.Println(err)
		log.Fatal("Error creating Discord session.", err)
	}

	// Get the account information.
	u, err := Session.User("@me")

	if err != nil {
		log.Println("Error obtaining account details. Make sure you have the correct bot token.")
		log.Fatal(err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	fmt.Println("Bot connected")
}

// Start - blocking function that starts a websocket listenting for discord callbacks
func Start() {

	// start new non blocking go routine
	go func() {
		// Open the websocket and begin listening.
		err := Session.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
			return
		}

		fmt.Println("Bot is now running...")

		// Simple way to keep program running until CTRL-C is pressed.
		<-make(chan struct{})
		return
	}()
}

func AddHandler(handler interface{}) {
	Session.AddHandler(handler)
}
