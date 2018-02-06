package handlers

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Downloader -
func Downloader(c *gin.Context) {
	url := c.Query("url")
	fileType := c.Query("fileType")

	// create youtube folder if it does not exist
	if _, err := os.Stat("youtube"); os.IsNotExist(err) {
		os.Mkdir("youtube", os.ModePerm)
	}

	// get the video title
	titleCmd := exec.Command("youtube-dl", "--get-title", url)
	var titleOut bytes.Buffer
	titleCmd.Stdout = &titleOut

	err := titleCmd.Run()

	if err != nil {
		log.Error(err)
		c.JSON(400, err)
		return
	}

	// TODO add video id to tile to not get collisions

	// ------------------------------------------------

	// remove all special characters from title
	cleanTitle := cleanseTitle(titleOut.String())
	log.Debug(cleanTitle)

	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "-o", "./youtube/"+cleanTitle+".%(ext)s", url)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		log.Error(out.String())
		log.Error(err)
		c.JSON(400, err)
		return
	}

	c.JSON(200, map[string]interface{}{"fileName": cleanTitle + "." + fileType})
}

func cleanseTitle(title string) string {

	// Make a Regex to say we only want
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Error(err)
	}

	return reg.ReplaceAllString(title, "")
}
