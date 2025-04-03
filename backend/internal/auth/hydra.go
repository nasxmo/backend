package auth

import (
	"context"
	"fmt"

	"github.com/ory/client-go"
)

type HydraClient struct {
	admin *client.APIClient
	public *client.APIClient
}

func NewHydraClient(publicURL, adminURL string) *HydraClient {
	publicConfig := client.NewConfiguration()
	publicConfig.Servers = client.ServerConfigurations{{URL: publicURL}}

	adminConfig := client.NewConfiguration()
	adminConfig.Servers = client.ServerConfigurations{{URL: adminURL}}

	return &HydraClient{
		admin: client.NewAPIClient(adminConfig),
		public: client.NewAPIClient(publicConfig),
	}
}

func (h *HydraClient) CreateOAuth2Client(ctx context.Context, clientConfig client.OAuth2Client) (*client.OAuth2Client, error) {
	client, _, err := h.admin.OAuth2Api.CreateOAuth2Client(ctx).OAuth2Client(clientConfig).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create OAuth2 client: %w", err)
	}
	return client, nil
}

func (h *HydraClient) VerifyToken(ctx context.Context, token string) (*client.IntrospectedOAuth2Token, error) {
	introspection, _, err := h.admin.OAuth2Api.IntrospectOAuth2Token(ctx).
		Token(token).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to introspect token: %w", err)
	}
	return introspection, nil
}