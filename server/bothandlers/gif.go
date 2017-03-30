package bothandlers

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tidwall/gjson"
)

const (
	gifPrefix string = "!gif "
	userAgent string = "go-discord-bot"
	giphyURL  string = "http://api.giphy.com/v1/stickers/random?api_key=dc6zaTOxFJmzC&rating=r&tag="
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

	data := gjson.Get(string(body), "data.url")

	return data.String()
}
