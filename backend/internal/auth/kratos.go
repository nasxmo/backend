package auth

import (
	"context"
	"fmt"

	"github.com/ory/client-go"
)

type KratosClient struct {
	public *client.APIClient
	admin *client.APIClient
}

func NewKratosClient(publicURL, adminURL string) *KratosClient {
	publicConfig := client.NewConfiguration()
	publicConfig.Servers = client.ServerConfigurations{{URL: publicURL}}

	adminConfig := client.NewConfiguration()
	adminConfig.Servers = client.ServerConfigurations{{URL: adminURL}}

	return &KratosClient{
		public: client.NewAPIClient(publicConfig),
		admin: client.NewAPIClient(adminConfig),
	}
}

func (k *KratosClient) GetSession(ctx context.Context, cookie string) (*client.Session, error) {
	session, _, err := k.public.FrontendApi.ToSession(ctx).
		Cookie(cookie).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}
	return session, nil
}

func (k *KratosClient) CreateIdentity(ctx context.Context, identity client.Identity) (*client.Identity, error) {
	created, _, err := k.admin.IdentityApi.CreateIdentity(ctx).
		CreateIdentityBody(client.CreateIdentityBody(identity)).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create identity: %w", err)
	}
	return created, nil
}