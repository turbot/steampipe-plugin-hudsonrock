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

func tableHudsonrockUrlsByDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hudsonrock_urls_by_domain",
		Description: "List URLs identified by infostealer infections for a given domain using Hudson Rock's API.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "domain", Require: plugin.Required},
			},
			Hydrate: listHudsonrockUrlsByDomain,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain searched."},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "API message about the domain."},
			{Name: "employees_urls", Type: proto.ColumnType_JSON, Description: "List of employee URLs from the API response."},
			{Name: "clients_urls", Type: proto.ColumnType_JSON, Description: "List of client URLs from the API response."},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "Raw data from the API response."},
		},
	}
}

func listHudsonrockUrlsByDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	domain := d.EqualsQuals["domain"].GetStringValue()
	if domain == "" {
		return nil, nil
	}

	url := "https://cavalier.hudsonrock.com/api/json/v2/osint-tools/urls-by-domain?domain=" + domain
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
		Message string `json:"message"`
		Data    struct {
			EmployeesUrls []interface{} `json:"employees_urls"`
			ClientsUrls   []interface{} `json:"clients_urls"`
		} `json:"data"`
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
		"message":        result.Message,
		"employees_urls": result.Data.EmployeesUrls,
		"clients_urls":   result.Data.ClientsUrls,
		"data":           result.Data,
	})
	return nil, nil
}
