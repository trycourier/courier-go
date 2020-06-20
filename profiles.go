package courier

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

// Profile represents the return of the Courier profiles API
type Profile struct {
	Profile interface{} `json:"profile"`
}

// GetProfile calls the GET /profiles/:id endpoint of the Courier API
func (c *Client) GetProfile(ctx context.Context, id string) (*Profile, error) {
	var response Profile
	err := c.HTTPSendJSON(ctx, "GET", fmt.Sprintf("/profiles/%s", id), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// MergeProfile calls the POST /profiles/:id endpoint of the Courier API
func (c *Client) MergeProfile(ctx context.Context, id string, profile []byte) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("/profiles/%s", id), bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = c.HTTPRequest(req)
	return err
}

// UpdateProfile calls the PUT /profiles/:id endpoint of the Courier API
func (c *Client) UpdateProfile(ctx context.Context, id string, profile []byte) error {
	url := fmt.Sprintf(c.BaseURL+"/profiles/%s", id)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = c.HTTPRequest(req)
	return err
}
