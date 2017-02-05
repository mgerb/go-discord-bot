package handlers

import (
	"../../config"
	"github.com/valyala/fasthttp"
	"io"
	"os"
)

func FileUpload(ctx *fasthttp.RequestCtx) {
	password := ctx.FormValue("password")

	if string(password) != config.Config.UploadPassword {
		ctx.Error("Invalid password.", 400)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Error("Error reading file.", 400)
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.Error("Error opening file.", 400)
		return
	}

	defer src.Close()

	// create uploads folder if it does not exist
	if _, err := os.Stat(config.Config.SoundsPath); os.IsNotExist(err) {
		os.Mkdir(config.Config.SoundsPath, os.ModePerm)
	}

	// check if file already exists
	if _, err := os.Stat(config.Config.SoundsPath + file.Filename); err == nil {
		ctx.Error("File already exists.", 400)
		return
	}

	dst, err := os.Create(config.Config.SoundsPath + file.Filename)
	if err != nil {
		ctx.Error("Error creating file.", 400)
		return
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		ctx.Error("Error writing file.", 400)
		return
	}

	// repopulate sound list
	err = PopulateSoundList()

	if err != nil {
		ctx.Error("File uploaded, but error populating sound list.", 400)
		return
	}

	ctx.Success("application/json", []byte("Success!"))
}
