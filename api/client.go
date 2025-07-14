package api

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"resty.dev/v3"
)

const BaseURL = "https://cavalier.hudsonrock.com"

// Client is a reusable HTTP client for the Hudson Rock API using Resty.
type Client struct {
	Resty      *resty.Client
	BaseURL    string
	MaxRetries int
	MinDelay   time.Duration
	rand       *rand.Rand
}

// NewClient returns a new Client with a Resty client and the Hudson Rock API base URL.
func NewClient() *Client {
	client := resty.New()

	// Configure timeouts
	client.SetTimeout(30 * time.Second)

	// Basic retry configuration if available
	if err := client.SetRetryCount(3); err != nil {
		log.Printf("[WARNING] Could not set retry count: %v", err)
	}

	return &Client{
		Resty:      client,
		BaseURL:    BaseURL,
		MaxRetries: 3,                                               // Default to 3 retries
		MinDelay:   100 * time.Millisecond,                          // Default minimum delay
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())), // Modern random source
	}
}

// WithMaxRetries sets the maximum number of retries for the client
func (c *Client) WithMaxRetries(maxRetries int) *Client {
	c.MaxRetries = maxRetries
	return c
}

// WithMinDelay sets the minimum delay for backoff calculations
func (c *Client) WithMinDelay(minDelay time.Duration) *Client {
	c.MinDelay = minDelay
	return c
}

// buildURL constructs the full URL for an API endpoint path.
func (c *Client) buildURL(path string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, path)
}

// BackoffDelay returns the duration to wait before the next attempt should be
// made. Returns an error if unable get a duration.
func (c *Client) BackoffDelay(attempt int, err error) (time.Duration, error) {
	minDelay := c.MinDelay

	// The calculated jitter will be between [0.8, 1.2)
	var jitter = float64(c.rand.Intn(120-80)+80) / 100

	retryTime := time.Duration(int(float64(int(minDelay.Nanoseconds())*int(math.Pow(3, float64(attempt)))) * jitter))

	// Cap retry time at 5 minutes to avoid too long a wait
	if retryTime > time.Duration(5*time.Minute) {
		retryTime = time.Duration(5 * time.Minute)
	}

	// Low level method to log retries since we don't have context etc here.
	// Logging is helpful for visibility into retries and choke points in using
	// the API.
	log.Printf("[INFO] BackoffDelay: attempt=%d, retryTime=%s, err=%v", attempt, retryTime.String(), err)

	return retryTime, nil
}

// executeWithRetry performs an HTTP request with manual retry logic
func (c *Client) executeWithRetry(request func() (*resty.Response, error), maxRetries int) (*resty.Response, error) {
	var lastErr error
	var resp *resty.Response

	for attempt := 1; attempt <= maxRetries; attempt++ {
		log.Printf("[REQUEST] Attempt %d/%d", attempt, maxRetries)

		resp, lastErr = request()

		if lastErr == nil && resp != nil {
			statusCode := resp.StatusCode()
			log.Printf("[RESPONSE] Status: %d", statusCode)

			// Success - no need to retry
			if statusCode >= 200 && statusCode < 300 {
				return resp, nil
			}

			// Check if we should retry based on status code
			shouldRetry := false
			switch statusCode {
			case 429: // Rate limited
				log.Printf("[RETRY] Rate limited (429), retrying...")
				shouldRetry = true
			case 408: // Request timeout
				log.Printf("[RETRY] Request timeout (408), retrying...")
				shouldRetry = true
			case 500, 502, 503, 504: // Server errors
				log.Printf("[RETRY] Server error (%d), retrying...", statusCode)
				shouldRetry = true
			default:
				if statusCode >= 400 && statusCode < 500 {
					log.Printf("[NO RETRY] Client error (%d), not retrying", statusCode)
					return resp, fmt.Errorf("HTTP client error: %d %s", statusCode, resp.Status())
				}
			}

			if !shouldRetry {
				return resp, fmt.Errorf("HTTP error: %d %s", statusCode, resp.Status())
			}
		} else if lastErr != nil {
			log.Printf("[RETRY] Network error: %v", lastErr)
		}

		// Don't sleep after the last attempt
		if attempt < maxRetries {
			backoff, err := c.BackoffDelay(attempt, lastErr)
			if err != nil {
				log.Printf("[ERROR] Failed to calculate backoff delay: %v", err)
				backoff = 1 * time.Second // fallback
			}

			time.Sleep(backoff)
		}
	}

	if lastErr != nil {
		return resp, fmt.Errorf("request failed after %d attempts: %w", maxRetries, lastErr)
	}

	return resp, fmt.Errorf("request failed after %d attempts with status %d", maxRetries, resp.StatusCode())
}

// executeWithRetryDefault performs an HTTP request using the client's default retry settings
func (c *Client) executeWithRetryDefault(request func() (*resty.Response, error)) (*resty.Response, error) {
	return c.executeWithRetry(request, c.MaxRetries)
}
