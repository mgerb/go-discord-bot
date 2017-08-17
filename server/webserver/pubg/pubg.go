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

	api, err := pubgClient.New(apiKey)

	if err != nil {
		log.Fatal(err)
	}

	// fetch new stats every 30 seconds
	for {

		for _, player := range players {
			newStats, err := api.GetPlayer(player)

			if err != nil {
				log.Println(err)
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
func Handler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, stats)
}
