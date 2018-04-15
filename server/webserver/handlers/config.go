package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
)

func GetClientID(c *gin.Context) {
	c.JSON(200, map[string]string{"id": config.Config.ClientID})
}
