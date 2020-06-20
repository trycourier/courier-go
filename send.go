package courier

import (
	"context"
)

// SendResponse is the response returned by the CourierAPI upon success
type SendResponse struct {
	MessageID string `json:"messageId"`
}

// SendRequest is the JSON body expected by Courier's /send endpoint
type SendRequest struct {
	EventID   string      `json:"event"`
	Recipient string      `json:"recipient"`
	Profile   interface{} `json:"profile"`
	Data      interface{} `json:"data"`
}

// Send calls the /send endpoint of the Courier API
func (c *Client) Send(ctx context.Context, eventID, recipientID string, profile, data interface{}) (string, error) {
	payload := SendRequest{
		EventID:   eventID,
		Recipient: recipientID,
		Profile:   profile,
		Data:      data,
	}

	var response SendResponse
	err := c.HTTPSendJSON(ctx, "POST", "/send", payload, &response)
	if err != nil {
		return "", err
	}

	return response.MessageID, nil
}

// func (c *Client) SendMap(ctx context.Context, eventID, recipientID string, porfile map[string]interface{}, data interface{}) (string, error) {
// }
