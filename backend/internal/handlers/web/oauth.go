package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/client-go"
)

func OAuthCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	// Exchange code for tokens
	token, _, err := hydraClient.OAuth2Api.GetOAuth2Token(c.Request.Context()).
		Code(code).
		Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user info
	userInfo, _, err := hydraClient.OidcApi.GetUserInfo(c.Request.Context()).
		Authorization("Bearer " + token.AccessToken).
		Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"userInfo": userInfo,
	})
}