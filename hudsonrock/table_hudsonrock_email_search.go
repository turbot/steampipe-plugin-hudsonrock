package hudsonrock

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-hudsonrock/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableHudsonrockEmailSearch(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_email_search",
		Description: "Search for compromised credentials and infostealer data by email using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "email", Require: plugin.Required},
			},
			Hydrate: listHudsonrockEmailSearch,
		},
		Columns: []*plugin.Column{
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email searched.", Transform: transform.FromQual("email")},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the email."},
			{Name: "date_compromised", Type: proto.ColumnType_TIMESTAMP, Description: "Date the credentials were compromised."},
			{Name: "stealer_family", Type: proto.ColumnType_STRING, Description: "Infostealer malware family."},
			{Name: "computer_name", Type: proto.ColumnType_STRING, Description: "Name of the infected computer."},
			{Name: "operating_system", Type: proto.ColumnType_STRING, Description: "Operating system of the infected computer."},
			{Name: "malware_path", Type: proto.ColumnType_STRING, Description: "Path to the malware on the system."},
			{Name: "antiviruses", Type: proto.ColumnType_JSON, Description: "Antivirus software detected."},
			{Name: "ip", Type: proto.ColumnType_STRING, Description: "IP address of the infected machine."},
			{Name: "top_passwords", Type: proto.ColumnType_JSON, Description: "Top passwords found."},
			{Name: "top_logins", Type: proto.ColumnType_JSON, Description: "Top logins found."},
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
	result, err := client.EmailSearch(email)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}

	plugin.Logger(ctx).Error("listHudsonrockDomainSearch", "connection_error", result)

	d.StreamListItem(ctx, result)
	return nil, nil

}
