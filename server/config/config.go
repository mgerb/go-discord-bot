package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Variables used for command line parameters
var Config configStruct

type configStruct struct {
	Token          string `json:"Token"`
	BotPrefix      string `json:"BotPrefix"` //prefix to use for bot commands
	SoundsPath     string `json:"SoundsPath"`
	UploadPassword string `json:"UploadPassword"`
	ServerAddr     string `json:"ServerAddr`
}

func Init() {

	log.Println("Reading config file...")

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		log.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	log.Printf("%s\n", string(file))

	err := json.Unmarshal(file, &Config)

	if err != nil {
		log.Println(err)
	}

}
