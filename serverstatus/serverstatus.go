package serverstatus

import (
	"../bot"
	"../config"
	"github.com/anvie/port-scanner"
	"time"
)

const serverAddr string = "149.202.207.235"

func Start() {
	go loop()
}

func loop() {
	prevServerUp := true
	elysiumPvP := portscanner.NewPortScanner(serverAddr, time.Second*2)

	for {
		serverUp := elysiumPvP.IsOpen(8099)

		if serverUp && serverUp != prevServerUp {
			sendMessage("@here Elysium PVP is now online!")
		} else if !serverUp && serverUp != prevServerUp {
			sendMessage("@here Elysium PVP is now offline!")
		}

		prevServerUp = serverUp
		time.Sleep(time.Second * 5)
	}
}

func sendMessage(message string) {
	bot.Session.ChannelMessageSend(config.Config.AlertRoomID, message)
}
