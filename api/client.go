package api

import (
	"resty.dev/v3"
)

const BaseURL = "https://cavalier.hudsonrock.com"

// Client is a reusable HTTP client for the Hudson Rock API using Resty.
type Client struct {
	Resty   *resty.Client
	BaseURL string
}

// NewClient returns a new Client with a Resty client and the Hudson Rock API base URL.
func NewClient() *Client {
	return &Client{
		Resty:   resty.New(),
		BaseURL: BaseURL,
	}
}
