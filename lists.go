package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ListResponse struct {
	Id      string
	Name    string
	Created string
	Updated string
}

type ListsResponse struct {
	Paging *PagingResponse
	Items  []*ListResponse
}

type UpdateListRequest struct {
	Name string `json:"name"`
}

type Subscription struct {
	RecipientId string
	Created     string
}

type ListSubscriptionsResponse struct {
	Paging *PagingResponse
	Items  []*Subscription
}

type Recipient struct {
	RecipientId string `json:"recipientId,omitempty"`
}

type SubscribeRecipientsRequest struct {
	Recipients []Recipient `json:"recipients,omitempty"`
}

func (c *Client) GetLists(ctx context.Context, cursor string, pattern string) (*ListsResponse, error) {
	url := c.BaseUrl + "/lists"

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

	bytes, err := c.doRequest(req)
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

func (c *Client) GetList(ctx context.Context, listID string) (*ListResponse, error) {
	if listID == "" {
		return nil, errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s", listID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
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

func (c *Client) PutList(ctx context.Context, listID string, listName string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	if listName == "" {
		return errors.New("List Name is required")
	}

	payload := UpdateListRequest{
		Name: listName,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s", listID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var data string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteList(ctx context.Context, listID string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s", listID)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var data string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RestoreList(ctx context.Context, listID string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s/restore", listID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var data string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetListSubscriptions(ctx context.Context, listID string, cursor string) (*ListSubscriptionsResponse, error) {
	if listID == "" {
		return nil, errors.New("List ID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s/subscriptions", listID)

	if cursor != "" {
		url = url + "?cursor=" + cursor
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
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

func (c *Client) SubscribeMultipleRecipientsToList(ctx context.Context, listID string, recipients []Recipient) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	if recipients == nil || len(recipients) == 0 {
		return errors.New("Recipients cannot be empty")
	}

	payload := SubscribeRecipientsRequest{
		Recipients: recipients,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s/subscriptions", listID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var data string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SubscribeRecipientToList(ctx context.Context, listID string, recipientID string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	if recipientID == "" {
		return errors.New("Recipient ID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s/subscriptions/%s", listID, recipientID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var data string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteListSubscription(ctx context.Context, listID string, recipientID string) error {
	if listID == "" {
		return errors.New("List ID is required")
	}

	if recipientID == "" {
		return errors.New("Recipient ID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/lists/%s/subscriptions/%s", listID, recipientID)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var data string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}
