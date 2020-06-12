package main

import (
	"os"

	"github.com/mgerb/go-discord-bot/server/bot"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	log "github.com/sirupsen/logrus"
)

// There's an issue with sqlite currently: https://github.com/mattn/go-sqlite3/issues/803
// Set env variables to fix temporarily
// export CGO_CFLAGS="-g -O2 -Wno-return-local-addr"

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	//read config file
	config.Init()
	db.Init(model.Migrations...)
}

func main() {

	// start the bot
	bot.Start(config.Config.Token)

	// start the web server
	webserver.Start()
}
