package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHudsonrockUrlByDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_url_by_domain",
		Description: "List URLs identified by infostealer infections for a given domain using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "domain", Require: plugin.Required},
			},
			Hydrate: listHudsonrockUrlByDomain,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain searched.", Transform: transform.FromQual("domain")},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the domain."},
			{Name: "employees_urls", Type: proto.ColumnType_JSON, Description: "List of URLs associated with employees for the given domain.", Transform: transform.FromField("Data.EmployeesURLs")},
			{Name: "clients_urls", Type: proto.ColumnType_JSON, Description: "List of URLs associated with clients for the given domain.", Transform: transform.FromField("Data.ClientsURLs")},
			{Name: "all_urls", Type: proto.ColumnType_JSON, Description: "List of all URLs (employees and clients) associated with the given domain.", Transform: transform.FromField("Data.AllURLs")},
		},
	}
}

func listHudsonrockUrlByDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	domain := d.EqualsQuals["domain"].GetStringValue()
	if domain == "" {
		return nil, nil
	}

	client := NewClient(ctx, d)
	result, err := client.UrlByDomain(ctx, domain)
	if err != nil {
		plugin.Logger(ctx).Error("hudsonrock_url_by_domain.listHudsonrockUrlByDomain", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, result)
	return nil, nil
}
