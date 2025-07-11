package api

import (
	"fmt"
	"net/url"
)

// Define response struct
type EmailSearchResponse struct {
	Message                string           `json:"message"`
	Stealers               []StealerDetails `json:"stealers"`
	TotalCorporateServices int              `json:"total_corporate_services"`
	TotalUserServices      int              `json:"total_user_services"`
}

// StealerDetails contains information about each stealer compromise.
type StealerDetails struct {
	TotalCorporateServices int      `json:"total_corporate_services"`
	TotalUserServices      int      `json:"total_user_services"`
	DateCompromised        string   `json:"date_compromised"`
	ComputerName           string   `json:"computer_name"`
	OperatingSystem        string   `json:"operating_system"`
	MalwarePath            string   `json:"malware_path"`
	Antiviruses            []string `json:"antiviruses"`
	IP                     string   `json:"ip"`
	TopPasswords           []string `json:"top_passwords"`
	TopLogins              []string `json:"top_logins"`
}

func (c *Client) EmailSearch(email string) (EmailSearchResponse, error) {
	// Build full URL using BaseURL constant
	endpoint, err := url.Parse(BaseURL)
	if err != nil {
		return EmailSearchResponse{}, err
	}
	endpoint.Path = "/api/json/v2/osint-tools/search-by-email"

	// Add query parameters
	query := endpoint.Query()
	query.Set("email", email)
	endpoint.RawQuery = query.Encode()

	var result EmailSearchResponse

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
