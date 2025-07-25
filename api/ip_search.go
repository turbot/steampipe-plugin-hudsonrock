package api

import (
	"context"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"resty.dev/v3"
)

// IPSearchResponse represents the response for an IP compromise search.
type IPSearchResponse struct {
	Message                string      `json:"message"`
	Stealers               []IPStealer `json:"stealers"`
	TotalCorporateServices int         `json:"total_corporate_services"`
	TotalUserServices      int         `json:"total_user_services"`
}

// IPStealer contains details about the infection associated with the IP.
type IPStealer struct {
	TotalCorporateServices int      `json:"total_corporate_services"`
	TotalUserServices      int      `json:"total_user_services"`
	DateCompromised        string   `json:"date_compromised"`
	IP                     string   `json:"ip"`
	ComputerName           string   `json:"computer_name"`
	OperatingSystem        string   `json:"operating_system"`
	MalwarePath            string   `json:"malware_path"`
	Antiviruses            []string `json:"antiviruses"`
	TopPasswords           []string `json:"top_passwords"`
	TopLogins              []string `json:"top_logins"`
}

func (c *Client) IpSearch(ctx context.Context, ip string) (IPSearchResponse, error) {
	// Build full URL using BaseURL constant
	endpoint, err := url.Parse(BaseURL)
	if err != nil {
		return IPSearchResponse{}, err
	}
	endpoint.Path = "/api/json/v2/osint-tools/search-by-ip"

	// Add query parameters
	query := endpoint.Query()
	query.Set("ip", ip)
	endpoint.RawQuery = query.Encode()

	var result IPSearchResponse

	// Create the request function for retry logic
	requestFunc := func() (*resty.Response, error) {
		return c.Resty.R().
			SetHeader("Accept", "application/json").
			SetResult(&result).
			Get(endpoint.String())
	}

	// Execute with client's default retry settings
	resp, err := c.executeWithRetryDefault(requestFunc)
	if err != nil {
		plugin.Logger(ctx).Error("IP search failed", "ip", ip, "error", err)
		return result, err
	}

	plugin.Logger(ctx).Debug("IP search completed successfully",
		"ip", ip,
		"status", resp.StatusCode(),
		"max_retries", c.MaxRetries)

	return result, nil
}
