package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-hudsonrock/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tableHudsonrockEmailSearch(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_email_search",
		Description: "Search for compromised credentials and infostealer data by email using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "email", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
			},
			Hydrate: listHudsonrockEmailSearch,
		},
		Columns: []*plugin.Column{
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email searched.", Transform: transform.FromQual("email")},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the email."},
			{Name: "stealers", Type: proto.ColumnType_JSON, Description: "List of stealer compromise details for the email."},
			{Name: "total_corporate_services", Type: proto.ColumnType_INT, Description: "Total corporate services found."},
			{Name: "total_user_services", Type: proto.ColumnType_INT, Description: "Total user services found."},
		},
	}
}

func listHudsonrockEmailSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	email := d.EqualsQuals["email"].GetStringValue()
	if email == "" {
		return nil, nil
	}

	client := api.NewClient()
	result, err := client.EmailSearch(ctx, email)
	if err != nil {
		plugin.Logger(ctx).Error("listHudsonrockEmailSearch", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, result)
	return nil, nil
}
