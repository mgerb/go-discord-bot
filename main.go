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
	//read config file
	config.Init()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	if config.Flags.Prod {
		// prod
		gin.SetMode(gin.ReleaseMode)
		// Only log the warning severity or above.
		log.SetLevel(log.WarnLevel)
	} else {
		// debug
		log.SetLevel(log.DebugLevel)
	}
}

func main() {

	// start the bot
	bot.Start(config.Config.Token)

	// start the web server
	webserver.Start()
}
