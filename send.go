package courier

import (
	"context"
	"encoding/json"
)

// SendBody is the JSON body expected by Courier's /send endpoint
type SendBody struct {
	Profile interface{} `json:"profile"`
	Data    interface{} `json:"data"`
}

// Send calls the /send endpoint of the Courier API (passing a struct)
func (c *Client) Send(ctx context.Context, eventID, recipientID string, body interface{}) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}

	// these are required, so we accept them as separate params
	bodyMap["event"] = eventID
	bodyMap["recipient"] = recipientID

	return c.SendMap(ctx, eventID, recipientID, bodyMap)
}

// SendMap calls the /send endpoint of the Courier API (passing maps)
func (c *Client) SendMap(ctx context.Context, eventID, recipientID string, body map[string]interface{}) (string, error) {
	body["event"] = eventID
	body["recipient"] = recipientID

	response, err := c.http.SendRequestWithMaps(ctx, "POST", "/send", body)
	if err != nil {
		return "", err
	}

	var messageID string
	err = json.Unmarshal(response["messageId"], &messageID)
	if err != nil {
		return "", err
	}
	return messageID, nil
}
