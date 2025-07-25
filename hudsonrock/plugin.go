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
			"hudsonrock_domain_search":   tableHudsonrockDomainSearch(ctx),
			"hudsonrock_email_search":    tableHudsonrockEmailSearch(ctx),
			"hudsonrock_ip_search":       tableHudsonrockIpSearch(ctx),
			"hudsonrock_urls_by_domain":  tableHudsonrockUrlsByDomain(ctx),
			"hudsonrock_username_search": tableHudsonrockUsernameSearch(ctx),
		},
	}
}

