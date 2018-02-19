package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/go-discord-bot/server/config"
	"github.com/mgerb/go-discord-bot/server/webserver/discord"
	log "github.com/sirupsen/logrus"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

// CustomClaims -
type CustomClaims struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Email         string `json:"email"`
	Permissions   string `json:"permissions"`
	jwt.StandardClaims
}

// GetJWT - get json web token
func GetJWT(user discord.User) (string, error) {

	permissions := "user"

	// check if email is in config admin list
	for _, email := range config.Config.AdminEmails {
		if user.Email == email {
			permissions = "admin"
			break
		}
	}

	claims := CustomClaims{
		user.ID,
		user.Username,
		user.Discriminator,
		user.Email,
		permissions,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(), // one month
			Issuer:    "Go Discord Bot",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.JWTKey))
}

// AuthorizedJWT - jwt middleware
func AuthorizedJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// grab token from authorization header: Bearer token
		tokenString := strings.Split(c.GetHeader("Authorization"), " ")

		if len(tokenString) != 2 {
			c.JSON(401, "Unauthorized")
			c.Abort()
			return
		}

		// parse and verify token
		token, err := jwt.ParseWithClaims(tokenString[1], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWTKey), nil
		})

		if err != nil {
			log.Error(err)
			c.JSON(401, "Unauthorized")
			c.Abort()
			return
		}

		// get claims and set on gin context
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			log.Error(err)
			c.JSON(401, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
