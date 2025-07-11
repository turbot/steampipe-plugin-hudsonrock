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

func tableHudsonrockIpSearch(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_ip_search",
		Description: "Search for info-stealer data by IP address using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "ip", Require: plugin.Required},
			},
			Hydrate: listHudsonrockIpSearch,
		},
		Columns: []*plugin.Column{
			{Name: "ip", Type: proto.ColumnType_STRING, Description: "IP address searched."},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the IP address."},
			{Name: "stealers", Type: proto.ColumnType_JSON, Description: "Stealer data from the API response."},
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

	url := "https://cavalier.hudsonrock.com/api/json/v2/osint-tools/search-by-ip?ip=" + ip
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
		Message                string      `json:"message"`
		Stealers               interface{} `json:"stealers"`
		TotalCorporateServices int         `json:"total_corporate_services"`
		TotalUserServices      int         `json:"total_user_services"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	d.StreamListItem(ctx, map[string]interface{}{
		"ip":                       ip,
		"message":                  result.Message,
		"stealers":                 result.Stealers,
		"total_corporate_services": result.TotalCorporateServices,
		"total_user_services":      result.TotalUserServices,
	})
	return nil, nil
}
