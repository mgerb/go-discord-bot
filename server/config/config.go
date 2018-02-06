package config

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

// Variables used for command line parameters
var (
	Config configFile
	Flags  configFlags
)

type configFile struct {
	Token          string `json:"Token"`
	BotPrefix      string `json:"BotPrefix"` //prefix to use for bot commands
	SoundsPath     string `json:"SoundsPath"`
	ClipsPath      string `json:"ClipsPath"`
	UploadPassword string `json:"UploadPassword"`
	ServerAddr     string `json:"ServerAddr`
	Pubg           struct {
		Enabled bool     `json:"enabled"`
		APIKey  string   `json:"apiKey"`
		Players []string `json:"players"`
	} `json:"pubg"`
}

type configFlags struct {
	Prod bool
}

// Init -
func Init() {
	parseConfig()
	parseFlags()
}

func parseConfig() {

	log.Debug("Reading config file...")

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		log.Fatal("File error: %v\n", e)
	}

	log.Debug("%s\n", string(file))

	err := json.Unmarshal(file, &Config)

	if err != nil {
		log.Error(err)
	}
}

func parseFlags() {

	Flags.Prod = false

	prod := flag.Bool("p", false, "Run in production")

	flag.Parse()

	Flags.Prod = *prod

	if Flags.Prod {
		log.Warn("Running in production mode")
	}

}
