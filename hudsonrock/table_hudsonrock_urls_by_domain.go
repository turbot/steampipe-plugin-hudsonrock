package hudsonrock

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-hudsonrock/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tableHudsonrockUrlsByDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_urls_by_domain",
		Description: "List URLs identified by infostealer infections for a given domain using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "domain", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
			},
			Hydrate: listHudsonrockUrlsByDomain,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain searched.", Transform: transform.FromQual("domain")},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the domain."},
			{Name: "employees_urls", Type: proto.ColumnType_JSON, Description: "List of employee URLs from the API response.", Transform: transform.FromField("Data.EmployeesURLs")},
			{Name: "clients_urls", Type: proto.ColumnType_JSON, Description: "List of client URLs from the API response.", Transform: transform.FromField("Data.ClientsURLs")},
		},
	}
}

func listHudsonrockUrlsByDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	domain := d.EqualsQuals["domain"].GetStringValue()
	if domain == "" {
		return nil, nil
	}

	client := api.NewClient()
	result, err := client.URLsByDomain(domain)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}

	plugin.Logger(ctx).Error("listHudsonrockUrlsByDomain", "connection_error", result)

	d.StreamListItem(ctx, result)
	return nil, nil
}
