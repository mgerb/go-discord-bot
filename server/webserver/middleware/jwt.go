package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver/model"
	log "github.com/sirupsen/logrus"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

// permission levels
const (
	PermAdmin = 3
	PermMod   = 2
	PermUser  = 1
)

// CustomClaims -
type CustomClaims struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Email         string `json:"email"`
	Permissions   int    `json:"permissions"`
	jwt.StandardClaims
}

// GetJWT - get json web token
func GetJWT(user model.User) (string, error) {

	claims := CustomClaims{
		user.ID,
		user.Username,
		user.Discriminator,
		user.Email,
		*user.Permissions,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 12, 0).Unix(), // twelve months
			Issuer:    "Go Discord Bot",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.JWTSecret))
}

// AuthPermissions - secure end points based on auth levels
func AuthPermissions(p int) gin.HandlerFunc {
	return func(c *gin.Context) {
		cl, _ := c.Get("claims")

		if claims, ok := cl.(*CustomClaims); ok {
			if p <= claims.Permissions {
				c.Next()
				return
			}
		}

		unauthorizedResponse(c, nil)
	}
}

// AuthorizedJWT - jwt middleware
func AuthorizedJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// grab token from authorization header: Bearer token
		tokenString := strings.Split(c.GetHeader("Authorization"), " ")

		if len(tokenString) != 2 {
			unauthorizedResponse(c, nil)
			return
		}

		// parse and verify token
		token, err := jwt.ParseWithClaims(tokenString[1], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWTSecret), nil
		})

		if err != nil {
			unauthorizedResponse(c, err)
			return
		}

		// get claims and set on gin context
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			unauthorizedResponse(c, err)
			return
		}

		c.Next()
	}
}

func unauthorizedResponse(c *gin.Context, err error) {
	if err != nil {
		log.Error(err)
	}
	c.JSON(401, "unauthorized")
	c.Abort()
}
