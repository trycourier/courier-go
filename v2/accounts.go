package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Account is the type for account entity
type Account struct {
	Id                 string      `json:"id"`
	Name               string      `json:"name"`
	ParentAccountId    string      `json:"parent_account_id,omitempty"`
	DefaultPreferences interface{} `json:"default_preferences,omitempty"`
	Properties         interface{} `json:"properties,omitempty"`
	UserProfile        interface{} `json:"user_profile,omitempty"`
	BrandId            string      `json:"brand_id,omitempty"`
}

// ListAccountsResponse is type for the GET /accounts response
type ListAccountsResponse struct {
	Items   []Account `json:"items"`
	HasMore bool      `json:"has_more'`
	Url     string    `json:"url"`
	NextUrl string    `json:"next_url,omitempty"`
}

// GetAccount calls the GET /accounts/:accountID endpoint of the Courier API
func (c *Client) GetAccount(ctx context.Context, accountID string) (*Account, error) {
	if accountID == "" {
		return nil, errors.New("account ID is required")
	}
	url := fmt.Sprintf(c.API.BaseURL+"/accounts/%s", accountID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data Account
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetAccounts calls the GET /accounts endpoint of the Courier API
func (c *Client) GetAccounts(ctx context.Context, limit string, cursor string) (*ListAccountsResponse, error) {
	url := fmt.Sprintf(c.API.BaseURL + "/accounts")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if limit != "" {
		q.Add("limit", limit)
	}
	if cursor != "" {
		q.Add("cursor", cursor)
	}
	req.URL.RawQuery = q.Encode()

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data ListAccountsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Calls the PUT /accounts/:accountID endpoint of the Courier API
func (c *Client) PutAccount(ctx context.Context, accountID string, body interface{}) (*Account, error) {
	if accountID == "" {
		return nil, errors.New("account ID is required")
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	bytes, err := c.API.SendRequestWithBytes(ctx, http.MethodPut, fmt.Sprintf("/accounts/%s", accountID), jsonBody)

	if err != nil {
		return nil, err
	}

	var data Account

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// DeleteAccount calls the DELETE /accounts/:account_id endpoint of the Courier API
func (c *Client) DeleteAccount(ctx context.Context, accountID string) error {
	if accountID == "" {
		return errors.New("account ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/accounts/%s", accountID)

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
