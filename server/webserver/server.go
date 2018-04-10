package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver/handlers"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./dist/static")
	router.Static("/public/sounds", config.Config.SoundsPath)
	router.Static("/public/youtube", "./youtube")
	router.Static("/public/clips", config.Config.ClipsPath)

	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/static/index.html")
	})

	api := router.Group("/api")
	api.GET("/ytdownloader", handlers.Downloader)
	api.GET("/soundlist", handlers.SoundList)
	api.GET("/cliplist", handlers.ClipList)
	api.POST("/oauth", handlers.Oauth)

	authorizedAPI := router.Group("/api")
	authorizedAPI.Use(middleware.AuthorizedJWT())
	authorizedAPI.POST("/upload", handlers.FileUpload)

	return router
}

// Start -
func Start() {

	router := getRouter()
	router.Run(config.Config.ServerAddr)
}
