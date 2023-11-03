package courier

const version = "2.12.0"

// Client lets you communicate with the Courier API
type Client struct {
	API *APIConfiguration
}

// CreateClient creates a new client for communicating with the Courier API
func CreateClient(authToken string, baseURL *string) *Client {
	var url string
	if baseURL == nil {
		url = "https://api.courier.com"
	} else {
		url = *baseURL
	}

	return &Client{
		API: &APIConfiguration{
			AuthToken:  authToken,
			BaseURL:    url,
			SDKVersion: version,
		},
	}
}
