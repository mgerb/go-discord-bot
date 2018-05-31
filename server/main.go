package main

import (
	"os"

	"github.com/mgerb/go-discord-bot/server/bot"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/logger"
	"github.com/mgerb/go-discord-bot/server/webserver"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	//read config file
	config.Init()

	if config.Config.Logger {
		db.Init()
		db.Conn.AutoMigrate(&logger.Message{}, &logger.Attachment{}, &logger.User{})
	}
}

func main() {

	// start the bot
	bot.Start(config.Config.Token)

	// start the web server
	webserver.Start()
}
