package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/bot"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	//read config file
	config.Init()

	if config.Flags.Prod {
		// prod
		gin.SetMode(gin.ReleaseMode)
		// Only log the warning severity or above.
		log.SetLevel(log.WarnLevel)
	}
}

func main() {

	// start the bot
	bot.Start(config.Config.Token)

	// start the web server
	webserver.Start()
}
