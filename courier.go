package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const version = "2.0.0"

// Client represents the core data needed to communicate with the Courier API
type Client struct {
	AuthToken  string
	BaseURL    string
	SDKVersion string
}

// CreateClient creates a new client for communicating with the Courier API
func CreateClient(authToken string, baseURL *string) *Client {
	if baseURL == nil {
		return &Client{
			AuthToken:  authToken,
			BaseURL:    "https://api.trycourier.app",
			SDKVersion: version,
		}
	}

	deref := *baseURL
	return &Client{
		AuthToken:  authToken,
		BaseURL:    deref,
		SDKVersion: version,
	}
}

// HTTPSendJSON wraps HTTPSendBytes
func (c *Client) HTTPSendJSON(ctx context.Context, method string, url string, body interface{}, response interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bytes, err := c.HTTPSendBytes(ctx, method, url, bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return err
	}

	return nil
}

// HTTPSendBytes wraps HTTPRequest
func (c *Client) HTTPSendBytes(ctx context.Context, method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+url, body)
	if err != nil {
		return nil, err
	}

	return c.HTTPRequest(req)
}

// HTTPRequest issues an HTTP request and sets headers expected by the Courier API
func (c *Client) HTTPRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "courier-go/"+c.SDKVersion)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
