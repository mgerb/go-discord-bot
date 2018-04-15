package webserver

import (
	"github.com/gobuffalo/packr"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver/handlers"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	box := packr.NewBox("../../dist/static")

	router.StaticFS("/static", box)
	router.Static("/public/sounds", config.Config.SoundsPath)
	router.Static("/public/youtube", "./youtube")
	router.Static("/public/clips", config.Config.ClipsPath)

	router.NoRoute(func(c *gin.Context) {
		c.Data(200, "text/html", box.Bytes("index.html"))
	})

	api := router.Group("/api")
	api.GET("/ytdownloader", handlers.Downloader)
	api.GET("/soundlist", handlers.SoundList)
	api.GET("/cliplist", handlers.ClipList)
	api.POST("/oauth", handlers.Oauth)
	api.GET("/logger/messages", handlers.GetMessages)
	api.GET("/logger/linkedmessages", handlers.GetLinkedMessages)

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
