package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type SendResponse struct {
	MessageId string
}

type sendRequest struct {
	EventID     string      `json:"event"`
	Recipient   string      `json:"recipient"`
	Profile     interface{} `json:"profile,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Brand       string      `json:"brand,omitempty"`
	Override    interface{} `json:"override,omitempty"`
	Preferences interface{} `json:"preferences,omitempty"`
}

type sendToListRequest struct {
	EventID  string      `json:"event"`
	ListID   string      `json:"list,omitempty"`
	Pattern  string      `json:"pattern,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Brand    string      `json:"brand,omitempty"`
	Override interface{} `json:"override,omitempty"`
}

func (c *Client) Send(ctx context.Context, eventID string, recipientID string, profile interface{}, data interface{}, brand string, override interface{}, preferences interface{}, idempotencyKey string) (string, error) {
	payload := sendRequest{
		EventID:     eventID,
		Recipient:   recipientID,
		Profile:     profile,
		Data:        data,
		Brand:       brand,
		Override:    override,
		Preferences: preferences,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseUrl+"/send", bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	// Idempotent Request check
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
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

func (c *Client) SendToList(ctx context.Context, eventID string, listID string, pattern string, data interface{}, brand string, override interface{}, idempotencyKey string) (string, error) {
	if (listID == "" && pattern == "") || (listID != "" && pattern != "") {
		return "", errors.New("list.send requires a listID or a pattern")
	}

	payload := sendToListRequest{
		EventID:  eventID,
		ListID:   listID,
		Pattern:  pattern,
		Data:     data,
		Brand:    brand,
		Override: override,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseUrl+"/send/list", bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	// Idempotent Request check
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
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
