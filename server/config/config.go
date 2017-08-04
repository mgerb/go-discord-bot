package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
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
	TLS  bool
}

// Init -
func Init() {

	parseConfig()
	parseFlags()

}

func parseConfig() {

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

func parseFlags() {

	Flags.Prod = false
	Flags.TLS = false

	prod := flag.Bool("p", false, "Run in production")
	tls := flag.Bool("tls", false, "Use TLS")

	flag.Parse()

	Flags.Prod = *prod
	Flags.TLS = *tls

	if *prod {
		log.Println("Running in production mode")
	}

}
