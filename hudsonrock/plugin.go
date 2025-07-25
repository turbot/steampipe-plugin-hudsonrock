package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-hudsonrock"

// Plugin returns the Hudson Rock plugin definition.
func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo().NullIfEmptySlice(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"hudsonrock_search_by_domain":   tableHudsonrockSearchByDomain(ctx),
			"hudsonrock_search_by_email":    tableHudsonrockSearchByEmail(ctx),
			"hudsonrock_search_by_ip":       tableHudsonrockSearchByIp(ctx),
			"hudsonrock_url_by_domain":      tableHudsonrockUrlByDomain(ctx),
			"hudsonrock_search_by_username": tableHudsonrockSearchByUsername(ctx),
		},
	}
}
