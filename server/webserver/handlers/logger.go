package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/logger"
)

// GetLogs - get all logs
func GetLogs(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		page = 0
	}

	messages := []logger.Message{}
	err = db.Conn.Offset(page*100).Limit(100).Order("timestamp desc", true).Preload("User").Find(&messages).Error

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, messages)
}
