package bothandlers

/***** DEPRECATED *****/

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/tidwall/gjson"
)

const (
	gifPrefix string = "!gif "
	userAgent string = "go-discord-bot"
	giphyURL  string = "http://api.giphy.com/v1/gifs/search?&api_key=dc6zaTOxFJmzC&limit=10&q="
)

// GifHandler - handler for giphy api
func GifHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// check if valid command
	if strings.HasPrefix(m.Content, gifPrefix) {

		searchText := strings.TrimPrefix(m.Content, gifPrefix)

		gifLink := getGiphy(searchText)

		if gifLink == "null" {
			gifLink = "No gif found."
		}

		s.ChannelMessageSend(m.ChannelID, gifLink)

	}
}

// send http request to reddit
func getGiphy(searchTerm string) string {

	rand.Seed(time.Now().UnixNano())

	client := &http.Client{}

	req, err := http.NewRequest("GET", giphyURL+searchTerm, nil)

	response, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}

	data := gjson.Get(string(body), "data").Array()

	if len(data) < 1 {
		return "null"
	}

	return data[rand.Intn(len(data))].Get("images.fixed_height.url").String()
}
