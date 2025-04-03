package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/client-go"
)

func GetProfile(c *gin.Context) {
	session, _ := c.Get("session")
	c.JSON(http.StatusOK, session.(*client.Session))
}

func ChangePassword(c *gin.Context) {
	session := c.MustGet("session").(*client.Session)
	
	var body struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update password through Kratos
	flow, _, err := kratosClient.FrontendApi.CreateBrowserSettingsFlow(c.Request.Context()).
		Execute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = kratosClient.FrontendApi.UpdateSettingsFlow(c.Request.Context()).
		Flow(flow.Id).
		UpdateSettingsFlowBody(map[string]interface{}{
			"method":   "password",
			"password": body.NewPassword,
		}).
		Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "password updated"})
}