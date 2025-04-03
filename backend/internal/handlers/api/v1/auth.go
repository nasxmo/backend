package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/client-go"
)

type AuthHandler struct {
	kratos *client.APIClient
	hydra  *client.APIClient
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	kratosCfg := client.NewConfiguration()
	kratosCfg.Servers = []client.ServerConfiguration{
		{URL: cfg.KratosURL},
	}

	hydraCfg := client.NewConfiguration()
	hydraCfg.Servers = []client.ServerConfiguration{
		{URL: cfg.HydraURL},
	}

	return &AuthHandler{
		kratos: client.NewAPIClient(kratosCfg),
		hydra:  client.NewAPIClient(hydraCfg),
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		// Create new login flow
		flow, _, err := h.kratos.FrontendApi.CreateBrowserLoginFlow(c.Request.Context()).Execute()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, flow)
		return
	}

	// Complete login flow
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, _, err := h.kratos.FrontendApi.UpdateLoginFlow(c.Request.Context()).
		Flow(flowId).
		UpdateLoginFlowBody(body).
		Execute()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *AuthHandler) Authenticate(c *gin.Context) {
	session, _, err := h.kratos.FrontendApi.ToSession(c.Request.Context()).Execute()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Store session in context
	c.Set("session", session)
	c.Next()
}