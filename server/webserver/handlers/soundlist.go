package handlers

import (
	"../../config"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"strings"
)

var soundList = make([]string, 0)

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

	soundList = make([]string, 0)

	var fileName string
	files, err := ioutil.ReadDir(config.Config.SoundsPath)

	if err != nil {
		return err
	}

	for _, f := range files {
		fileName = config.Config.BotPrefix + strings.Split(f.Name(), ".")[0]
		soundList = append(soundList, fileName)
	}

	return nil
}
