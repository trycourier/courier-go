package courier

import (
	"context"
)

// ProvidersChannelResponse represents the channel section of the ProvidersResponse
type ProvidersChannelResponse struct {
	Key      string
	Name     string
	Template string
}

// ProvidersResponse represents the providers section of the MessageResponse
type ProvidersResponse struct {
	Channel   *ProvidersChannelResponse
	Error     string
	Status    string
	Delivered int64
	Sent      int64
	Clicked   int64
	Provider  string
	Reference interface{} // provider specific response
}

// MessageResponse represents the return of the /messages/* endpoints on the Courier API
type MessageResponse struct {
	ID           string
	Event        string
	Notification string
	Status       string
	Error        string
	Reason       string
	Recipient    string
	Enqueued     int64
	Delivered    int64
	Sent         int64
	Clicked      int64
	Providers    []*ProvidersResponse
}

// GetMessage calls the /messages/:id endpoint of the Courier API
func (c *Client) GetMessage(ctx context.Context, messageID string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.API.SendRequestWithJSON(ctx, "GET", "/messages/"+messageID, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
