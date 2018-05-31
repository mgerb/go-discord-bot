package handlers

import (
	"os"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	log "github.com/sirupsen/logrus"
)

// FileUpload -
func FileUpload(c *gin.Context) {

	// originalClaims, _ := c.Get("claims")
	// claims, _ := originalClaims.(*middleware.CustomClaims)
	// TODO: verify user for upload

	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "Error reading file.")
		return
	}

	// create uploads folder if it does not exist
	if _, err := os.Stat(config.Config.SoundsPath); os.IsNotExist(err) {
		os.Mkdir(config.Config.SoundsPath, os.ModePerm)
	}

	// convert file name to lower case and trim spaces
	file.Filename = strings.ToLower(file.Filename)
	file.Filename = strings.Replace(file.Filename, " ", "", -1)

	// check if file already exists
	if _, err := os.Stat(config.Config.SoundsPath + "/" + file.Filename); err == nil {
		c.JSON(http.StatusInternalServerError, "File already exists.")
		return
	}

	err = c.SaveUploadedFile(file, config.Config.SoundsPath+"/"+file.Filename)
	log.Debug("Saving file", config.Config.SoundsPath+"/"+file.Filename)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "Error creating file.")
		return
	}

	c.JSON(200, "Success")
}
