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
			{Name: "date_compromised", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the computer was compromised."},
			{Name: "computer_name", Type: proto.ColumnType_STRING, Description: "Name of the infected computer."},
			{Name: "operating_system", Type: proto.ColumnType_STRING, Description: "Operating system of the infected computer."},
			{Name: "malware_path", Type: proto.ColumnType_STRING, Description: "File path of the detected malware on the infected computer."},
			{Name: "antiviruses", Type: proto.ColumnType_JSON, Description: "List of antivirus products found on the infected computer."},
			{Name: "top_passwords", Type: proto.ColumnType_JSON, Description: "Top passwords found on the infected computer."},
			{Name: "top_logins", Type: proto.ColumnType_JSON, Description: "Top logins found on the infected computer."},
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
		Message  string `json:"message"`
		Stealers []struct {
			TotalCorporateServices int      `json:"total_corporate_services"`
			TotalUserServices      int      `json:"total_user_services"`
			DateCompromised        string   `json:"date_compromised"`
			Ip                     string   `json:"ip"`
			ComputerName           string   `json:"computer_name"`
			OperatingSystem        string   `json:"operating_system"`
			MalwarePath            string   `json:"malware_path"`
			Antiviruses            []string `json:"antiviruses"`
			TopPasswords           []string `json:"top_passwords"`
			TopLogins              []string `json:"top_logins"`
		} `json:"stealers"`
		TotalCorporateServices int `json:"total_corporate_services"`
		TotalUserServices      int `json:"total_user_services"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// If there are multiple stealers, stream each as a row; otherwise, stream one row with details or nulls.
	if len(result.Stealers) > 0 {
		for _, s := range result.Stealers {
			d.StreamListItem(ctx, map[string]interface{}{
				"ip":                       ip,
				"message":                  result.Message,
				"date_compromised":         s.DateCompromised,
				"computer_name":            s.ComputerName,
				"operating_system":         s.OperatingSystem,
				"malware_path":             s.MalwarePath,
				"antiviruses":              s.Antiviruses,
				"top_passwords":            s.TopPasswords,
				"top_logins":               s.TopLogins,
				"total_corporate_services": s.TotalCorporateServices,
				"total_user_services":      s.TotalUserServices,
			})
		}
	} else {
		d.StreamListItem(ctx, map[string]interface{}{
			"ip":                       ip,
			"message":                  result.Message,
			"date_compromised":         nil,
			"computer_name":            nil,
			"operating_system":         nil,
			"malware_path":             nil,
			"antiviruses":              nil,
			"top_passwords":            nil,
			"top_logins":               nil,
			"total_corporate_services": result.TotalCorporateServices,
			"total_user_services":      result.TotalUserServices,
		})
	}
	return nil, nil
}
