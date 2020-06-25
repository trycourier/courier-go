package courier

import "github.com/trycourier/courier-go/api"

const version = "2.0.0"

// Client lets you communicate with the Courier API
type Client struct {
	API *api.Configuration
}

// CreateClient creates a new client for communicating with the Courier API
func CreateClient(authToken string, baseURL *string) *Client {
	var url string
	if baseURL == nil {
		url = "https://api.trycourier.app"
	} else {
		url = *baseURL
	}

	return &Client{
		API: &api.Configuration{
			AuthToken:  authToken,
			BaseURL:    url,
			SDKVersion: version,
		},
	}
}
