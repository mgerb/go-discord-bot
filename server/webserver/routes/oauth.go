package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/discord"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	log "github.com/sirupsen/logrus"
)

type oauthReq struct {
	Code string `json:"code"`
}

// AddOauthRoutes -
func AddOauthRoutes(group *gin.RouterGroup) {
	group.POST("/oauth", oauthHandler)
}

func oauthHandler(c *gin.Context) {

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

	// save/update user in database
	err = model.UserSave(db.GetConn(), &user)

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
