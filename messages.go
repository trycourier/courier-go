package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type ProvidersChannelResponse struct {
	Key      string
	Name     string
	Template string
}
type ProvidersResponse struct {
	Channel   *ProvidersChannelResponse
	Error     string
	Status    string
	Delivered int64
	Sent      int64
	Clicked   int64
	Opened    int64
	Provider  string
	Reference interface{} // provider specific response
}

type MessageResponse struct {
	Id           string
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
	Opened       int64
	Providers    []*ProvidersResponse
}

type PagingResponse struct {
	Cursor string
	More   bool
}

type MessagesResponse struct {
	Paging  *PagingResponse
	Results []*MessageResponse
}

type MessageHistoryResponse struct {
	Results []interface{}
}

func (c *Client) GetMessage(ctx context.Context, messageID string) (*MessageResponse, error) {
	if messageID == "" {
		return nil, errors.New("messageID is required")
	}

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

func (c *Client) GetMessages(ctx context.Context, cursor string, event string, list string, messageId string, notification string, recipient string) (*MessagesResponse, error) {
	params := url.Values{}
	if cursor != "" {
		params.Add("cursor", cursor)
	}
	if event != "" {
		params.Add("event", event)
	}
	if list != "" {
		params.Add("list", list)
	}
	if messageId != "" {
		params.Add("messageId", messageId)
	}
	if notification != "" {
		params.Add("notification", notification)
	}
	if recipient != "" {
		params.Add("recipient", recipient)
	}

	url := fmt.Sprintf(c.BaseUrl+"/messages?%s", params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data MessagesResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) GetMessageHistory(ctx context.Context, messageID string, _type string) (*MessageHistoryResponse, error) {
	if messageID == "" {
		return nil, errors.New("messageID is required")
	}

	requestURL := fmt.Sprintf(c.BaseUrl+"/messages/%s/history", messageID)

	if _type != "" {
		requestURL = requestURL + "?type=" + _type
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data MessageHistoryResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
