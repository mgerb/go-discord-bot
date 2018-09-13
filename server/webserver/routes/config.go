package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
)

// AddConfigRoutes -
func AddConfigRoutes(group *gin.RouterGroup) {
	group.GET("/config/client_id", getConfigHandler)
}

func getConfigHandler(c *gin.Context) {
	c.JSON(200, map[string]string{"id": config.Config.ClientID})
}
