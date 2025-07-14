package api

import (
	"fmt"
	"net/url"
)

// UsernameSearchResponse represents the response for a username compromise search.
type UsernameSearchResponse struct {
	Message                string            `json:"message"`
	Stealers               []UsernameStealer `json:"stealers"`
	TotalCorporateServices int               `json:"total_corporate_services"`
	TotalUserServices      int               `json:"total_user_services"`
}

// UsernameStealer contains details about each compromise event for the username.
type UsernameStealer struct {
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
}

func (c *Client) UsernameSearch(username string) (UsernameSearchResponse, error) {
	// Build full URL using BaseURL constant
	endpoint, err := url.Parse(BaseURL)
	if err != nil {
		return UsernameSearchResponse{}, err
	}
	endpoint.Path = "/api/json/v2/osint-tools/search-by-username"

	// Add query parameters
	query := endpoint.Query()
	query.Set("username", username)
	endpoint.RawQuery = query.Encode()

	var result UsernameSearchResponse

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