package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// BadRequest -
func BadRequest(c *gin.Context, message string) {
	dataMap := map[string]string{
		"data": message,
	}
	c.JSON(http.StatusBadRequest, dataMap)
}

// TODO: don't return the error on production

// InternalError -
func InternalError(c *gin.Context, err error) {
	log.Error(err)
	c.JSON(http.StatusInternalServerError, err)
}

// Success -
func Success(c *gin.Context, data interface{}) {
	dataMap := map[string]interface{}{
		"data": data,
	}
	c.JSON(http.StatusOK, dataMap)
}
