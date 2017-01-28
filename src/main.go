package main

import (
	"./bot"
	"./config"
	"./handlers"
	"./webserver"
)

func main() {
	//read config file
	config.Init()

	//connect bot to account with token
	bot.Connect(config.Config.Token)

	//add handlers
	bot.AddHandler(handlers.SoundsHandler)

	// start new go routine for the discord websockets
	go bot.Start()

	// start the web server
	webserver.Start()
}
