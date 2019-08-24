package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/db"
	"github.com/mgerb/go-discord-bot/server/webserver/middleware"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	"github.com/mgerb/go-discord-bot/server/webserver/response"
)

// AddUserEventLogRoutes -
func AddUserEventLogRoutes(group *gin.RouterGroup) {
	group.GET("/user-event-log", middleware.AuthorizedJWT(), middleware.AuthPermissions(middleware.PermAdmin), listEventLogHandler)
}

func listEventLogHandler(c *gin.Context) {

	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		page = 0
	}

	userEventLogs, err := model.UserEventLogGet(db.GetConn(), page)

	if err != nil {
		response.InternalError(c, err)
	} else {
		response.Success(c, userEventLogs)
	}
}
