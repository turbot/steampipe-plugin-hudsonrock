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
			{Name: "logo", Type: proto.ColumnType_STRING, Description: "Logo URL."},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "Raw data from the API response."},
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

	var result struct {
		Total         int         `json:"total"`
		TotalStealers int         `json:"totalStealers"`
		Employees     int         `json:"employees"`
		Users         int         `json:"users"`
		ThirdParties  int         `json:"third_parties"`
		Logo          string      `json:"logo"`
		Data          interface{} `json:"data"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	d.StreamListItem(ctx, map[string]interface{}{
		"domain":         domain,
		"total":          result.Total,
		"total_stealers": result.TotalStealers,
		"employees":      result.Employees,
		"users":          result.Users,
		"third_parties":  result.ThirdParties,
		"logo":           result.Logo,
		"data":           result.Data,
	})
	return nil, nil
}
