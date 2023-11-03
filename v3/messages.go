package courier

import (
	"context"
	"fmt"
)

// ProvidersChannelResponse represents the channel section of the ProvidersResponse
type ProvidersChannelResponse struct {
	Key      string
	Name     string
	Template string
}

// ProvidersResponse represents the providers section of the MessageResponse
type ProvidersResponse struct {
	Channel          *ProvidersChannelResponse
	Error            string
	Status           string
	Delivered        int64
	Sent             int64
	Clicked          int64
	Provider         string
	ProviderResponse interface{}
	Reference        interface{}
}

// MessageResponse represents the return of the /messages/* endpoints on the Courier API
type MessageResponse struct {
	ID            string
	Event         string
	Notification  string
	Status        string
	Error         string
	Reason        string
	Recipient     string
	JobId         string
	ListId        string
	ListMessageId string
	RunId         string
	Enqueued      int64
	Delivered     int64
	Sent          int64
	Clicked       int64
	Opened        int64
	CanceledAt    int64
	Providers     []*ProvidersResponse
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

// CancelMessage calls the /messages/:id/cancel endpoint of the Courier API
func (c *Client) CancelMessage(ctx context.Context, messageID string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.API.SendRequestWithJSON(ctx, "POST", fmt.Sprintf("/messages/%s/cancel", messageID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
