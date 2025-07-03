package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FileSearchRequest represents the request body for the /search-by-file endpoint.
type FileSearchRequest struct {
	FileName      string `json:"file_name"`
	StartDate     string `json:"start_date,omitempty"`
	EndDate       string `json:"end_date,omitempty"`
	SortBy        string `json:"sort_by,omitempty"`
	SortDirection string `json:"sort_direction,omitempty"`
	Cursor        string `json:"cursor,omitempty"`
}

// FileSearchResponse represents the response from the /search-by-file endpoint.
type FileSearchResponse struct {
	Data       []map[string]interface{} `json:"data"`
	NextCursor string                   `json:"nextCursor"`
}

// Client is a Hudson Rock API client.
type Client struct {
	BaseURL string
}

// NewClient creates a new Hudson Rock API client.
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}
