package api

import (
	"context"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"resty.dev/v3"
)

// URLSearchResponse represents the response for URL compromise search.
type URLSearchResponse struct {
	Message string       `json:"message"`
	Data    URLDataGroup `json:"data"`
}

// URLDataGroup holds lists of URLs for employees and clients.
type URLDataGroup struct {
	EmployeesURLs []URLInfo `json:"employees_urls"`
	ClientsURLs   []URLInfo `json:"clients_urls"`
}

func (c *Client) UrlByDomain(ctx context.Context, domain string) (URLSearchResponse, error) {
	// Build full URL using BaseURL constant
	endpoint, err := url.Parse(BaseURL)
	if err != nil {
		return URLSearchResponse{}, err
	}
	endpoint.Path = "/api/json/v2/osint-tools/urls-by-domain"

	// Add query parameters
	query := endpoint.Query()
	query.Set("domain", domain)
	endpoint.RawQuery = query.Encode()

	var result URLSearchResponse

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
		plugin.Logger(ctx).Error("Domain search failed", "domain", domain, "error", err)
		return result, err
	}

	plugin.Logger(ctx).Debug("Domain search completed successfully",
		"domain", domain,
		"status", resp.StatusCode(),
		"max_retries", c.MaxRetries)

	return result, nil
}
