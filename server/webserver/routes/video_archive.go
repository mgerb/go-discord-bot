package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	"github.com/mgerb/go-discord-bot/server/webserver/response"
	"github.com/rylio/ytdl"
)

// AddVideoArchiveRoutes -
func AddVideoArchiveRoutes(group *gin.RouterGroup) {
	group.GET("/video-archives", listVideoArchivesHandler)

	authGroup := group.Group("", middleware.AuthorizedJWT())
	authGroup.POST("/video-archives", middleware.AuthPermissions(middleware.PermMod), postVideoArchivesHandler)
	authGroup.DELETE("/video-archives/:id", middleware.AuthPermissions(middleware.PermAdmin), deleteVideoArchivesHandler)
}

func listVideoArchivesHandler(c *gin.Context) {
	archives, err := model.VideoArchiveList(db.GetConn())

	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, archives)
}

func deleteVideoArchivesHandler(c *gin.Context) {
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

func postVideoArchivesHandler(c *gin.Context) {
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

	info, err := ytdl.GetVideoInfo(params.URL)

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
		DatePublished: info.DatePublished,
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

	response.Success(c, "saved")
}
