package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
)

// AddLoggerRoutes -
func AddLoggerRoutes(group *gin.RouterGroup) {
	group.GET("/logger/messages", getMessagesHandler)
	group.GET("/logger/linkedmessages", getLinkedMessagesHandler)
}

func getMessagesHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		page = 0
	}

	messages, err := model.MessageGet(db.GetConn(), page)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, messages)
}

func getLinkedMessagesHandler(c *gin.Context) {
	posts, err := model.MessageGetLinked(db.GetConn())

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, posts)
}
