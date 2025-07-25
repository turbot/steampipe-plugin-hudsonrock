package hudsonrock

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type HudsonRockConfig struct {
	MaxRetries *int   `hcl:"max_retries,optional"`
	MinDelay   *int64 `hcl:"min_delay,optional"`
}

func ConfigInstance() interface{} {
	return &HudsonRockConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) HudsonRockConfig {
	if connection == nil || connection.Config == nil {
		return HudsonRockConfig{}
	}
	config, _ := connection.Config.(HudsonRockConfig)

	return config
}
