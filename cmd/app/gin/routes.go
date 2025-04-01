package gin

import (
	"github.com/ShudderStorm/go-github-tracker/cmd/app/internal/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	LoginEndpoint    string = "/login"
	CallbackEndpoint string = "/callback"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		url, err := oauth.Client.GetAuthorizationUrl()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create auth URL"})
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func CallbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
