package pubg

/**
 * DEPRECATED
 * I no longer have a use for this so I'm ripping it out
 */

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	pubgClient "github.com/mgerb/go-pubg"
	log "github.com/sirupsen/logrus"
)

var (
	apiKey string
	stats  = map[string]*pubgClient.Player{}
	mut    = &sync.Mutex{}
)

// Start -
func Start(key string, players []string) {

	apiKey = key

	log.Debug("Gathering pubg data...")

	go fetchStats(players)
}

func fetchStats(players []string) {

	api, err := pubgClient.New(apiKey)

	if err != nil {
		log.Fatal(err)
	}

	// fetch new stats every 30 seconds
	for {

		for _, player := range players {
			newStats, err := api.GetPlayer(player)

			if err != nil {
				log.Error(err)
				continue
			}

			mut.Lock()
			stats[player] = newStats
			mut.Unlock()

			time.Sleep(time.Second * 2)
		}

		time.Sleep(time.Second * 30)
	}

}

// Handler - returns the pubg stats
func Handler(c *gin.Context) {
	c.JSON(200, stats)
}
