package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-hudsonrock/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tableHudsonrockUsernameSearch(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_username_search",
		Description: "Search for compromised credentials and infostealer data by username using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "username", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
			},
			Hydrate: listHudsonrockUsernameSearch,
		},
		Columns: []*plugin.Column{
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Username searched.", Transform: transform.FromQual("username")},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the username."},
			{Name: "stealers", Type: proto.ColumnType_JSON, Description: "List of stealer compromise details for the username."},
			{Name: "total_corporate_services", Type: proto.ColumnType_INT, Description: "Total corporate services found."},
			{Name: "total_user_services", Type: proto.ColumnType_INT, Description: "Total user services found."},
		},
	}
}

func listHudsonrockUsernameSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	username := quals["username"].GetStringValue()
	if username == "" {
		return nil, nil
	}

	client := api.NewClient()
	result, err := client.UsernameSearch(ctx, username)
	if err != nil {
		plugin.Logger(ctx).Error("listHudsonrockUsernameSearch", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, result)
	return nil, nil

}
