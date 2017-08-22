package handlers

import (
	"io"
	"os"
	"strings"

	"net/http"

	"github.com/mgerb/chi_auth_server/response"
	"github.com/mgerb/go-discord-bot/server/config"
)

// FileUpload
func FileUpload(w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("password")

	if string(password) != config.Config.UploadPassword {
		response.ERR(w, http.StatusInternalServerError, []byte("Invalid password."))
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Error reading file."))
		return
	}

	defer file.Close()

	src, err := header.Open()
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Error opening file."))
		return
	}

	defer src.Close()

	// create uploads folder if it does not exist
	if _, err := os.Stat(config.Config.SoundsPath); os.IsNotExist(err) {
		os.Mkdir(config.Config.SoundsPath, os.ModePerm)
	}

	// convert file name to lower case
	header.Filename = strings.ToLower(header.Filename)

	// check if file already exists
	if _, err := os.Stat(config.Config.SoundsPath + header.Filename); err == nil {
		response.ERR(w, http.StatusInternalServerError, []byte("File already exists."))
		return
	}

	dst, err := os.Create(config.Config.SoundsPath + header.Filename)
	if err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Error creating file."))
		return
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Error writing file."))
		return
	}

	// repopulate sound list
	err = PopulateSoundList()

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Error populating sound list."))
		return
	}

	response.JSON(w, []byte("Success"))
}
