package main

import (
	"./bot"
	"./bothandlers"
	"./config"
	"./webserver"
)

func main() {
	//read config file
	config.Init()

	//connect bot to account with token
	bot.Connect(config.Config.Token)

	//add handlers
	bot.AddHandler(bothandlers.SoundsHandler)
	bot.AddHandler(bothandlers.GifHandler)

	// start new go routine for the discord websockets
	go bot.Start()

	// start the web server
	webserver.Start()
}
