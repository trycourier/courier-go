package courier

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProvidersChannelResponse struct {
	Key string
	Name string
	Template string
}
type ProvidersResponse struct {
	Channel *ProvidersChannelResponse
	Error string
	Status string
	Delivered int64
	Sent int64
	Clicked int64
	Provider string
	Reference interface{} // provider specific response
}

type MessageResponse struct {
	Id string
	Event string
	Notification string
	Status string
	Error string
	Reason string
	Recipient string
	Enqueued int64
	Delivered int64
	Sent int64
	Clicked int64
	Providers []*ProvidersResponse
}

func (c *Client) GetMessage(ctx context.Context, messageID string) (*MessageResponse, error) {

	url := fmt.Sprintf(c.BaseUrl+"/messages/%s", messageID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data MessageResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
