package serverstatus

import (
	"../bot"
	"../config"
	"github.com/anvie/port-scanner"
	"time"
)

func Start() {
	go loop()
}

func loop() {
	prevServerUp := true
	elysiumPvP := portscanner.NewPortScanner("149.202.207.235", time.Second*2)

	for {
		serverUp := elysiumPvP.IsOpen(8099)

		if serverUp && serverUp != prevServerUp {
			sendMessage("@everyone Elysium PVP is now online!")
		} else if !serverUp && serverUp != prevServerUp {
			sendMessage("@everyone Elysium PVP is offline.")
		}

		prevServerUp = serverUp
		time.Sleep(time.Second * 5)
	}
}

func sendMessage(message string) {
	bot.Session.ChannelMessageSend(config.Config.AlertRoomID, message)
}
