package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver/handlers"
	"github.com/mgerb/go-discord-bot/server/webserver/pubg"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./dist/static")
	router.Static("/public/sounds", config.Config.SoundsPath)
	router.Static("/public/youtube", "./youtube")
	router.Static("/public/clips", config.Config.ClipsPath)

	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	api := router.Group("/api")
	api.GET("/stats/pubg", pubg.Handler)
	api.GET("/ytdownloader", handlers.Downloader)
	api.GET("/soundlist", handlers.SoundList)
	api.GET("/cliplist", handlers.ClipList)
	api.POST("/upload", handlers.FileUpload)

	return router
}

// Start -
func Start() {

	// start gathering pubg data from the api
	if config.Config.Pubg.Enabled {
		pubg.Start(config.Config.Pubg.APIKey, config.Config.Pubg.Players)
	}

	router := getRouter()
	router.Run(config.Config.ServerAddr)
}
