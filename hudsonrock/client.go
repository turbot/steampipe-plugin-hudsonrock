package hudsonrock

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-hudsonrock/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func NewClient(ctx context.Context, d *plugin.QueryData) *api.Client {
	config := GetConfig(d.Connection)

	client := api.NewClient()

	if config.MaxRetries != nil {
		client.WithMaxRetries(*config.MaxRetries)
	}
	if config.MinDelay != nil {
		client.WithMinDelay(time.Duration(*config.MinDelay) * time.Second)
	}

	return client
}
