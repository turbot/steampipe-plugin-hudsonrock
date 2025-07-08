package hudsonrock

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email searched."},
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
			{Name: "data", Type: proto.ColumnType_JSON, Description: "Raw data from the API response."},
		},
	}
}

func listHudsonrockEmailSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	email := d.EqualsQuals["email"].GetStringValue()
	if email == "" {
		return nil, nil
	}

	url := "https://cavalier.hudsonrock.com/api/json/v2/osint-tools/search-by-email?email=" + email
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s - %s", resp.Status, string(b))
	}

	var result struct {
		Message  string `json:"message"`
		Stealers []struct {
			TotalCorporateServices int      `json:"total_corporate_services"`
			TotalUserServices      int      `json:"total_user_services"`
			DateCompromised        string   `json:"date_compromised"`
			StealerFamily          string   `json:"stealer_family"`
			ComputerName           string   `json:"computer_name"`
			OperatingSystem        string   `json:"operating_system"`
			MalwarePath            string   `json:"malware_path"`
			Antiviruses            []string `json:"antiviruses"`
			IP                     string   `json:"ip"`
			TopPasswords           []string `json:"top_passwords"`
			TopLogins              []string `json:"top_logins"`
		} `json:"stealers"`
		TotalCorporateServices int `json:"total_corporate_services"`
		TotalUserServices      int `json:"total_user_services"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	for _, s := range result.Stealers {
		d.StreamListItem(ctx, map[string]interface{}{
			"email":                    email,
			"message":                  result.Message,
			"date_compromised":         s.DateCompromised,
			"stealer_family":           s.StealerFamily,
			"computer_name":            s.ComputerName,
			"operating_system":         s.OperatingSystem,
			"malware_path":             s.MalwarePath,
			"antiviruses":              s.Antiviruses,
			"ip":                       s.IP,
			"top_passwords":            s.TopPasswords,
			"top_logins":               s.TopLogins,
			"total_corporate_services": s.TotalCorporateServices,
			"total_user_services":      s.TotalUserServices,
			"data":                     s,
		})
	}
	return nil, nil
}
