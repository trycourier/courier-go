package courier

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageResponse struct {
	Id string
	Status string
	Enqueued int64
	Delivered int64
	Provider string
	RecipientId string
	EventId string
	Configuration string
	ProviderResponse string
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
