package hudsonrock

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
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
	if connection == nil || connection.GetConfig() == nil {
		return HudsonRockConfig{}
	}
	config, _ := connection.GetConfig().(HudsonRockConfig)

	return config
}
