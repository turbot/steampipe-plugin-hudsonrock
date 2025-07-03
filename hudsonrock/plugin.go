package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const pluginName = "steampipe-plugin-hudsonrock"

// Plugin returns the Hudson Rock plugin definition.
func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: pluginName,
		TableMap: map[string]*plugin.Table{
			"hudsonrock_username_search": tableHudsonrockUsernameSearch(ctx),
		},
	}
}
