package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type SendResponse struct {
	MessageId string
}

type sendRequest struct {
	EventID   string      `json:"event"`
	Recipient string      `json:"recipient"`
	Brand     string      `json:"brand"`
	Profile   interface{} `json:"profile"`
	Data      interface{} `json:"data"`
}

func (c *Client) Send(ctx context.Context, eventID, recipientID, brandID string, profile, data interface{}) (string, error) {
	payload := sendRequest{
		EventID:   eventID,
		Recipient: recipientID,
		Brand:     brandID,
		Profile:   profile,
		Data:      data,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseUrl+"/send", bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return "", err
	}
	var responseData SendResponse
	err = json.Unmarshal(bytes, &responseData)
	if err != nil {
		return "", err
	}

	return responseData.MessageId, nil
}
