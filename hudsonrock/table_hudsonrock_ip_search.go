package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-hudsonrock/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tableHudsonrockIpSearch(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_ip_search",
		Description: "Search for info-stealer data by IP address using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "ip", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
			},
			Hydrate: listHudsonrockIpSearch,
		},
		Columns: []*plugin.Column{
			{Name: "ip", Type: proto.ColumnType_STRING, Description: "IP address searched.", Transform: transform.FromQual("ip")},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the IP address."},
			{Name: "stealers", Type: proto.ColumnType_JSON, Description: "List of stealer compromise details for the IP address."},
			{Name: "total_corporate_services", Type: proto.ColumnType_INT, Description: "Total corporate services found."},
			{Name: "total_user_services", Type: proto.ColumnType_INT, Description: "Total user services found."},
		},
	}
}

func listHudsonrockIpSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ip := d.EqualsQuals["ip"].GetStringValue()
	if ip == "" {
		return nil, nil
	}

	client := api.NewClient()
	result, err := client.IpSearch(ctx, ip)
	if err != nil {
		plugin.Logger(ctx).Error("listHudsonrockIpSearch", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Error("listHudsonrockIpSearch", "connection_error", result)

	d.StreamListItem(ctx, result)
	return nil, nil
}
