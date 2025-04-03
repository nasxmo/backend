package auth

import (
	"context"
	"fmt"

	"github.com/ory/oathkeeper-client-go"
)

type OathkeeperClient struct {
	client *oathkeeper.APIClient
}

func NewOathkeeperClient(url string) *OathkeeperClient {
	config := oathkeeper.NewConfiguration()
	config.Servers = oathkeeper.ServerConfigurations{{URL: url}}

	return &OathkeeperClient{
		client: oathkeeper.NewAPIClient(config),
	}
}

func (o *OathkeeperClient) CreateRule(ctx context.Context, rule oathkeeper.Rule) (*oathkeeper.Rule, error) {
	created, _, err := o.client.ApiApi.CreateRule(ctx).Rule(rule).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create rule: %w", err)
	}
	return created, nil
}

func (o *OathkeeperClient) VerifyAccess(ctx context.Context, request oathkeeper.DecisionsRequest) (*oathkeeper.DecisionsForbidden, error) {
	decision, _, err := o.client.DecisionApi.GetDecisions(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("access denied: %w", err)
	}
	return decision, nil
}