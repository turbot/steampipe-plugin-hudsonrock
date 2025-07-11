package hudsonrock

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableHudsonrockDomainSearch(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_domain_search",
		Description: "Search for domain-related cybercrime and infostealer intelligence using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "domain", Require: plugin.Required},
			},
			Hydrate: listHudsonrockDomainSearch,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain searched."},
			{Name: "total", Type: proto.ColumnType_INT, Description: "Total records found."},
			{Name: "total_stealers", Type: proto.ColumnType_INT, Description: "Total stealers found."},
			{Name: "employees", Type: proto.ColumnType_INT, Description: "Number of employees."},
			{Name: "users", Type: proto.ColumnType_INT, Description: "Number of users."},
			{Name: "third_parties", Type: proto.ColumnType_INT, Description: "Number of third parties."},
			{Name: "logo", Type: proto.ColumnType_STRING, Description: "Logo URL for the domain."},
			{Name: "total_urls", Type: proto.ColumnType_INT, Description: "Total number of unique URLs associated with the domain."},
			{Name: "stats", Type: proto.ColumnType_JSON, Description: "Statistical breakdown of employees and users, including top URLs and counts."},
			{Name: "is_shopify", Type: proto.ColumnType_BOOL, Description: "Indicates if the domain is a Shopify store."},
			{Name: "last_employee_compromised", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the last employee compromise for the domain."},
			{Name: "last_user_compromised", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the last user compromise for the domain."},
			{Name: "antiviruses", Type: proto.ColumnType_JSON, Description: "Antivirus statistics and list of antivirus products found in the dataset."},
			{Name: "applications", Type: proto.ColumnType_JSON, Description: "List of detected application keywords related to the domain."},
			{Name: "employee_passwords", Type: proto.ColumnType_JSON, Description: "Password strength statistics for employees of the domain."},
			{Name: "user_passwords", Type: proto.ColumnType_JSON, Description: "Password strength statistics for users of the domain."},
			{Name: "third_party_domains", Type: proto.ColumnType_JSON, Description: "List of third-party domains associated with the main domain, with occurrence counts."},
			{Name: "stealer_families", Type: proto.ColumnType_JSON, Description: "Breakdown of stealer malware families found in the dataset for the domain."},
		},
	}
}

func listHudsonrockDomainSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	domain := d.EqualsQuals["domain"].GetStringValue()
	if domain == "" {
		return nil, nil
	}

	url := "https://cavalier.hudsonrock.com/api/json/v2/osint-tools/search-by-domain?domain=" + domain
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s - %s", resp.Status, string(b))
	}

	type DomainSearchResponse struct {
		Total                   int                      `json:"total"`
		TotalStealers           int                      `json:"totalStealers"`
		Employees               int                      `json:"employees"`
		Users                   int                      `json:"users"`
		ThirdParties            int                      `json:"third_parties"`
		Logo                    string                   `json:"logo"`
		TotalUrls               int                      `json:"totalUrls"`
		Stats                   map[string]interface{}   `json:"stats"`
		IsShopify               bool                     `json:"is_shopify"`
		LastEmployeeCompromised string                   `json:"last_employee_compromised"`
		LastUserCompromised     string                   `json:"last_user_compromised"`
		Antiviruses             map[string]interface{}   `json:"antiviruses"`
		Applications            []map[string]interface{} `json:"applications"`
		EmployeePasswords       map[string]interface{}   `json:"employeePasswords"`
		UserPasswords           map[string]interface{}   `json:"userPasswords"`
		ThirdPartyDomains       []map[string]interface{} `json:"thirdPartyDomains"`
		StealerFamilies         map[string]interface{}   `json:"stealerFamilies"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	var result DomainSearchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	d.StreamListItem(ctx, map[string]interface{}{
		"domain":                    domain,
		"total":                     result.Total,
		"total_stealers":            result.TotalStealers,
		"employees":                 result.Employees,
		"users":                     result.Users,
		"third_parties":             result.ThirdParties,
		"logo":                      result.Logo,
		"total_urls":                result.TotalUrls,
		"stats":                     result.Stats,
		"is_shopify":                result.IsShopify,
		"last_employee_compromised": result.LastEmployeeCompromised,
		"last_user_compromised":     result.LastUserCompromised,
		"antiviruses":               result.Antiviruses,
		"applications":              result.Applications,
		"employee_passwords":        result.EmployeePasswords,
		"user_passwords":            result.UserPasswords,
		"third_party_domains":       result.ThirdPartyDomains,
		"stealer_families":          result.StealerFamilies,
	})
	return nil, nil
}
