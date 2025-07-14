package api

import (
	"context"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"resty.dev/v3"
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

func (c *Client) EmailSearch(ctx context.Context, email string) (EmailSearchResponse, error) {
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
		plugin.Logger(ctx).Error("Email search failed", "email", email, "error", err)
		return result, err
	}

	plugin.Logger(ctx).Debug("Email search completed successfully",
		"email", email,
		"status", resp.StatusCode(),
		"max_retries", c.MaxRetries)

	return result, nil
}
