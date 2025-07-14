package api

import (
	"context"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"resty.dev/v3"
)

// Define response struct
type DomainSearchResponse struct {
	Total                   int                `json:"total"`
	TotalStealers           int64              `json:"totalStealers"`
	Employees               int                `json:"employees"`
	Users                   int                `json:"users"`
	ThirdParties            int                `json:"third_parties"`
	Logo                    string             `json:"logo"`
	TotalUrls               int                `json:"totalUrls"`
	Stats                   Stats              `json:"stats"`
	IsShopify               bool               `json:"is_shopify"`
	LastEmployeeCompromised string             `json:"last_employee_compromised"`
	LastUserCompromised     string             `json:"last_user_compromised"`
	Antiviruses             Antiviruses        `json:"antiviruses"`
	Applications            []Application      `json:"applications"`
	EmployeePasswords       PasswordStats      `json:"employeePasswords"`
	UserPasswords           PasswordStats      `json:"userPasswords"`
	ThirdPartyDomains       []DomainOccurrence `json:"thirdPartyDomains"`
	StealerFamilies         map[string]int     `json:"stealerFamilies"`
	Data                    DomainSearchData   `json:"data"`
}

type DomainSearchData struct {
	EmployeesURLs []URLInfo `json:"employees_urls"`
	ClientsURLs   []URLInfo `json:"clients_urls"`
	AllURLs       []URLInfo `json:"all_urls"`
}

type URLInfo struct {
	Occurrence int    `json:"occurrence"`
	Type       string `json:"type"`
	URL        string `json:"H"`
}

type Stats struct {
	TotalEmployees int      `json:"totalEmployees"`
	TotalUsers     int      `json:"totalUsers"`
	EmployeesURLs  []string `json:"employees_urls"`
	ClientsURLs    []string `json:"clients_urls"`
	EmployeesCount []int    `json:"employees_count"`
	ClientsCount   []int    `json:"clients_count"`
}

type Antiviruses struct {
	Total    int         `json:"total"`
	Found    float64     `json:"found"`
	NotFound float64     `json:"not_found"`
	Free     float64     `json:"free"`
	List     []AVProduct `json:"list"`
}

type AVProduct struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type Application struct {
	Keyword string `json:"keyword"`
}

type PasswordStats struct {
	TotalPass int         `json:"totalPass"`
	HasStats  bool        `json:"has_stats"`
	TooWeak   PasswordBin `json:"too_weak"`
	Weak      PasswordBin `json:"weak"`
	Medium    PasswordBin `json:"medium"`
	Strong    PasswordBin `json:"strong"`
}

type PasswordBin struct {
	Qty  int     `json:"qty"`
	Perc float64 `json:"perc"`
}

type DomainOccurrence struct {
	Occurrence int     `json:"occurrence"`
	Domain     *string `json:"domain"` // pointer to handle nulls
}

func (c *Client) DomainSearch(ctx context.Context, domain string) (DomainSearchResponse, error) {
	// Build full URL using BaseURL constant
	endpoint, err := url.Parse(BaseURL)
	if err != nil {
		return DomainSearchResponse{}, err
	}
	endpoint.Path = "/api/json/v2/osint-tools/search-by-domain"

	// Add query parameters
	query := endpoint.Query()
	query.Set("domain", domain)
	endpoint.RawQuery = query.Encode()

	var result DomainSearchResponse

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
		"total_results", result.Total,
		"max_retries", c.MaxRetries)

	return result, nil
}
