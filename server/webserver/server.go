package webserver

import (
	"strings"

	"github.com/mgerb/go-discord-bot/server/webserver/response"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver/routes"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./dist/static")
	router.Static("/public/sounds", config.Config.SoundsPath)
	router.Static("/public/youtube", config.Config.YoutubePath)
	router.Static("/public/clips", config.Config.ClipsPath)

	api := router.Group("/api")

	// add api routes
	routes.AddSoundListRoutes(api)
	routes.AddOauthRoutes(api)
	routes.AddLoggerRoutes(api)
	routes.AddDownloaderRoutes(api)
	routes.AddConfigRoutes(api)
	routes.AddSoundRoutes(api)
	routes.AddVideoArchiveRoutes(api)
	routes.AddUserEventLogRoutes(api)
	routes.AddUserRoutes(api)

	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.String(), "/api/") {
			response.BadRequest(c, "404 Not Found")
		} else {
			c.File("./dist/static/index.html")
		}
	})

	return router
}

// Start -
func Start() {
	router := getRouter()
	router.Run(config.Config.ServerAddr)
}
