package pubg

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/mgerb/chi_auth_server/response"
	pubgClient "github.com/mgerb/go-pubg"
)

var (
	apiKey string
	stats  = map[string]*pubgClient.Player{}
	mut    = &sync.Mutex{}
)

// Start -
func Start(key string, players []string) {

	apiKey = key

	log.Println("Gathering pubg data...")

	go fetchStats(players)
}

func fetchStats(players []string) {

	api := pubgClient.New(apiKey)

	// fetch new stats every 30 seconds
	for {

		for _, player := range players {
			newStats := api.GetPlayer(player)

			mut.Lock()
			stats[player] = newStats
			mut.Unlock()

			time.Sleep(time.Second * 2)
		}

		time.Sleep(time.Second * 30)
	}

}

// Handler - returns the pubg stats
func Handler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, stats)
}
