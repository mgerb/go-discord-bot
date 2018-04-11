package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/logger"
)

// GetMessages - get all messages
func GetMessages(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		page = 0
	}

	messages, err := logger.GetMessages(page)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, messages)
}

// GetLinkedMessages -
func GetLinkedMessages(c *gin.Context) {
	posts, err := logger.GetLinkedMessages()

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, posts)
}
