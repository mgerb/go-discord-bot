package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/bothandlers"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	"github.com/mgerb/go-discord-bot/server/webserver/response"
	log "github.com/sirupsen/logrus"
)

// AddSoundRoutes -
func AddSoundRoutes(group *gin.RouterGroup) {
	group.GET("/sound", listSoundHandler)
	group.POST("/sound", middleware.AuthorizedJWT(), postSoundHandler)
	group.POST("/sound/play", middleware.AuthorizedJWT(), middleware.AuthPermissions(middleware.PermMod), postSoundPlayHandler)
}

func listSoundHandler(c *gin.Context) {
	archives, err := model.SoundList(db.GetConn())

	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, archives)
}

func postSoundPlayHandler(c *gin.Context) {
	connections := bothandlers.ActiveConnections

	params := struct {
		Name string `json:"name"`
	}{}
	c.BindJSON(&params)

	// loop through all connections and play audio
	// currently only used with one server
	// will need selector on UI if used for multiple servers
	if len(connections) == 1 && params.Name != "" {
		for _, con := range connections {
			if params.Name == "random" {
				con.PlayRandomAudio(nil)
			} else {
				con.PlayAudio(params.Name, nil)
			}
		}
	}

	response.Success(c, "test")
}

func postSoundHandler(c *gin.Context) {

	oc, _ := c.Get("claims")
	claims, _ := oc.(*middleware.CustomClaims)

	// TODO: verify user for upload

	file, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "Error reading file.")
		return
	}

	// create uploads folder if it does not exist
	if _, err := os.Stat(config.Config.SoundsPath); os.IsNotExist(err) {
		os.Mkdir(config.Config.SoundsPath, os.ModePerm)
	}

	// convert file name to lower case and trim spaces
	file.Filename = strings.Replace(strings.ToLower(file.Filename), " ", "", -1)

	// check if file already exists
	if _, err := os.Stat(config.Config.SoundsPath + "/" + file.Filename); err == nil {
		c.JSON(http.StatusInternalServerError, "File already exists.")
		return
	}

	err = c.SaveUploadedFile(file, config.Config.SoundsPath+"/"+file.Filename)
	log.Info(claims.Username, "uploaded", config.Config.SoundsPath+"/"+file.Filename)

	// save who uploaded the clip into the database
	uploadSaveDB(claims.ID, file.Filename)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "Error creating file.")
		return
	}

	c.JSON(200, "Success")
}

// save new sound to database
func uploadSaveDB(userID, filename string) {
	splitFilename := strings.Split(filename, ".")
	extension := splitFilename[len(splitFilename)-1]
	name := strings.Join(splitFilename[:len(splitFilename)-1], ".")

	model.SoundCreate(db.GetConn(), &model.Sound{
		UserID:    userID,
		Name:      name,
		Extension: extension,
	})
}
