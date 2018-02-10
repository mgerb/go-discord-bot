package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	log "github.com/sirupsen/logrus"
)

const discordApi = "https://discordapp.com/api/v6/oauth2/token"

type oauthReq struct {
	Code string `json:"code"`
}

type oauthResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func Oauth(c *gin.Context) {

	var json oauthReq

	err := c.ShouldBindJSON(&json)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	oauth, err := postReq(json.Code)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, oauth)
}

func postReq(code string) (oauthResp, error) {

	form := url.Values{}

	form.Set("client_id", config.Config.ClientId)
	form.Set("client_secret", config.Config.ClientSecret)
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", config.Config.RedirectUri)

	req, err := http.NewRequest("POST", discordApi, strings.NewReader(form.Encode()))

	if err != nil {
		return oauthResp{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}
	resp, err := c.Do(req)

	if err != nil {
		return oauthResp{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return oauthResp{}, err
	}

	var oauth oauthResp

	err = json.Unmarshal(data, &oauth)

	if err != nil {
		return oauthResp{}, err
	}

	return oauth, nil
}
