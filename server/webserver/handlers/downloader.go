package handlers

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	"github.com/mgerb/chi_auth_server/response"
)

// Downloader -
func Downloader(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	fileType := r.FormValue("fileType")

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
		log.Println(err)
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
		return
	}

	// TODO add video id to tile to not get collisions

	// ------------------------------------------------

	// remove all special characters from title
	cleanTitle := cleanseTitle(titleOut.String())
	log.Println(cleanTitle)

	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "-o", "./youtube/"+cleanTitle+".%(ext)s", url)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		log.Println(out.String())
		log.Println(err)
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
		return
	}

	response.JSON(w, map[string]interface{}{"fileName": cleanTitle + "." + fileType})
}

func cleanseTitle(title string) string {

	// Make a Regex to say we only want
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(title, "")
}
