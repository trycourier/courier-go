package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ListResponse struct {
	ID      string
	Name    string
	Created string
	Updated string
}

type ListsResponse struct {
	Paging *PagingResponse
	Items  []*ListResponse
}

type PutListBody struct {
	Name        string      `json:"name"`
	Preferences interface{} `json:"preferences,omitempty"`
}

type ListSubscriptionItem struct {
	RecipientID string
	Created     string
	Preferences interface{}
}

type ListSubscriptionsResponse struct {
	Paging *PagingResponse
	Items  []*ListSubscriptionItem
}

type ListRecipient struct {
	RecipientID string      `json:"recipientId,omitempty"`
	Preferences interface{} `json:"preferences,omitempty"`
}

type ListSubscriptionBody struct {
	Recipients []ListRecipient `json:"recipients"`
}

type ListSubscriptionRecipientBody struct {
	Preferences interface{} `json:"preferences,omitempty"`
}

// GetLists calls the GET /lists endpoint of the Courier API
func (c *Client) GetLists(ctx context.Context, cursor string, pattern string) (*ListsResponse, error) {
	url := c.API.BaseURL + "/lists"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if cursor != "" {
		q.Add("cursor", cursor)
	}
	if pattern != "" {
		q.Add("pattern", pattern)
	}
	req.URL.RawQuery = q.Encode()

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data ListsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetList calls the GET /lists/{list_id} endpoint of the Courier API
func (c *Client) GetList(ctx context.Context, listID string) (*ListResponse, error) {
	if listID == "" {
		return nil, errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/lists/%s", listID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data ListResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// PutList calls the PUT /lists/{list_id} endpoint of the Courier API
func (c *Client) PutList(ctx context.Context, listID string, body interface{}) error {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return err
	}

	_, err = c.API.SendRequestWithMaps(ctx, "PUT", "/lists/"+listID, bodyMap)
	if err != nil {
		return err
	}

	return nil
}

// DeleteList calls the DELETE /lists/{list_id} endpoint of the Courier API
func (c *Client) DeleteList(ctx context.Context, listID string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/lists/%s", listID)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = c.API.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}

// RestoreList calls the PUT /lists/{list_id}/restore endpoint of the Courier API
func (c *Client) RestoreList(ctx context.Context, listID string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/lists/%s/restore", listID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return err
	}

	_, err = c.API.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}

// GetListSubscriptions calls the GET /lists/{list_id}/subscriptions endpoint of the Courier API
func (c *Client) GetListSubscriptions(ctx context.Context, listID string, cursor string) (*ListSubscriptionsResponse, error) {
	if listID == "" {
		return nil, errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/lists/%s/subscriptions", listID)

	if cursor != "" {
		url = url + "?cursor=" + cursor
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data ListSubscriptionsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// PutListSubscriptions calls the PUT /lists/{list_id}/subscriptions endpoint of the Courier API
func (c *Client) PutListSubscriptions(ctx context.Context, listID string, body interface{}) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	bodyMap, err := toJSONMap(body)
	if err != nil {
		return err
	}

	_, err = c.API.SendRequestWithMaps(ctx, "PUT", "/lists/"+listID+"/subscriptions", bodyMap)
	if err != nil {
		return err
	}

	return nil
}

// PostListSubscriptions calls the POST /lists/{list_id}/subscriptions endpoint of the Courier API
func (c *Client) PostListSubscriptions(ctx context.Context, listID string, body interface{}) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	bodyMap, err := toJSONMap(body)
	if err != nil {
		return err
	}

	_, err = c.API.SendRequestWithMaps(ctx, "POST", "/lists/"+listID+"/subscriptions", bodyMap)
	if err != nil {
		return err
	}

	return nil
}

// ListSubscribe calls the PUT /lists/{list_id}/subscriptions/{recipient_id} endpoint of the Courier API
func (c *Client) ListSubscribe(ctx context.Context, listID string, recipientID string, body interface{}) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	if recipientID == "" {
		return errors.New("Recipient ID is required")
	}

	_, err := toJSONMap(body)
	if err != nil {
		return err
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.API.SendRequestWithBytes(ctx, "PUT", "/lists/"+listID+"/subscriptions/"+recipientID, jsonBody)
	if err != nil {
		return err
	}

	return nil
}

// ListUnsubscribe calls the PUT /lists/{list_id}/subscriptions/{recipient_id} endpoint of the Courier API
func (c *Client) ListUnsubscribe(ctx context.Context, listID string, recipientID string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	if recipientID == "" {
		return errors.New("Recipient ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/lists/%s/subscriptions/%s", listID, recipientID)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = c.API.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}
