package handlers

import (
	"io/ioutil"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	log "github.com/sirupsen/logrus"
)

var soundList []sound

type sound struct {
	Prefix    string `json:"prefix"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

// SoundList -
func SoundList(c *gin.Context) {

	if len(soundList) < 1 {
		err := PopulateSoundList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(200, soundList)
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
func ClipList(c *gin.Context) {

	clipList := []sound{}

	files, err := ioutil.ReadDir(config.Config.ClipsPath)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
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

	c.JSON(200, clipList)
}
