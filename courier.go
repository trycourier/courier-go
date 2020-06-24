package courier

import "github.com/trycourier/courier-go/http"

const version = "2.0.0"

// Client lets you communicate with the Courier API
type Client struct {
	API *http.APIConfiguration
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
		API: &http.APIConfiguration{
			AuthToken:  authToken,
			BaseURL:    url,
			SDKVersion: version,
		},
	}
}
