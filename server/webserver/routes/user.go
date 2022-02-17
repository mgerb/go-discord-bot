package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	"github.com/mgerb/go-discord-bot/server/webserver/response"
)

func AddUserRoutes(group *gin.RouterGroup) {
	group.GET("/user", middleware.AuthorizedJWT(), middleware.AuthPermissions(middleware.PermAdmin), userHandler)
	group.PUT("/user", middleware.AuthorizedJWT(), middleware.AuthPermissions(middleware.PermAdmin), userUpdateHandler)
}

func userHandler(c *gin.Context) {
	users, err := model.UserGetAll(db.GetConn())

	if err != nil {
		response.InternalError(c, err)
	} else {
		response.Success(c, users)
	}
}

func userUpdateHandler(c *gin.Context) {
	params := struct {
		Users []model.User `json:"users"`
	}{}
	c.BindJSON(&params)

	for _, user := range params.Users {
		_, err := model.UserUpdate(db.GetConn(), &user)

		if err != nil {
			response.InternalError(c, err)
			return
		}
	}

	response.Success(c, params.Users)
}
