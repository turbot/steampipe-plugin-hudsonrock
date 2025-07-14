package api

import (
	"fmt"
	"net/url"
)

// IPSearchResponse represents the response for an IP compromise search.
type IPSearchResponse struct {
	Message                string           `json:"message"`
	Stealers               []IPStealer      `json:"stealers"`
	TotalCorporateServices int              `json:"total_corporate_services"`
	TotalUserServices      int              `json:"total_user_services"`
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

func (c *Client) IpSearch(ip string) (IPSearchResponse, error) {
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

	// Make the request with proper error handling
	resp, err := c.Resty.R().
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get(endpoint.String())

	if err != nil {
		return result, err
	}

	// Handle HTTP errors
	if resp.IsError() {
		return result, fmt.Errorf("HTTP error: %s", resp.Status())
	}

	return result, nil
}

