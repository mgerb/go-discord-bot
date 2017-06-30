package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/valyala/fasthttp"
)

var soundList []sound

type sound struct {
	Prefix    string `json:"prefix"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

func SoundList(ctx *fasthttp.RequestCtx) {

	if len(soundList) < 1 {
		err := PopulateSoundList()
		if err != nil {
			ctx.Error(err.Error(), 400)
			return
		}
	}

	response, err := json.Marshal(soundList)

	if err != nil {
		ctx.Error("Error marshaling json", 400)
		return
	}

	ctx.SetContentType("application/json")
	ctx.Write(response)
}

func PopulateSoundList() error {
	fmt.Println("Populating sound list.")

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
