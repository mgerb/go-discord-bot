package handlers

import (
	"io/ioutil"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	log "github.com/sirupsen/logrus"
)

type sound struct {
	Prefix    string `json:"prefix"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

// SoundList -
func SoundList(c *gin.Context) {

	soundList, err := readSoundsDir(config.Config.SoundsPath)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, soundList)
}

// ClipList -
func ClipList(c *gin.Context) {

	clipList, err := readSoundsDir(config.Config.ClipsPath)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, clipList)
}

func readSoundsDir(dir string) ([]sound, error) {

	soundList := []sound{}

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return soundList, err
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

	return soundList, nil
}
