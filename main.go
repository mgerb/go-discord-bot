package main

import (
	"github.com/mgerb/go-discord-bot/server/bot"
	"github.com/mgerb/go-discord-bot/server/bothandlers"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver"
)

func main() {
	//read config file
	config.Init()

	//connect bot to account with token
	bot.Connect(config.Config.Token)

	//add handlers
	bot.AddHandler(bothandlers.SoundsHandler)

	// start the bot
	bot.Start()

	// start the web server
	webserver.Start()
}
