package api

import (
	"fmt"
	"net/url"
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

func (c *Client) URLsByDomain(domain string) (URLSearchResponse, error) {
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
