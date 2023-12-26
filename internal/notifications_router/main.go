package notifications_router

import (
	"context"
	"net/http"

	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/tokend/connectors/signed"
)

const registerServicePath = "/internal/services"

type (
	Connector interface {
		Register(ctx context.Context) error
	}
	connector struct {
		cfg       *config
		connector *jsonapi.Connector
	}
)

func NewConnector(cfg *config) Connector {
	client := signed.NewClient(http.DefaultClient, cfg.RegisterEndpoint)
	return &connector{
		cfg:       cfg,
		connector: jsonapi.NewConnector(client),
	}
}

func (c *connector) Register(ctx context.Context) error {
	request := c.composeRegisterServiceRequest()
	endpoint := c.cfg.RegisterEndpoint.JoinPath(registerServicePath)
	return c.connector.PostJSON(endpoint, request, ctx, nil)
}

func (c *connector) composeRegisterServiceRequest() *registerServiceRequest {
	return &registerServiceRequest{
		Data: registerServiceData{
			Type: registerServiceType,
			Attributes: registerServiceAttributes{
				Endpoint: c.cfg.Upstream.String(),
				Channels: []string{c.cfg.Channel},
			},
		},
	}
}
