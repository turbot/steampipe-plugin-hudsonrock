package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const pluginName = "steampipe-plugin-hudsonrock"

// Plugin returns the Hudson Rock plugin definition.
func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"hudsonrock_username_search": tableHudsonrockUsernameSearch(ctx),
			"hudsonrock_email_search":    tableHudsonrockEmailSearch(ctx),
			"hudsonrock_domain_search":   tableHudsonrockDomainSearch(ctx),
			"hudsonrock_ip_search":       tableHudsonrockIpSearch(ctx),
			"hudsonrock_urls_by_domain":  tableHudsonrockUrlsByDomain(ctx),
		},
	}
}
