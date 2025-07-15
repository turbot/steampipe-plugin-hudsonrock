package hudsonrock

import (
	"context"

	"github.com/turbot/steampipe-plugin-hudsonrock/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tableHudsonrockDomainSearch(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_domain_search",
		Description: "Search for domain-related cybercrime and infostealer intelligence using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "domain", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
			},
			Hydrate: listHudsonrockDomainSearch,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain searched.", Transform: transform.FromQual("domain")},
			{Name: "total", Type: proto.ColumnType_INT, Description: "Total records found."},
			{Name: "total_stealers", Type: proto.ColumnType_INT, Description: "Total stealers found."},
			{Name: "employees", Type: proto.ColumnType_INT, Description: "Number of employees."},
			{Name: "users", Type: proto.ColumnType_INT, Description: "Number of users."},
			{Name: "third_parties", Type: proto.ColumnType_INT, Description: "Number of third parties."},
			{Name: "logo", Type: proto.ColumnType_STRING, Description: "Logo URL for the domain."},
			{Name: "total_urls", Type: proto.ColumnType_INT, Description: "Total number of unique URLs associated with the domain."},
			{Name: "stats", Type: proto.ColumnType_JSON, Description: "Statistical breakdown of employees and users, including top URLs and counts."},
			{Name: "is_shopify", Type: proto.ColumnType_BOOL, Description: "Indicates if the domain is a Shopify store."},
			{Name: "last_employee_compromised", Type: proto.ColumnType_STRING, Description: "Timestamp of the last employee compromise for the domain."},
			{Name: "last_user_compromised", Type: proto.ColumnType_STRING, Description: "Timestamp of the last user compromise for the domain."},
			{Name: "antiviruses", Type: proto.ColumnType_JSON, Description: "Antivirus statistics and list of antivirus products found in the dataset."},
			{Name: "applications", Type: proto.ColumnType_JSON, Description: "List of detected application keywords related to the domain."},
			{Name: "employee_passwords", Type: proto.ColumnType_JSON, Description: "Password strength statistics for employees of the domain."},
			{Name: "user_passwords", Type: proto.ColumnType_JSON, Description: "Password strength statistics for users of the domain."},
			{Name: "third_party_domains", Type: proto.ColumnType_JSON, Description: "List of third-party domains associated with the main domain, with occurrence counts."},
			{Name: "stealer_families", Type: proto.ColumnType_JSON, Description: "Breakdown of stealer malware families found in the dataset for the domain."},
			{Name: "employees_urls", Type: proto.ColumnType_JSON, Description: "List of URLs associated with employees for the given domain.", Transform: transform.FromField("Data.EmployeesURLs")},
			{Name: "clients_urls", Type: proto.ColumnType_JSON, Description: "List of URLs associated with clients for the given domain.", Transform: transform.FromField("Data.ClientsURLs")},
			{Name: "all_urls", Type: proto.ColumnType_JSON, Description: "List of all URLs (employees and clients) associated with the given domain.", Transform: transform.FromField("Data.AllURLs")},
		},
	}
}


func listHudsonrockDomainSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	domain := d.EqualsQuals["domain"].GetStringValue()
	if domain == "" {
		return nil, nil
	}

	client := api.NewClient()
	result, err := client.DomainSearch(ctx, domain)
	if err != nil {
		plugin.Logger(ctx).Error("listHudsonrockDomainSearch", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, result)
	return nil, nil
}
