package main

import (
	"./bot"
	"./config"
	"./serverstatus"
)

// Variables used for command line parameters
var (
	BotID string
)

func main() {
	config.Configure()
	bot.Connect(config.Config.Token)
	serverstatus.Start()
	bot.Start()

}
