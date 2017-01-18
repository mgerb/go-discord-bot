package main

import (
	"./bot"
	"./config"
	"./handlers"
)

// Variables used for command line parameters
var (
	BotID string
)

func main() {
	//read config file
	config.Configure()

	//connect bot to account with token
	bot.Connect(config.Config.Token)

	//load sound files into memory
	handlers.LoadSounds()

	//add handlers
	bot.AddHandler(handlers.SoundsHandler)

	//start websock to listen for messages
	bot.Start()
}
