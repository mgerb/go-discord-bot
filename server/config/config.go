package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Config configType
)

type configType struct {
	Token        string   `json:"token"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURI  string   `json:"redirect_uri"`
	BotPrefix    string   `json:"bot_prefix"` //prefix to use for bot commands
	AdminEmails  []string `json:"admin_emails"`
	ModEmails    []string `json:"mod_emails"`
	ServerAddr   string   `json:"server_addr"`
	JWTSecret    string   `json:"jwt_secret"`
	Logger       bool     `json:"logger"`

	// hard coded folder paths
	SoundsPath  string
	ClipsPath   string
	YoutubePath string
}

// Init -
func Init() {
	parseConfig()

	Config.SoundsPath = "./sounds"
	Config.ClipsPath = "./clips"
	Config.YoutubePath = "./youtube"
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
		log.Fatal(err)
	}
}
