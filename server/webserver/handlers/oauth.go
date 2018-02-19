package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/webserver/discord"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
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

	// get users oauth code
	oauth, err := discord.Oauth(json.Code)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	// verify and grab user information
	user, err := discord.GetUserInfo(oauth.AccessToken)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	// generate json web token
	token, err := middleware.GetJWT(user)

	if err != nil {
		log.Error(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, token)
}
