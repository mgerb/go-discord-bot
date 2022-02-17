package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"
	"github.com/mgerb/go-discord-bot/server/bot"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	"github.com/mgerb/go-discord-bot/server/webserver/response"
)

// AddVideoArchiveRoutes -
func AddVideoArchiveRoutes(group *gin.RouterGroup) {
	group.GET("/video-archive", listVideoArchiveHandler)

	authGroup := group.Group("", middleware.AuthorizedJWT())
	authGroup.POST("/video-archive", middleware.AuthPermissions(middleware.PermMod), postVideoArchiveHandler)
	authGroup.DELETE("/video-archive/:id", middleware.AuthPermissions(middleware.PermAdmin), deleteVideoArchiveHandler)
}

func listVideoArchiveHandler(c *gin.Context) {
	archives, err := model.VideoArchiveList(db.GetConn())

	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, archives)
}

func deleteVideoArchiveHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.BadRequest(c, "Invalid ID")
		return
	}

	err := model.VideoArchiveDelete(db.GetConn(), id)

	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, "deleted")
}

func postVideoArchiveHandler(c *gin.Context) {
	params := struct {
		URL string `json:"url"`
	}{}

	c.BindJSON(&params)

	if params.URL == "" {
		response.BadRequest(c, "URL Required")
		return
	}

	cl, _ := c.Get("claims")
	claims, ok := cl.(*middleware.CustomClaims)

	if !ok {
		response.InternalError(c, errors.New("Claims error"))
		return
	}

	client := youtube.Client{}

	info, err := client.GetVideo(params.URL)

	if err != nil {
		response.InternalError(c, err)
		return
	}

	// if title and author are blank
	if info.Title == "" && info.Author == "" {
		response.BadRequest(c, "Invalid URL")
		return
	}

	videoArchive := model.VideoArchive{
		Author:        info.Author,
		DatePublished: info.PublishDate,
		Description:   info.Description,
		Duration:      int(info.Duration.Seconds()),
		Title:         info.Title,
		URL:           params.URL,
		YoutubeID:     info.ID,
		UploadedBy:    claims.Username,
	}

	err = model.VideoArchiveSave(db.GetConn(), &videoArchive)

	if err != nil {
		response.InternalError(c, err)
		return
	}

	hostURL := "[Click here to see the full archive!](https://" + c.Request.Host + "/video-archive)"
	youtubeURL := "https://youtu.be/" + videoArchive.YoutubeID
	bot.SendEmbeddedNotification(videoArchive.Title, "**"+videoArchive.UploadedBy+"** archived a new video:\n"+youtubeURL+"\n\n"+hostURL)

	response.Success(c, "saved")
}
