package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/webserver/discord"
	log "github.com/sirupsen/logrus"
)

const cashGuildID = "101198129352691712"

type oauthReq struct {
	Code string `json:"code"`
}

// Oauth -
func Oauth(c *gin.Context) {

	var json oauthReq

	err := c.ShouldBindJSON(&json)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	oauth, err := discord.Oauth(json.Code)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	user, err := discord.GetUserInfo(oauth.AccessToken)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	// TODO: generate jwt for user
	fmt.Println(user)

	c.JSON(200, oauth)
}
