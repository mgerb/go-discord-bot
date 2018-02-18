package discord

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/mgerb/go-discord-bot/server/config"
)

const discordAPI = "https://discordapp.com/api/v6"

// OauthResp -
type OauthResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// Oauth -
func Oauth(code string) (OauthResp, error) {

	form := url.Values{}

	form.Set("client_id", config.Config.ClientID)
	form.Set("client_secret", config.Config.ClientSecret)
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", config.Config.RedirectURI)

	req, err := http.NewRequest("POST", discordAPI+"/oauth2/token", strings.NewReader(form.Encode()))

	if err != nil {
		return OauthResp{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return OauthResp{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return OauthResp{}, err
	}

	var oauth OauthResp

	err = json.Unmarshal(data, &oauth)

	if err != nil {
		return OauthResp{}, err
	}

	return oauth, nil
}
