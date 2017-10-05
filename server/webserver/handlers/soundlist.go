package handlers

import (
	"io/ioutil"
	"log"
	"strings"

	"net/http"

	"github.com/mgerb/chi_auth_server/response"
	"github.com/mgerb/go-discord-bot/server/config"
)

var soundList []sound

type sound struct {
	Prefix    string `json:"prefix"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

// SoundList -
func SoundList(w http.ResponseWriter, r *http.Request) {

	if len(soundList) < 1 {
		err := PopulateSoundList()
		if err != nil {
			response.ERR(w, http.StatusInternalServerError, []byte(response.DefaultInternalError))
			return
		}
	}

	response.JSON(w, soundList)
}

// PopulateSoundList -
func PopulateSoundList() error {

	soundList = []sound{}

	files, err := ioutil.ReadDir(config.Config.SoundsPath)

	if err != nil {
		return err
	}

	for _, f := range files {
		fileName := strings.Split(f.Name(), ".")[0]
		extension := strings.Split(f.Name(), ".")[1]

		listItem := sound{
			Name:      fileName,
			Extension: extension,
			Prefix:    config.Config.BotPrefix,
		}

		soundList = append(soundList, listItem)
	}

	return nil
}

// ClipList -
func ClipList(w http.ResponseWriter, r *http.Request) {

	clipList := []sound{}

	files, err := ioutil.ReadDir(config.Config.ClipsPath)

	if err != nil {
		log.Println(err)
		response.ERR(w, 500, response.DefaultInternalError)
		return
	}

	for _, f := range files {
		fileName := strings.Split(f.Name(), ".")[0]
		extension := strings.Split(f.Name(), ".")[1]

		listItem := sound{
			Name:      fileName,
			Extension: extension,
		}

		clipList = append(clipList, listItem)
	}

	response.JSON(w, clipList)
}
