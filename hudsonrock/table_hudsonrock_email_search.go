package hudsonrock

import (
	"context"

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
			{Name: "stealer_total_corporate_services", Type: proto.ColumnType_INT, Description: "Stealer total corporate services found.", Transform: transform.FromField("Stealer.TotalCorporateServices")},
			{Name: "stealer_total_user_services", Type: proto.ColumnType_INT, Description: "Stealer total user services found.", Transform: transform.FromField("Stealer.TotalUserServices")},
			{Name: "date_compromised", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the computer was compromised.", Transform: transform.FromField("Stealer.DateCompromised")},
			{Name: "computer_name", Type: proto.ColumnType_STRING, Description: "Name of the infected computer.", Transform: transform.FromField("Stealer.ComputerName")},
			{Name: "operating_system", Type: proto.ColumnType_STRING, Description: "Operating system of the infected computer.", Transform: transform.FromField("Stealer.OperatingSystem")},
			{Name: "malware_path", Type: proto.ColumnType_STRING, Description: "File path of the detected malware on the infected computer.", Transform: transform.FromField("Stealer.MalwarePath")},
			{Name: "antiviruses", Type: proto.ColumnType_JSON, Description: "List of antivirus products found on the infected computer.", Transform: transform.FromField("Stealer.Antiviruses")},
			{Name: "top_passwords", Type: proto.ColumnType_JSON, Description: "Top passwords found on the infected computer.", Transform: transform.FromField("Stealer.TopPasswords")},
			{Name: "top_logins", Type: proto.ColumnType_JSON, Description: "Top logins found on the infected computer.", Transform: transform.FromField("Stealer.TopLogins")},
			{Name: "total_corporate_services", Type: proto.ColumnType_INT, Description: "Total corporate services found."},
			{Name: "total_user_services", Type: proto.ColumnType_INT, Description: "Total user services found."},
		},
	}
}

type EmailDetails struct {
	Message                string
	TotalCorporateServices int
	TotalUserServices      int
	Stealer                api.EmailStealer
}

func listHudsonrockEmailSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	email := d.EqualsQuals["email"].GetStringValue()
	if email == "" {
		return nil, nil
	}

	client := api.NewClient()
	output, err := client.EmailSearch(ctx, email)
	if err != nil {
		plugin.Logger(ctx).Error("hudsonrock_email_search.listHudsonrockEmailSearch", "api_error", err)
		return nil, err
	}

	for _, result := range output.Stealers {
		d.StreamListItem(ctx, &EmailDetails{output.Message, output.TotalCorporateServices, output.TotalUserServices, result})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
