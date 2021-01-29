package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type EventResponse struct {
	Event string
	Id    string
	Type  string
}

type EventsResponse struct {
	Results []*EventResponse
}

type UpsertEventRequest struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type UpsertEventResponse struct {
	Status string
}

func (c *Client) GetEvents(ctx context.Context) (*EventsResponse, error) {
	url := fmt.Sprintf(c.BaseUrl + "/events")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data EventsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) GetEvent(ctx context.Context, eventID string) (*EventResponse, error) {
	if eventID == "" {
		return nil, errors.New("eventID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/events/%s", eventID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data EventResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) PutEvent(ctx context.Context, eventID string, notificationID string, eventType string) (*UpsertEventResponse, error) {
	if eventID == "" {
		return nil, errors.New("eventID is required")
	}

	if notificationID == "" {
		return nil, errors.New("notificationID is required")
	}

	if eventType == "" {
		return nil, errors.New("eventType is required")
	}

	payload := UpsertEventRequest{
		Id:   notificationID,
		Type: eventType,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(c.BaseUrl+"/events/%s", eventID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data UpsertEventResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
